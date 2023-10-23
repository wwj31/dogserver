package niuniu

import (
	"time"

	"github.com/wwj31/dogactor/tools"

	"server/common"
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
	s.lastMasterShort = s.niuniuPlayers[s.masterIndex].ShortId
	s.currentStateEndAt = tools.Now().Add(SettlementDuration)
	s.Log().Infow("[NiuNiu] enter state settlement", "room", s.room.RoomId,
		"master", s.masterIndex, "endAt", s.currentStateEndAt.UnixMilli())

	settlementMsg := &outer.NiuNiuSettlementNtf{
		EndAt:            s.currentStateEndAt.UnixMilli(),
		GameSettlementAt: tools.Now().UnixMilli(),
		PlayerData:       make(map[int32]*outer.NiuNiuSettlementPlayerData),
	}

	s.RangePartInPlayer(func(seat int, player *niuniuPlayer) {
		settlementMsg.PlayerData[int32(seat)] = &outer.NiuNiuSettlementPlayerData{}
	})

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

	s.RangePartInPlayer(func(seat int, player *niuniuPlayer) {
		ntf.PlayerData[int32(seat)].Player = allPlayerInfo[seat]
		ntf.PlayerData[int32(seat)].TotalScore = player.winScore
	})

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
func (s *StateSettlement) profit() (totalProfit int64) {
	var (
		winners []*niuniuPlayer
	)

	for i, player := range s.niuniuPlayers {
		if player.winScore > 0 {
			winners = append(winners, s.niuniuPlayers[i])
		}
	}

	// 处理每一位赢家抽水
	for _, winner := range winners {
		rangeCfg := s.room.ProfitRange(winner.winScore, s.gameParams().ReBate)
		winScore := winner.winScore
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
