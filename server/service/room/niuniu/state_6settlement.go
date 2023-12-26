package niuniu

import (
	"context"
	"time"

	"github.com/wwj31/dogactor/tools"

	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/common/rds"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/rdsop"
	"server/service/room"
)

// 结算状态

type StateSettlement struct {
	*NiuNiu
	history map[int64]int64
}

func (s *StateSettlement) State() int {
	return Settlement
}

func (s *StateSettlement) Enter() {
	s.history = make(map[int64]int64)
	s.lastMasterShort = s.niuniuPlayers[s.masterIndex].ShortId
	s.currentStateEndAt = tools.Now().Add(SettlementDuration).Add(time.Duration(s.participantCount()) * time.Second)
	s.Log().Infow("[NiuNiu] enter state settlement", "room", s.room.RoomId,
		"master", s.masterIndex, "endAt", s.currentStateEndAt.UnixMilli())

	s.settlementMsg = &outer.NiuNiuSettlementNtf{
		EndAt:            s.currentStateEndAt.UnixMilli(),
		GameSettlementAt: tools.Now().UnixMilli(),
		WinScores:        map[int32]int64{},
	}

	// 托管计数
	for i := 0; i < len(s.niuniuPlayers); i++ {
		player := s.niuniuPlayers[i]
		if player != nil && player.trusteeship {
			player.trusteeshipCount++
		}
	}

	// 先分别统计输家和赢家,排除庄家
	var (
		winners          []int
		losers           []int
		playerCardsTypes = map[int]CardsGroup{} // 闲家牌型，计算后加入，防止多次计算
	)

	master := s.niuniuPlayers[s.masterIndex]
	masterCardsType := s.niuniuPlayers[s.masterIndex].cardsGroup // 先拿到庄家牌型
	s.RangePartInPlayer(func(seat int, player *niuniuPlayer) {
		if seat == s.masterIndex {
			return
		}

		// 闲家赢加入winners,闲家输加入losers
		playerCardsType := player.handCards.AnalyzeCards(s.gameParams())
		playerCardsTypes[seat] = playerCardsType
		playerIsWin := playerCardsType.GreaterThan(masterCardsType)
		if playerIsWin {
			winners = append(winners, seat)
		} else {
			losers = append(losers, seat)
		}
		s.Log().Infow("compare niuniu",
			"master card", masterCardsType.String(),
			"player card", playerCardsType.String(), "player is win?", playerIsWin, "player shortId", player.ShortId, "player gold", player.score)
	})

	var cardTypeTimes map[PokerCardsType]int32
	if s.gameParams().NiuNiuTimes < 0 || s.gameParams().NiuNiuTimes > 1 {
		s.Log().Warnw(" niu niu params times err", "times", s.gameParams().NiuNiuTimes)
		cardTypeTimes = CardsTypeTimes[0]
	} else {
		cardTypeTimes = CardsTypeTimes[int(s.gameParams().NiuNiuTimes)]
	}

	// 赢家赢钱计算公式
	winFunc := func(seat int, winnerCardsGroup CardsGroup) (winScore int64) {
		return int64(s.masterTimesSeats[int32(s.masterIndex)]) * s.betGoldSeats[int32(seat)] * int64(cardTypeTimes[winnerCardsGroup.Type])
	}

	// 先计算庄家赢的钱
	for _, loserSeat := range losers {
		loser := s.niuniuPlayers[loserSeat]
		calcScore := winFunc(loserSeat, masterCardsType)
		winScore := common.Min(loser.score, calcScore)

		master.updateScore(winScore)
		s.niuniuPlayers[loserSeat].updateScore(-winScore)
		s.settlementMsg.WinScores[int32(loserSeat)] = s.niuniuPlayers[loserSeat].winScore
		s.Log().Infow("settle losers", "short", loser.ShortId, "win score", winScore, "calc score", calcScore)
	}

	// 统计庄家总共需要输的钱，如果够输直接算分，不够输就按照比例算分
	var totalMasterLoseScore int64
	winScores := map[int]int64{} // 每个位置分别需要赔多少钱
	for _, winSeat := range winners {
		winScore := winFunc(winSeat, playerCardsTypes[winSeat])
		totalMasterLoseScore += winScore
		winScores[winSeat] = winScore
		s.Log().Infow("settle winners", "win seat", winSeat)
	}
	s.Log().Infow("totalMasterLoseScore", "total", totalMasterLoseScore, "master score", master.score, "winScores", winScores)

	if master.score >= totalMasterLoseScore {
		for seat, score := range winScores {
			master.updateScore(-score)
			s.niuniuPlayers[seat].updateScore(score)
			s.settlementMsg.WinScores[int32(seat)] = s.niuniuPlayers[seat].winScore
		}
	} else {
		for seat, score := range winScores {
			winScore := master.score * score / totalMasterLoseScore
			master.updateScore(-winScore)
			s.niuniuPlayers[seat].updateScore(winScore)
			s.settlementMsg.WinScores[int32(seat)] = s.niuniuPlayers[seat].winScore
		}
	}
	// 所有输赢的闲家都统计完了，再把庄家的数据加入结算中
	s.settlementMsg.WinScores[int32(s.masterIndex)] = s.niuniuPlayers[s.masterIndex].winScore

	/////////////////////////////// 抽水返利 ///////////////////////////////
	s.profitAndRebate()

	///////////////////// 结算分数为最终金币 //////////////////////////////////
	var (
		count          int
		modifyRspCount = make(map[string]struct{}) // 必须等待所有玩家金币修改成功后，才能发送结算
	)

	s.RangePartInPlayer(func(seat int, player *niuniuPlayer) { count++ })

	var partInShortIds []int64
	s.RangePartInPlayer(func(seat int, player *niuniuPlayer) {
		finalScore := player.score
		presentScore := player.PlayerInfo.Gold
		changes := finalScore - presentScore
		partInShortIds = append(partInShortIds, player.ShortId)
		s.room.Request(actortype.PlayerId(player.RID), &inner.ModifyGoldReq{
			Set:       false,
			Gold:      finalScore,
			SmallZero: true, // 允许扣为负数
		}).Handle(func(resp any, err error) {
			modifyRspCount[player.RID] = struct{}{}
			if err == nil {
				modifyRsp := resp.(*inner.ModifyGoldRsp)
				player.PlayerInfo = modifyRsp.Info
			}

			s.Log().Infow("modify gold result", "room", s.room.RoomId, "seat", seat,
				"player", player.ShortId, "finalScore", finalScore, "presentScore", presentScore, "changes", changes, "err", err)

			// 记录本场游戏的输赢变化
			s.history[player.ShortId] = changes
			rdsop.SetUpdateGoldRecord(player.ShortId, rdsop.GoldUpdateReason{
				Type:      rdsop.GameWinOrLose, // 跑的快游戏输赢记录
				Gold:      changes,
				AfterGold: finalScore,
				OccurAt:   tools.Now(),
			})
			if len(modifyRspCount) == count {
				s.afterSettle(s.settlementMsg)
			}
		})
	})
	rdsop.SetTodayPlaying(partInShortIds...)
}

func (s *StateSettlement) Leave() {
	s.settlementMsg = nil
	s.Log().Infow("[NiuNiu] leave state settlement ==================SETTLEMENT==================", "room", s.room.RoomId)
	s.Log().Infof(" ")
	s.Log().Infof(" ")
	s.Log().Infof(" ")

}

func (s *StateSettlement) Handle(shortId int64, v any) (result any) {
	return outer.ERROR_NIUNIU_STATE_MSG_INVALID
}

func (s *StateSettlement) afterSettle(ntf *outer.NiuNiuSettlementNtf) {
	ntf.Players = s.playersToPB(0) // 组装结算消息

	s.Log().Infow(" settlement broadcast", "room", s.room.RoomId, "master", s.masterIndex, "ntf", ntf.String())

	s.room.Broadcast(ntf)

	s.room.GameRecordingOver(s.baseScore(), s.history)
	s.clear()

	// 结算给个短暂的时间
	s.room.AddTimer(tools.XUID(), s.currentStateEndAt, func(dt time.Duration) {
		s.SwitchTo(Ready)
	})
}

func (s *StateSettlement) finalSettlement() bool {
	return true
}

// profit 抽水和返利
func (s *StateSettlement) profitAndRebate() {
	// 庄家对每个人的赢分都单独计算抽水
	var (
		profitInfos []struct {
			profitScore int64
			winner      *niuniuPlayer
			loser       *niuniuPlayer
		}
	)

	master := s.niuniuPlayers[s.masterIndex]
	s.RangePartInPlayer(func(seat int, player *niuniuPlayer) {
		if seat == s.masterIndex {
			return
		}

		profitInfo := struct {
			profitScore int64
			winner      *niuniuPlayer
			loser       *niuniuPlayer
		}{
			profitScore: player.winScore,
			winner:      nil,
			loser:       nil,
		}

		// 记录每一笔的赢家和输家
		if player.winScore > 0 {
			profitInfo.winner = player
			profitInfo.loser = master
		} else {
			profitInfo.winner = master
			profitInfo.loser = player
		}
		profitInfos = append(profitInfos, profitInfo)
	})

	// 记录返利信息
	record := &outer.RebateDetailInfo{
		Type:      outer.GameType_NiuNiu,
		BaseScore: s.gameParams().BaseScore,
		CreateAt:  tools.Now().UnixMilli(),
	}
	pip := rds.Ins.Pipeline()

	// 处理每一位赢家抽水
	for _, profit := range profitInfos {
		rangeCfg := s.room.ProfitRange(profit.profitScore, s.gameParams().ReBate)
		winScore := profit.profitScore
		if rangeCfg == nil {
			s.Log().Warnw("winner profit score not in any range", "profitScore", profit.profitScore, "rebate", s.gameParams().ReBate.String())
			continue
		}

		// 检查是否达到抽水最低要求
		minimumRebate := rangeCfg.MinimumRebate * common.Gold1000Times
		if winScore < minimumRebate {
			s.Log().Warnw("the winner win the score did not meet the expected score",
				"winner", profit.winner.ShortId, "win score", winScore, "minimumRebate", minimumRebate)
			continue
		}

		baseProfit := (winScore * rangeCfg.RebateRatio) / 100                        // 基础抽水
		profitVal := baseProfit + (rangeCfg.MinimumGuarantee * common.Gold1000Times) // +抽水保底值
		profit.winner.score = common.Max(0, profit.winner.score-profitVal)
		s.Log().Infow("profit", "winner", profit.winner.ShortId, "current score", profit.winner.score,
			"baseProfit", baseProfit, "profitVal", profitVal, "winScore", winScore, "profit", profit, "range param", rangeCfg.String())

		s.room.Rebate(record, profitVal, []*room.Player{profit.winner.Player, profit.loser.Player}, pip)
	}

	if _, err := pip.Exec(context.Background()); err != nil {
		log.Errorw("rebate redis failed", "err", err)
	}

	return
}
