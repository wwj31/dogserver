package niuniu

import (
	"time"

	"server/rdsop"

	"github.com/wwj31/dogactor/tools"

	"server/common"
	"server/common/actortype"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
)

// 结算状态

type StateSettlement struct {
	*NiuNiu
}

func (s *StateSettlement) State() int {
	return Settlement
}

func (s *StateSettlement) Enter() {
	s.currentStateEndAt = tools.Now().Add(SettlementDuration)
	s.Log().Infow("[NiuNiu] enter state settlement", "room", s.room.RoomId,
		"master", s.masterIndex, "endAt", s.currentStateEndAt.UnixMilli())

	settlementMsg := &outer.NiuNiuSettlementNtf{
		EndAt:            s.currentStateEndAt.UnixMilli(),
		HasScoreZero:     s.scoreZeroOver,
		GameSettlementAt: tools.Now().UnixMilli(),
		PlayerData:       make([]*outer.NiuNiuSettlementPlayerData, 10, 10),
	}

	for seat := 0; seat < len(s.niuniuPlayers); seat++ {
		player := s.niuniuPlayers[seat]
		if player != nil && player.ready {
			settlementMsg.PlayerData[seat] = &outer.NiuNiuSettlementPlayerData{}
		}
	}

	// 结算输赢分
	var (
		winner     *niuniuPlayer
		winnerSeat int
	)
	for seat, player := range s.niuniuPlayers {
		if len(player.handCards) == 0 {
			winner = player
			winnerSeat = seat
			break
		}
	}

	if winner != nil {
		loserSeats := s.allSeats(winnerSeat)
		for _, seat := range loserSeats {
			loser := s.niuniuPlayers[seat]
			loseScore := int64(len(loser.handCards)) * s.baseScore()

			// 不允许负分，能扣多少扣多少
			if !s.gameParams().AllowScoreSmallZero {
				loseScore = common.Min(loser.score, loseScore)
			}

			loser.updateScore(-loseScore)
			winner.updateScore(loseScore)
			s.Log().Infow("settle loser", "seat", seat, "loser", loser.ShortId, "score", loseScore)
		}
	}

	// 大结算
	if s.finalSettlement() || s.scoreZeroOver {
		s.Log().Infow("final settlement",
			"scoreZeroOver", s.scoreZeroOver, "param", s.gameParams().PlayCountLimit)

		// 总抽水
		totalProfit := s.profit(false)

		// 记录返利信息
		record := &outer.RebateDetailInfo{
			Type:      outer.GameType_FasterRun,
			BaseScore: s.gameParams().BaseScore,
			CreateAt:  tools.Now().UnixMilli(),
		}
		s.room.Rebate(record, totalProfit, s.toRoomPlayers())

		ntf := &outer.NiuNiuFinialSettlement{}
		for seat := 0; seat < playerNumber; seat++ {
			player := s.niuniuPlayers[seat]
			rdsop.SetTodayPlaying(player.ShortId)
			ntf.PlayerInfo = append(ntf.PlayerInfo, player.finalStatsMsg)
			player.finalStatsMsg = &outer.NiuNiuFinialPlayerInfo{}
		}
		settlementMsg.FinalSettlement = ntf
	}

	// 结算分数为最终金币
	modifyRspCount := make(map[string]struct{}) // 必须等待所有玩家金币修改成功后，才能发送结算
	for i := 0; i < playerNumber; i++ {
		seat := i
		player := s.niuniuPlayers[seat]
		finalScore := player.score
		presentScore := player.PlayerInfo.Gold
		s.room.Request(actortype.PlayerId(player.RID), &inner.ModifyGoldReq{
			Set:       true,
			Gold:      finalScore,
			SmallZero: true, // 允许扣为负数
		}).Handle(func(resp any, err error) {
			modifyRspCount[player.RID] = struct{}{}
			if err == nil {
				modifyRsp := resp.(*inner.ModifyGoldRsp)
				player.PlayerInfo = modifyRsp.Info
			}

			s.Log().Infow("modify gold result", "room", s.room.RoomId, "seat", seat,
				"player", player.ShortId, "latest gold", player.Gold, "err", err)

			// 记录本场游戏的输赢变化
			changes := finalScore - presentScore
			rdsop.SetUpdateGoldRecord(player.ShortId, rdsop.GoldUpdateReason{
				Type:      rdsop.GameWinOrLose, // 跑的快游戏输赢记录
				Gold:      changes,
				AfterGold: finalScore,
				OccurAt:   tools.Now(),
			})
			if len(modifyRspCount) == playerNumber {
				s.afterSettle(settlementMsg)
			}
		})
	}
}

func (s *StateSettlement) Leave() {
	s.Log().Infow("[NiuNiu] leave state settlement ==================SETTLEMENT==================", "room", s.room.RoomId)
	s.Log().Infof(" ")
	s.Log().Infof(" ")
	s.Log().Infof(" ")

}

func (s *StateSettlement) Handle(shortId int64, v any) (result any) {
	return outer.ERROR_NIUNIU_STATE_MSG_INVALID
}

func (s *StateSettlement) afterSettle(ntf *outer.NiuNiuSettlementNtf) {
	allPlayerInfo := s.playersToPB(0) // 组装结算消息

	for seat, player := range s.niuniuPlayers {
		ntf.PlayerData[seat].Player = allPlayerInfo[seat]
		ntf.PlayerData[seat].TotalScore = player.totalWinScore
	}

	s.Log().Infow(" settlement broadcast", "room", s.room.RoomId,
		"master", s.masterIndex, "ntf", ntf.String())

	s.room.Broadcast(ntf)

	// 结算给个短暂的时间
	s.room.AddTimer(tools.XUID(), s.currentStateEndAt, func(dt time.Duration) {
		s.SwitchTo(Ready)
	})
}

func (s *StateSettlement) finalSettlement() bool {
	return true
}

// profit 抽水
func (s *StateSettlement) profit(bigWinner bool) (totalProfit int64) {
	var (
		winners []*niuniuPlayer
	)

	if bigWinner {
		var (
			winScore int64
			winner   *niuniuPlayer
		)

		for i, player := range s.niuniuPlayers {
			if player.finalStatsMsg.TotalScore > winScore {
				winner = s.niuniuPlayers[i]
				winScore = player.finalStatsMsg.TotalScore
			}
		}

		winners = append(winners, winner)
	} else {
		for i, player := range s.niuniuPlayers {
			if player.finalStatsMsg.TotalScore > 0 {
				winners = append(winners, s.niuniuPlayers[i])
			}
		}
	}

	// 处理每一位赢家抽水
	for _, winner := range winners {
		rangeCfg := s.room.ProfitRange(winner.finalStatsMsg.TotalScore, s.gameParams().ReBate)
		winScore := winner.finalStatsMsg.TotalScore
		if rangeCfg == nil {
			s.Log().Warnw("winner profit score not in any range",
				"big winners", winner.ShortId, "win", winScore)
			continue
		}

		// 检查是否达到抽水最低要求
		minimumRebate := rangeCfg.MinimumRebate * common.Gold1000Times
		if winScore < minimumRebate {
			s.Log().Warnw("the winner win the score did not meet the expected score",
				"big winners", winner.ShortId, "win", winScore)
			continue
		}

		baseProfit := (winScore * rangeCfg.RebateRatio) / 100                     // 基础抽水
		profit := baseProfit + (rangeCfg.MinimumGuarantee * common.Gold1000Times) // +抽水保底值
		totalProfit += profit
		val := winner.score - profit
		winner.score = common.Max(0, val)
		s.Log().Infow("profit", "winner", winner.ShortId, "current score", winner.score,
			"baseProfit", baseProfit, "val", val, "winScore", winScore, "profit", profit, "range param", rangeCfg.String())
	}

	return
}
