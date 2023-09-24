package mahjong

import (
	"time"

	"github.com/wwj31/dogactor/tools"

	"server/common"
	"server/common/actortype"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/rdsop"
)

// 结算状态

type StateSettlement struct {
	*Mahjong
}

func (s *StateSettlement) State() int {
	return Settlement
}

func (s *StateSettlement) Enter() {
	notHu := s.isNoHu()

	s.currentStateEndAt = tools.Now().Add(SettlementDuration)
	s.Log().Infow("[Mahjong] enter state settlement",
		"room", s.room.RoomId, "master", s.masterIndex, "notHu", notHu, "game count", s.gameCount,
		"endAt", s.currentStateEndAt.UnixMilli())

	settlementMsg := &outer.MahjongBTESettlementNtf{
		EndAt:            s.currentStateEndAt.UnixMilli(),
		NotHu:            notHu,
		HasScoreZero:     s.scoreZeroOver,
		GameCount:        int32(s.gameCount),
		GameSettlementAt: tools.Now().UnixMilli(),
		HuSeatIndex:      s.huSeat,
		PlayerData:       make([]*outer.MahjongBTESettlementPlayerData, maxNum, maxNum),
	}

	for seat := 0; seat < maxNum; seat++ {
		settlementMsg.PlayerData[seat] = &outer.MahjongBTESettlementPlayerData{}
	}

	// 流局
	if notHu {
		s.notHu(settlementMsg)
		for _, player := range s.mahjongPlayers {
			if player.score <= 0 {
				s.scoreZeroOver = true
			}
		}
	}

	// 大结算
	if s.finalSettlement() || s.scoreZeroOver {
		s.Log().Infow("final settlement",
			"scoreZeroOver", s.scoreZeroOver, "game count", s.gameCount, "param", s.gameParams().PlayCountLimit)
		s.gameCount = int(s.gameParams().PlayCountLimit)

		// 总抽水
		totalProfit := s.profit(s.gameParams().BigWinner)
		s.room.Rebate(totalProfit, s.toRoomPlayers())

		ntf := &outer.MahjongBTEFinialSettlement{}
		for seat := 0; seat < maxNum; seat++ {
			player := s.mahjongPlayers[seat]
			ntf.PlayerInfo = append(ntf.PlayerInfo, player.finalStatsMsg)
			player.finalStatsMsg = &outer.MahjongBTEFinialPlayerInfo{}
		}
		settlementMsg.FinalSettlement = ntf

	}

	// 结算分数为最终金币
	modifyRspCount := make(map[string]struct{}) // 必须等待所有玩家金币修改成功后，才能发送结算
	for i := 0; i < maxNum; i++ {
		seat := i
		player := s.mahjongPlayers[seat]
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
				Type: rdsop.GameWinOrLose, // 血战麻将游戏输赢记录
				Gold: changes,
			})

			if len(modifyRspCount) == maxNum {
				s.afterSettle(settlementMsg)
			}
		})
	}

}

func (s *StateSettlement) Leave() {
	s.Log().Infow("[Mahjong] leave state settlement ==================SETTLEMENT==================", "room", s.room.RoomId, "count", s.gameCount)
	s.Log().Infof(" ")
	s.Log().Infof(" ")
	s.Log().Infof(" ")

}

func (s *StateSettlement) Handle(shortId int64, v any) (result any) {
	return outer.ERROR_MAHJONG_STATE_MSG_INVALID
}

// 流局结算
func (s *StateSettlement) notHu(ntf *outer.MahjongBTESettlementNtf) {
	// 筛选出有叫、没叫、花猪的玩家
	var (
		hasTingSeat, hasNotTingSeat, pigSeats []int
	)

	allTingCards := make(map[int]map[Card]HuType)

	for seat, player := range s.mahjongPlayers {
		tingCards, err := player.handCards.ting(player.ignoreColor,
			player.lightGang,
			player.darkGang,
			player.pong,
			s.gameParams(),
		)

		if err != nil {
			s.Log().Errorw("player ting error",
				"room", s.room.RoomId, "player", player.ShortId, "seat", seat, "hand", player.handCards, "err", err)
			continue
		}

		allTingCards[seat] = tingCards

		if len(tingCards) > 0 {
			hasTingSeat = append(hasTingSeat, seat)
		} else {
			// 没有叫，需要退杠分
			for _, gang := range player.gangInfos {
				for loseSeat, score := range gang.loserSeats {
					s.mahjongPlayers[loseSeat].gangTotalScore += score // 退杠得分
					s.mahjongPlayers[loseSeat].updateScore(score)

					player.gangTotalScore -= score // 没叫，退杠分
					player.updateScore(-score)
				}
			}

			// 花猪单独赔钱，不计入查大叫
			if player.handCards.HasColorCard(player.ignoreColor) {
				pigSeats = append(pigSeats, seat) // 花猪
				ntf.PigSeat = append(ntf.PigSeat, int32(seat))
			} else {
				hasNotTingSeat = append(hasNotTingSeat, seat) // 没叫的人
				ntf.NotTingSeat = append(ntf.NotTingSeat, int32(seat))
			}
		}
	}

	// 4家全是花猪, 或者4家都没有叫  ？？?
	if len(pigSeats) == maxNum || len(hasNotTingSeat) == maxNum {
		s.Log().Infow("all player is the pig", "pigs", pigSeats, "not ting", hasNotTingSeat)
		return
	}

	baseScore := s.baseScore()

	// 先把猪儿的钱赔了
	if len(pigSeats) > 0 {
		winSeats := s.allSeats(pigSeats...) // 非花猪的位置
		s.Log().Infow("pig seats", "seats", pigSeats)

		for _, pigSeat := range pigSeats {
			// 先算这个猪儿需要赔多少分
			playerPig := s.mahjongPlayers[pigSeat]
			fan := common.Min(6, int(s.fanUpLimit()))
			winScore := s.fanScore(fan, baseScore)

			// 猪儿已经没钱，赔不了了
			if playerPig.score <= 0 {
				continue
			}

			// 如果猪儿有钱不够赔，
			// 那么有多少就赔多少，赢的人均分
			if !s.gameParams().AllowScoreSmallZero {
				winScore = common.Min(playerPig.score, winScore)
			}

			totalLoseScore := winScore * int64(len(winSeats))
			playerPig.updateScore(-totalLoseScore)
			s.Log().Infow("pig!", "pig seat", pigSeat, "totalLoseScore", totalLoseScore, "winner seats", winSeats)
			// 赢钱的人加分
			for _, seat := range winSeats {
				player := s.mahjongPlayers[seat]
				player.updateScore(winScore)
			}
		}
	}

	// 查大叫
	if len(hasNotTingSeat) > 0 {
		var (
			allWinner     = make(map[int]int64) // 赢钱的位置和需要赢的分
			totalWinScore int64                 // 总赔付
		)

		for _, tingSeat := range hasTingSeat {
			// 算出听牌可胡牌的最大番+根
			tingCards := allTingCards[tingSeat]
			card, maxFan := s.maxFanTingCard(tingCards)
			s.mahjongPlayers[tingSeat].handCards = s.mahjongPlayers[tingSeat].handCards.Insert(card)
			gen := int32(s.huGen(tingSeat))
			maxFan += gen
			maxFan = common.Min(maxFan, s.fanUpLimit())
			winScore := s.fanScore(int(maxFan), baseScore)
			allWinner[tingSeat] = winScore
			totalWinScore += winScore
			s.Log().Infow("notHu ting",
				"ting seat", tingSeat, "tingCards", tingCards, "base score", baseScore, "card", card,
				"max fan", maxFan, "gen", gen, "win score", winScore)
		}

		// 没叫的挨个赔有叫的
		for _, notTingSeat := range hasNotTingSeat {
			// 够赔，或者允许负分，直接赔
			if s.gameParams().AllowScoreSmallZero || s.mahjongPlayers[notTingSeat].score < totalWinScore {
				for seat, winScore := range allWinner {
					s.mahjongPlayers[seat].updateScore(winScore)
					s.mahjongPlayers[notTingSeat].updateScore(-winScore)
					s.Log().Infow("notHu notTing 1",
						"notTingSeat", notTingSeat, "seat", seat, "winScore", winScore)
				}
			} else {
				// 不够赔，并且不允许负分，就按照比例赔付
				loserScore := s.mahjongPlayers[notTingSeat].score
				for seat, winScore := range allWinner {
					exactScore := loserScore * winScore / totalWinScore
					s.mahjongPlayers[seat].updateScore(exactScore)
					s.mahjongPlayers[notTingSeat].updateScore(-exactScore)
					s.Log().Infow("notHu notTing 2",
						"notTingSeat", notTingSeat, "seat", seat, "winScore", winScore, "totalWinScore", totalWinScore, "exactScore", exactScore)
				}
			}
		}
	}
}

// 从听牌能胡的牌型中，选出最大番
func (s *StateSettlement) maxFanTingCard(tingCards map[Card]HuType) (card Card, maxBaseFan int32) {
	for c, huType := range tingCards {
		if maxBaseFan == 0 || int32(huFan[huType]) > maxBaseFan {
			maxBaseFan = int32(huFan[huType])
			card = c
		}
	}

	return
}

func (s *StateSettlement) afterSettle(ntf *outer.MahjongBTESettlementNtf) {
	allPlayerInfo := s.playersToPB(0, true) // 组装结算消息

	darkGangMap := map[int32]map[int32]int64{}  // 表示每个人被哪些位置暗杠过
	lightGangMap := map[int32]map[int32]int64{} // 表示每个人被哪些位置明杠过

	// 根据每个人的杠牌，先分析出每个人扣的杠分
	for seat, player := range s.mahjongPlayers {
		var huPeer peerRecords
		if player.huPeerIndex != -1 {
			huPeer = s.peerRecords[player.huPeerIndex]
		}

		totalFan := int32(huFan[player.hu]) + player.huGen
		for _, ext := range player.huExtra {
			totalFan += int32(extraFan[ext])
		}
		ntf.PlayerData[seat].Player = allPlayerInfo[seat]
		ntf.PlayerData[seat].DianPaoSeatIndex = int32(huPeer.seat)
		ntf.PlayerData[seat].TotalFan = totalFan
		ntf.PlayerData[seat].TotalScore = player.totalWinScore
		ntf.PlayerData[seat].HuWinScoreSeatIndex = player.winScore

		// 本局该玩家所有的杠牌,以及每次杠成功后赔钱的位置
		for peerIndex, info := range player.gangInfos {
			for loserSeat, loserWin := range info.loserSeats {
				if s.peerRecords[peerIndex].typ == GangType4 {
					if darkGangMap[loserSeat] == nil {
						darkGangMap[loserSeat] = make(map[int32]int64)
					}
					darkGangMap[loserSeat][int32(seat)] += loserWin
				} else {
					if lightGangMap[loserSeat] == nil {
						lightGangMap[loserSeat] = make(map[int32]int64)
					}
					lightGangMap[loserSeat][int32(seat)] += loserWin
				}
			}
		}
	}

	// 组装暗杠数据
	for seat, gangSeat := range darkGangMap {
		ntf.PlayerData[seat].ByDarkGangSeatIndex = gangSeat
	}

	// 组装明杠数据
	for seat, gangSeat := range lightGangMap {
		ntf.PlayerData[seat].ByLightGangSeatIndex = gangSeat
	}

	s.Log().Infow(" settlement broadcast",
		"room", s.room.RoomId, "master", s.masterIndex, "ntf", ntf.String())
	s.room.Broadcast(ntf)

	s.clear()           // 分算完清理数据
	s.nextMasterIndex() // 计算下一局庄家

	// 结算给个短暂的时间
	s.room.AddTimer(tools.XUID(), s.currentStateEndAt, func(dt time.Duration) {
		s.SwitchTo(Ready)
	})
}

func (s *StateSettlement) nextMasterIndex() {
	if s.multiHuByIndex != -1 {
		s.masterIndex = s.multiHuByIndex
	} else if len(s.huSeat) > 0 {
		s.masterIndex = int(s.huSeat[0])
	} else {
		s.masterIndex = s.nextSeatIndexWithoutHu(s.masterIndex)
	}
}

// 检查是否流局
func (s *StateSettlement) isNoHu() bool {
	// 如果还有牌进入结算，要么是3家胡了，要么是有玩家输光强制结算
	if s.cards.Len() > 0 {
		return false
	}

	for _, player := range s.mahjongPlayers {
		if player.hu != HuInvalid {
			return false
		}
	}
	return true
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
		winners []*mahjongPlayer
	)

	if bigWinner {
		var (
			winScore int64
			winner   *mahjongPlayer
		)

		for i, player := range s.mahjongPlayers {
			if player.finalStatsMsg.TotalScore > winScore {
				winner = s.mahjongPlayers[i]
				winScore = player.finalStatsMsg.TotalScore
			}
		}

		if winners != nil {
			winners = append(winners, winner)
		}
	} else {
		for i, player := range s.mahjongPlayers {
			if player.finalStatsMsg.TotalScore > 0 {
				winners = append(winners, s.mahjongPlayers[i])
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
