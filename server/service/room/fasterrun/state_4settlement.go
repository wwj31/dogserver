package fasterrun

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"

	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/common/rds"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/rdsop"

	"github.com/wwj31/dogactor/tools"
)

// 结算状态

type StateSettlement struct {
	*FasterRun
}

func (s *StateSettlement) State() int {
	return Settlement
}

func (s *StateSettlement) Enter() {
	s.currentStateEndAt = tools.Now().Add(SettlementDuration)
	s.Log().Infow("[FasterRun] enter state settlement", "room", s.room.RoomId,
		"master", s.masterIndex, "game count", s.gameCount, "endAt", s.currentStateEndAt.UnixMilli())

	playerNumber := s.playerNumber()
	settlementMsg := &outer.FasterRunSettlementNtf{
		EndAt:            s.currentStateEndAt.UnixMilli(),
		HasScoreZero:     s.scoreZeroOver,
		GameCount:        int32(s.gameCount),
		GameSettlementAt: tools.Now().UnixMilli(),
		PlayerData:       make([]*outer.FasterRunSettlementPlayerData, playerNumber, playerNumber),
	}

	for seat := 0; seat < playerNumber; seat++ {
		settlementMsg.PlayerData[seat] = &outer.FasterRunSettlementPlayerData{}
	}

	// 大结算
	if s.finalSettlement() || s.scoreZeroOver {
		s.Log().Infow("final settlement",
			"scoreZeroOver", s.scoreZeroOver, "game count", s.gameCount, "param", s.gameParams().PlayCountLimit)
		s.gameCount = int(s.gameParams().PlayCountLimit)

		// 总抽水
		totalProfit := s.profit(s.gameParams().BigWinner)
		s.rebate(totalProfit)

		ntf := &outer.FasterRunFinialSettlement{}
		for seat := 0; seat < playerNumber; seat++ {
			player := s.fasterRunPlayers[seat]
			ntf.PlayerInfo = append(ntf.PlayerInfo, player.finalStatsMsg)
			player.finalStatsMsg = &outer.FasterRunFinialPlayerInfo{}
		}
		settlementMsg.FinalSettlement = ntf

	}

	// 结算分数为最终金币
	modifyRspCount := make(map[string]struct{}) // 必须等待所有玩家金币修改成功后，才能发送结算
	for i := 0; i < playerNumber; i++ {
		seat := i
		player := s.fasterRunPlayers[seat]
		finalScore := player.score
		s.room.Request(actortype.PlayerId(player.RID), &inner.ModifyGoldReq{
			SetOrAdd: true,
			Gold:     finalScore,
		}).Handle(func(resp any, err error) {
			modifyRspCount[player.RID] = struct{}{}
			if err == nil {
				modifyRsp := resp.(*inner.ModifyGoldRsp)
				player.PlayerInfo = modifyRsp.Info
			}

			s.Log().Infow("modify gold result", "room", s.room.RoomId, "seat", seat,
				"player", player.ShortId, "latest gold", player.Gold, "err", err)

			if len(modifyRspCount) == playerNumber {
				s.afterSettle(settlementMsg)
			}
		})
	}

}

func (s *StateSettlement) Leave() {
	s.Log().Infow("[FasterRun] leave state settlement ==================SETTLEMENT==================", "room", s.room.RoomId, "count", s.gameCount)
	s.Log().Infof(" ")
	s.Log().Infof(" ")
	s.Log().Infof(" ")

}

func (s *StateSettlement) Handle(shortId int64, v any) (result any) {
	return outer.ERROR_MAHJONG_STATE_MSG_INVALID
}

func (s *StateSettlement) afterSettle(ntf *outer.FasterRunSettlementNtf) {
	allPlayerInfo := s.playersToPB(0, true) // 组装结算消息

	for seat, player := range s.fasterRunPlayers {
		ntf.PlayerData[seat].Player = allPlayerInfo[seat]
		ntf.PlayerData[seat].BombsCount = player.BombsCount
		ntf.PlayerData[seat].TotalScore = player.totalWinScore
	}

	s.Log().Infow(" settlement broadcast", "room", s.room.RoomId,
		"master", s.masterIndex, "ntf", ntf.String())

	s.room.Broadcast(ntf)
	s.clear() // 分算完清理数据

	// 结算给个短暂的时间
	s.room.AddTimer(tools.XUID(), s.currentStateEndAt, func(dt time.Duration) {
		s.SwitchTo(Ready)
	})
}

func (s *StateSettlement) finalSettlement() bool {
	if int32(s.gameCount) >= s.gameParams().PlayCountLimit {
		return true
	}

	return false
}

// profit 抽水
func (s *StateSettlement) profit(bigWinner bool) (totalProfit int64) {
	var (
		winners []*fasterRunPlayer
	)

	if bigWinner {
		var (
			winScore int64
			winner   *fasterRunPlayer
		)

		for i, player := range s.fasterRunPlayers {
			if player.finalStatsMsg.TotalScore > winScore {
				winner = s.fasterRunPlayers[i]
				winScore = player.finalStatsMsg.TotalScore
			}
		}

		if winners != nil {
			winners = append(winners, winner)
		}
	} else {
		for i, player := range s.fasterRunPlayers {
			if player.finalStatsMsg.TotalScore > 0 {
				winners = append(winners, s.fasterRunPlayers[i])
			}
		}
	}

	// 处理每一位赢家抽水
	for _, winner := range winners {
		rangeCfg := s.profitRange(winner)
		winScore := winner.finalStatsMsg.TotalScore
		if rangeCfg == nil {
			s.Log().Warnw("winner profit score not in any range",
				"big winners", winner.ShortId, "win", winScore)
			continue
		}

		// 检查是否达到抽水最低要求
		if winScore < rangeCfg.MinimumRebate {
			s.Log().Warnw("the winner win the score did not meet the expected score",
				"big winners", winner.ShortId, "win", winScore)
			continue
		}

		baseProfit := (winScore * rangeCfg.RebateRatio) / 100 // 基础抽水
		profit := baseProfit + rangeCfg.MinimumGuarantee      // +抽水保底值
		totalProfit += profit
		val := winner.score - profit
		winner.score = common.Max(0, val)
		s.Log().Infow("profit", "winner", winner.ShortId, "current score", winner.score,
			"baseProfit", baseProfit, "val", val, "winScore", winScore, "profit", profit, "range param", rangeCfg.String())
	}

	return
}

// 找出玩家赢分所在区间的参数配置
func (s *StateSettlement) profitRange(winner *fasterRunPlayer) *outer.RangeParams {
	params := s.gameParams().ReBate
	if params == nil {
		return nil
	}

	for _, param := range []*outer.RangeParams{params.RangeL1, params.RangeL2, params.RangeL3, params.RangeL4} {
		if !param.Valid {
			continue
		}

		minLimit := param.Min * 1000
		MaxLimit := param.Max * 1000
		if winner.finalStatsMsg == nil {
			s.Log().Warnw("final stats msg is nil", "short", winner.ShortId)
			continue
		}

		totalWin := winner.finalStatsMsg.TotalScore
		if minLimit < totalWin && totalWin <= MaxLimit {
			return param
		}
	}
	return nil
}

// rebate 返利计算
func (s *StateSettlement) rebate(totalProfit int64) {
	divProfit := totalProfit / 4 // 每个玩家那条上级路线，都均分获得利润
	if divProfit == 0 {
		s.Log().Infow("profit is zero!!", "total profit", totalProfit)
		return
	}

	pip := rds.Ins.Pipeline()
	for _, player := range s.fasterRunPlayers {
		s.Log().Infow("recur profit start", "start short", player.ShortId)
		s.recurRebate(divProfit, player.UpShortId, player.ShortId, 0, pip)
	}

	if _, err := pip.Exec(context.Background()); err != nil {
		log.Errorw("rebate redis failed", "err", err)
	}
}

// 逐层向上返利
func (s *StateSettlement) recurRebate(profitGold, upShortId, shortId, downShortId int64, addPip redis.Pipeliner) {
	rebateInfo := rdsop.GetRebateInfo(shortId)

	var (
		exactPoint      int32 // 确切的获利点位
		exactProfitGold int64 // 确切的获利
	)

	// 自己有点位，先分自己一份
	if rebateInfo.Point > 0 {
		// 自己的获利=自己的分润-下级的分润
		exactPoint = common.Max(0, rebateInfo.Point-rebateInfo.DownPoints[downShortId])
		exactProfitGold = profitGold * int64(exactPoint) / 100
		rdsop.AddRebateGold(shortId, exactProfitGold, addPip)
	}

	s.Log().Infow("rebate calculating ...", "room", s.room.RoomId,
		"short", shortId, "up", upShortId, "down", downShortId,
		"profitGold", profitGold, "point", exactPoint, "gold", exactProfitGold, "rebateInfo", rebateInfo)

	if upShortId == 0 {
		return
	}

	// 向上级递归
	s.recurRebate(profitGold, rdsop.AgentUp(upShortId), upShortId, shortId, addPip)
}
