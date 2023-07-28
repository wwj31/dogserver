package mahjong

import (
	"math"
	"time"

	"server/common"
	"server/common/actortype"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"

	"github.com/wwj31/dogactor/tools"
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
	s.gameCount++
	s.Log().Infow("[Mahjong] enter state settlement",
		"room", s.room.RoomId, "master", s.masterIndex, "notHu", notHu, "game count", s.gameCount)

	settlementMsg := &outer.MahjongBTESettlementNtf{
		NotHu:            notHu,
		GameCount:        int32(s.gameCount),
		GameSettlementAt: tools.Now().UnixMilli(),
		HuSeatIndex:      s.huSeat,
		PlayerData:       make([]*outer.MahjongBTESettlementPlayerData, maxNum, maxNum),
	}

	for seat := 0; seat < maxNum; seat++ {
		settlementMsg.PlayerData[seat] = &outer.MahjongBTESettlementPlayerData{}
	}

	if notHu {
		s.notHu(settlementMsg)
	} else {

	}

	// 结算分数为最终金币
	modifyRspCount := make(map[string]struct{}) // 必须等待所有玩家金币修改成功后，才能发送结算
	for seat, player := range s.mahjongPlayers {
		rid := player.RID
		finalScore := common.Max(0, player.score)
		s.room.Request(actortype.PlayerId(rid), &inner.ModifyGoldReq{Gold: finalScore}).Handle(func(resp any, err error) {
			modifyRspCount[rid] = struct{}{}
			if err != nil {
				s.Log().Errorw("modify gold failed kick-out player",
					"room", s.room.RoomId, "player", player.ShortId, "seat", seat, "err", err)
			} else {
				modifyRsp := resp.(*inner.ModifyGoldRsp)
				player.PlayerInfo = modifyRsp.Info
			}
			s.Log().Infow("modify gold success", "room", s.room.RoomId, "player", player.ShortId, "seat", seat, "player info", *player.PlayerInfo)
			if len(modifyRspCount) == maxNum {
				s.settlementBroadcast(settlementMsg)
			}
		})
	}
}

func (s *StateSettlement) Leave() {
	s.Log().Infow("[Mahjong] leave state settlement", "room", s.room.RoomId)
}

func (s *StateSettlement) Handle(shortId int64, v any) (result any) {
	return outer.ERROR_MAHJONG_STATE_MSG_INVALID
}

// 流局结算
func (s *StateSettlement) notHu(ntf *outer.MahjongBTESettlementNtf) {
	// 筛选出有叫、没叫、花猪的玩家
	var (
		hasTingSeat, hasNotTingSeat, pig []int
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
			for peerIdx, losers := range player.gangScore {
				peer := s.peerRecords[peerIdx]
				loseScore := int64(float32(s.baseScore()) * gangScoreRatio[peer.typ])
				for _, loserSeatIndex := range losers {
					s.mahjongPlayers[loserSeatIndex].score += loseScore
					player.score -= loseScore
				}
			}

			// 花猪单独赔钱，不计入查大叫
			if player.handCards.HasColorCard(player.ignoreColor) {
				pig = append(pig, seat) // 花猪
			} else {
				hasNotTingSeat = append(hasNotTingSeat, seat) // 没叫的人
			}
		}
	}

	// 先把猪儿的钱赔了
	if len(pig) > 0 {
		winPigSeats := s.allSeats(pig...) // 非花猪的位置
		var winPigScore int64             // 非花猪的赢分
		for _, pigSeat := range winPigSeats {
			// 先算这个猪儿需要赔多少分
			playerPig := s.mahjongPlayers[pigSeat]
			needSubBaseScore := int64(math.Pow(float64(s.baseScore()), float64(s.fanUpLimit())))

			// 猪儿已经没钱，赔不了了
			if playerPig.score <= 0 {
				continue
			}

			// 如果猪儿有钱不够赔，那么有多少就赔多少，赢的人均分
			if playerPig.score < needSubBaseScore*int64(len(winPigSeats)) {
				winPigScore = playerPig.score / int64(len(winPigSeats))
			} else {
				winPigScore = needSubBaseScore
			}
			playerPig.score -= winPigScore * int64(len(winPigSeats))

			// 赢钱的人加分
			for _, seat := range winPigSeats {
				player := s.mahjongPlayers[seat]
				player.score += winPigScore
			}
		}
	}

	// 查大叫
	if len(hasNotTingSeat) > 0 {
		allWinner := make(map[int]int64) // 赢钱的位置和需要赢的分
		var totalWinScore int64          // 总赔付
		for _, tingSeat := range hasTingSeat {
			// 算出听牌可胡牌的最大番
			tingCards := allTingCards[tingSeat]
			maxFan := s.maxFanTingCard(tingCards)
			maxFan = common.Min(maxFan, s.fanUpLimit())
			winScore := int64(math.Pow(float64(s.baseScore()), float64(s.fanUpLimit())))
			allWinner[tingSeat] = winScore
			totalWinScore += winScore
		}

		// 没叫的挨个赔钱
		for _, notTingSeat := range hasNotTingSeat {
			notTingPlayer := s.mahjongPlayers[notTingSeat]
			if notTingPlayer.score <= 0 {
				continue // 没钱赔个屁
			}

			// 如果不够赔,就按照赔付比例赔付给有叫的人
			if notTingPlayer.score < totalWinScore {
				for winSeat, _ := range allWinner {
					divScore := float64(notTingPlayer.score) * (float64(allWinner[winSeat]) / float64(totalWinScore))
					s.mahjongPlayers[winSeat].score += int64(divScore)
				}
				notTingPlayer.score = 0
			} else {
				// 够赔就直接赔
				for seat, winScore := range allWinner {
					s.mahjongPlayers[seat].score += winScore
				}
				notTingPlayer.score -= totalWinScore
			}
		}
	}

}

// 从听牌能胡的牌型中，选出最大番
func (s *StateSettlement) maxFanTingCard(tingCards map[Card]HuType) int32 {
	var maxFan int32
	for _, huType := range tingCards {
		if maxFan == 0 || int32(huFan[huType]) > maxFan {
			maxFan = int32(huFan[huType])
		}
	}

	return maxFan
}

func (s *StateSettlement) settlementBroadcast(ntf *outer.MahjongBTESettlementNtf) {
	// 组装结算消息
	allPlayerInfo := s.playersToPB(0, true)

	darkGangMap := map[int32][]int32{}  // 表示每个人被哪些位置暗杠过
	lightGangMap := map[int32][]int32{} // 表示每个人被哪些位置明杠过

	// 根据每个人的杠牌，先分析出每个人扣的杠分
	for seat, player := range s.mahjongPlayers {
		var huPeer peerRecords
		if player.huPeerIndex != -1 {
			huPeer = s.peerRecords[player.huPeerIndex]
		}

		ntf.PlayerData[seat].Player = allPlayerInfo[seat]
		ntf.PlayerData[seat].DianPaoSeatIndex = int32(huPeer.seat)
		ntf.PlayerData[seat].TotalFan = int32(huFan[player.hu]+extraFan[player.huExtra]) + player.huGen
		ntf.PlayerData[seat].TotalScore = player.score

		// 本局该玩家所有的杠牌,以及每次杠成功后赔钱的位置
		for peerIndex, loserSeats := range player.gangScore {
			for _, loserSeat := range loserSeats {
				if s.peerRecords[peerIndex].typ == GangType4 {
					darkGangMap[loserSeat] = append(darkGangMap[loserSeat], int32(seat))
				} else {
					lightGangMap[loserSeat] = append(lightGangMap[loserSeat], int32(seat))
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

	s.Log().Infow(" settlement broadcast", "room", s.room.RoomId, "master", s.masterIndex, "ntf", ntf.String())
	s.room.Broadcast(ntf)

	s.clear()           // 分算完清理数据
	s.nextMasterIndex() // 计算下一局庄家

	// 结算给个短暂的时间
	s.currentStateEndAt = tools.Now().Add(SettlementDuration)
	s.room.AddTimer(tools.XUID(), s.currentStateEndAt, func(dt time.Duration) {
		s.SwitchTo(Ready)
	})
}
func (s *StateSettlement) nextMasterIndex() {
	if s.mutilHuByIndex != -1 {
		s.masterIndex = s.mutilHuByIndex
	} else if len(s.huSeat) > 0 {
		s.masterIndex = int(s.huSeat[0])
	} else {
		s.masterIndex = s.nextSeatIndexWithoutHu(s.masterIndex)
	}
}

// 检查是否流局
func (s *StateSettlement) isNoHu() bool {
	for _, player := range s.mahjongPlayers {
		if player.hu != HuInvalid {
			return false
		}
	}
	return true
}
