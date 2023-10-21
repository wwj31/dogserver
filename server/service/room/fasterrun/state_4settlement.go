package fasterrun

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
		SpareCards:       s.spareCards.ToPB(),
	}

	for seat := 0; seat < playerNumber; seat++ {
		settlementMsg.PlayerData[seat] = &outer.FasterRunSettlementPlayerData{}
	}

	// 结算输赢分
	var (
		winner     *fasterRunPlayer
		winnerSeat int
	)
	for seat, player := range s.fasterRunPlayers {
		if len(player.handCards) == 0 {
			winner = player
			winnerSeat = seat
			break
		}
	}

	if winner != nil {
		// 春天、反春判断
		isSpring := s.isSpring()
		var isAgainstSpring bool
		if !isSpring && s.gameParams().AgainstSpring {
			isAgainstSpring = s.isAgainstSpring(winner)
		}
		s.Log().Infow("againstspring ",
			"isAgainstSpring", isAgainstSpring, "isSpring", isSpring, "s.gameParams().AgainstSpring", s.gameParams().AgainstSpring, "s.playRecords", s.playRecords)
		settlementMsg.Spring = isSpring
		settlementMsg.AgainstSpring = isAgainstSpring

		loserSeats := s.allSeats(winnerSeat)
		for _, seat := range loserSeats {
			loser := s.fasterRunPlayers[seat]
			loseScore := int64(len(loser.handCards)) * s.baseScore()

			// 特殊规则剩一张不输
			if s.gameParams().SpareOnlyOneWithoutLose && len(loser.handCards) == 1 {
				s.Log().Infow("SpareOnlyOneWithoutLose ", "shortId", loser.ShortId)
				continue
			}

			// 春天和反春的算分
			if isSpring {
				loseScore *= 2
			} else if isAgainstSpring {
				// 反春，重新算分
				if seat == s.masterIndex {
					needNum := 15
					if s.gameParams().CardsNumber == 1 {
						needNum = 16
					}
					loseScore = int64(needNum) * s.baseScore()
				}
				loseScore *= 2
			}

			// 红桃10，输赢翻倍
			if winner.doubleHearts10 {
				loseScore *= 2
			} else if loser.doubleHearts10 {
				loseScore *= 2
			}

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
			"scoreZeroOver", s.scoreZeroOver, "game count", s.gameCount, "param", s.gameParams().PlayCountLimit)
		s.gameCount = int(s.gameParams().PlayCountLimit)

		// 总抽水
		totalProfit := s.profit(s.gameParams().BigWinner)

		// 记录返利信息
		record := &outer.RebateDetailInfo{
			Type:      outer.GameType_FasterRun,
			BaseScore: s.gameParams().BaseScore,
			CreateAt:  tools.Now().UnixMilli(),
		}
		s.room.Rebate(record, totalProfit, s.toRoomPlayers())

		ntf := &outer.FasterRunFinialSettlement{}
		for seat := 0; seat < playerNumber; seat++ {
			player := s.fasterRunPlayers[seat]
			rdsop.SetTodayPlaying(player.ShortId)
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

// 春天规则
func (s *StateSettlement) isSpring() bool {
	master := s.fasterRunPlayers[s.masterIndex]
	if len(master.handCards) > 0 {
		return false
	}

	var fullNum int
	if s.gameParams().CardsNumber == 0 {
		fullNum = 15
	} else {
		fullNum = 16
	}

	others := s.allSeats(s.masterIndex)
	for _, seat := range others {
		other := s.fasterRunPlayers[seat]
		if len(other.handCards) != fullNum {
			return false
		}
	}

	return true
}

// 反春规则
func (s *StateSettlement) isAgainstSpring(winner *fasterRunPlayer) bool {
	// 除了第一手是庄稼，其他所有出牌全部是赢的这家，肯定就是反春
	for i := 1; i < len(s.playRecords); i++ {
		if s.playRecords[i].shortId != winner.ShortId {
			return false
		}
	}
	return true
}

func (s *StateSettlement) Leave() {
	s.Log().Infow("[FasterRun] leave state settlement ==================SETTLEMENT==================", "room", s.room.RoomId, "count", s.gameCount)
	s.Log().Infof(" ")
	s.Log().Infof(" ")
	s.Log().Infof(" ")

}

func (s *StateSettlement) Handle(shortId int64, v any) (result any) {
	return outer.ERROR_FASTERRUN_STATE_MSG_INVALID
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

		if winner != nil {
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
