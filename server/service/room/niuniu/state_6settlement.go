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
		WinScores:        map[int32]int64{},
		CardsTypes:       map[int32]*outer.NiuNiuCardsGroup{},
	}

	// 先分别统计输家和赢家,排除庄家
	var (
		winners          []int
		losers           []int
		playerCardsTypes = map[int]CardsGroup{} // 闲家牌型，计算后加入，防止多次计算
	)

	master := s.niuniuPlayers[s.masterIndex]
	masterCardsType := s.niuniuPlayers[s.masterIndex].handCards.AnalyzeCards(s.gameParams()) // 先拿到庄家牌型
	settlementMsg.CardsTypes[int32(s.masterIndex)] = masterCardsType.ToPB()
	s.RangePartInPlayer(func(seat int, player *niuniuPlayer) {
		if seat == s.masterIndex {
			return
		}

		// 闲家赢加入winners,闲家输加入losers
		playerCardsType := player.handCards.AnalyzeCards(s.gameParams())
		playerCardsTypes[seat] = playerCardsType
		settlementMsg.CardsTypes[int32(seat)] = playerCardsType.ToPB()
		if playerCardsType.GreaterThan(masterCardsType) {
			winners = append(winners, seat)
		} else {
			losers = append(losers, seat)
		}
	})

	var cardTypeTimes map[PokerCardsType]int32
	if s.gameParams().NiuNiuTimes < 0 || s.gameParams().NiuNiuTimes > 1 {
		s.Log().Errorw(" niu niu params times err")
		cardTypeTimes = CardsTypeTimes[0]
	} else {
		cardTypeTimes = CardsTypeTimes[int(s.gameParams().NiuNiuTimes)]
	}

	// 赢家赢钱计算公式
	winFunc := func(loseSeat int, winnerCardsGroup CardsGroup) (winScore int64) {
		return int64(s.masterTimesSeats[int32(s.masterIndex)]) * s.betGoldSeats[int32(loseSeat)] * int64(cardTypeTimes[winnerCardsGroup.Type])
	}

	// 先计算庄家赢的钱
	for _, loserSeat := range losers {
		loser := s.niuniuPlayers[loserSeat]
		winScore := winFunc(loserSeat, masterCardsType)
		winScore = common.Min(loser.score, winScore)

		master.updateScore(winScore)
		s.niuniuPlayers[loserSeat].updateScore(-winScore)
		settlementMsg.WinScores[int32(loserSeat)] = s.niuniuPlayers[loserSeat].winScore
	}

	// 统计庄家总共需要输的钱，如果够输直接算分，不够输就按照比例算分
	var totalMasterLoseScore int64
	winScores := map[int]int64{} // 每个位置分别需要赔多少钱
	for _, winSeat := range winners {
		winScore := winFunc(s.masterIndex, playerCardsTypes[winSeat])
		totalMasterLoseScore += winScore
		winScores[winSeat] = winScore
	}

	if master.score >= totalMasterLoseScore {
		for seat, score := range winScores {
			master.updateScore(-score)
			s.niuniuPlayers[seat].updateScore(score)
			settlementMsg.WinScores[int32(seat)] = s.niuniuPlayers[seat].winScore
		}
	} else {
		for seat, score := range winScores {
			winScore := master.score * score / totalMasterLoseScore
			master.updateScore(-winScore)
			s.niuniuPlayers[seat].updateScore(winScore)
			settlementMsg.WinScores[int32(seat)] = s.niuniuPlayers[seat].winScore
		}
	}

	/////////////////////////////// 抽水返利 ///////////////////////////////
	s.profitAndRebate()

	///////////////////// 结算分数为最终金币 //////////////////////////////////
	modifyRspCount := make(map[string]struct{}) // 必须等待所有玩家金币修改成功后，才能发送结算
	count := int(s.playerCount())
	var partInShortIds []int64
	s.RangePartInPlayer(func(seat int, player *niuniuPlayer) {
		finalScore := player.score
		presentScore := player.PlayerInfo.Gold
		partInShortIds = append(partInShortIds, player.ShortId)
		s.room.Request(actortype.PlayerId(player.RID), &inner.ModifyGoldReq{
			Set:  true,
			Gold: finalScore,
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
			if len(modifyRspCount) == count {
				s.afterSettle(settlementMsg)
			}
		})
	})
	rdsop.SetTodayPlaying(partInShortIds...)
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
	ntf.Players = s.playersToPB(0) // 组装结算消息

	s.Log().Infow(" settlement broadcast", "room", s.room.RoomId, "master", s.masterIndex, "ntf", ntf.String())

	s.room.Broadcast(ntf)

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
