package mahjong

import (
	"server/proto/outermsg/outer"
)

var gangScoreRatio = map[checkCardType]float32{
	GangType1: 1, // 弯杠1分
	GangType3: 2, // 弯杠2分
	GangType4: 2, // 暗杠2分
}

// 杠牌操作
func (s *StatePlaying) operateGang(p *mahjongPlayer, seatIndex int, card Card, ntf *outer.MahjongBTEOperaNtf) (ok bool, err outer.ERROR) {
	if len(s.peerRecords) == 0 {
		s.Log().Errorw("operate gang failed peerRecords len = 0",
			"room", s.room.RoomId, "player", p.ShortId)
		return false, outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	if !s.currentAction.canGang(card) {
		s.Log().Errorw("operate gang failed invalid gang card",
			"room", s.room.RoomId, "player", p.ShortId, "currentGang", s.currentAction.gang, "card", card)
		return false, outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	// 获得最后一次操作的牌
	lastPeerIndex := len(s.peerRecords) - 1
	peer := s.peerRecords[lastPeerIndex]

	var (
		qiangGang  bool
		loseScores map[int32]int64
		gangFunc   func(opNtf *outer.MahjongBTEOperaNtf)
		gangType   checkCardType
	)

	// 检查能否被抢杠胡
	hasQiangGang := func() bool {
		b := false
		for seat, other := range s.mahjongPlayers {
			if seatIndex == seat {
				continue
			}

			if hu := other.handCards.Insert(card).IsHu(other.lightGang, other.darkGang, other.pong, card, s.gameParams()); hu != HuInvalid {
				newAction := action{}
				newAction.acts = append(newAction.acts, outer.ActionType_ActionHu, outer.ActionType_ActionPass)
				newAction.hus = append(newAction.hus, hu.PB())
				s.actionMap[seat] = &newAction // 抢杠胡操作
				b = true
			}
		}
		return b
	}

	// 统一计算赔付分
	loseScoreAnalyze := func(seats ...int) map[int32]int64 {
		result := make(map[int32]int64)
		loseScore := int64(float32(s.baseScore()) * gangScoreRatio[gangType])
		for _, seatIdx := range seats {
			result[int32(seatIdx)] = loseScore
		}
		return result
	}

	// 杠成功后实时算分
	execScoreFunc := func(opNtf *outer.MahjongBTEOperaNtf, loseScores map[int32]int64) {
		for loserSeat, score := range loseScores {
			rival := s.mahjongPlayers[loserSeat]
			rival.score -= score
			p.score += score
			p.gangScore[lastPeerIndex] = append(p.gangScore[lastPeerIndex], loserSeat)
		}

		// 先组装杠成功得通知消息
		gangResultNtf := &outer.MahjongBTEGangResultNtf{
			OpShortId:        p.ShortId,
			QiangGangShortId: 0,
			Card:             card.Int32(),
		}

		// 如果需要实时结算, 就把算分数据带上
		if s.gameParams().GangImmediatelyScore {
			// 所有人最新得分数
			for _, player := range s.mahjongPlayers {
				gangResultNtf.CurrentScores = append(gangResultNtf.CurrentScores, player.score)
			}
			gangResultNtf.LoseScores = loseScores // 每个赔分得人
		}

		// 不等于nil, 说明没有抢杠, 将通知带入操作中一并发出
		if opNtf != nil {
			opNtf.GangResult = gangResultNtf
		} else {
			// 如果有抢杠，单独广播此通知
			s.room.Broadcast(gangResultNtf)
		}
	}

	switch peer.typ {
	case drawCardType: // 摸牌
		if _, ok := p.pong[card.Int32()]; ok {
			ntf.GangType = 1 // 面下杠（刮风）
			gangType = GangType1

			loseScores = loseScoreAnalyze(s.allSeatsWithoutHu(seatIndex)...) // 其余三家输分
			gangFunc = func(opNtf *outer.MahjongBTEOperaNtf) {
				s.Log().Infow("gang ok by draw card with pong")
				delete(p.pong, card.Int32())
				p.lightGang[card.Int32()] = p.ShortId
				execScoreFunc(opNtf, loseScores)
			}
			qiangGang = hasQiangGang()
			ntf.Card = card.Int32() // 杠的牌
		} else {
			ntf.GangType = 2
			gangType = GangType4

			loseScores = loseScoreAnalyze(s.allSeatsWithoutHu(seatIndex)...) // 其余三家输分
			// 暗杠（下雨）
			gangFunc = func(opNtf *outer.MahjongBTEOperaNtf) {
				s.Log().Infow("gang ok by draw card")
				p.handCards = p.handCards.Remove(card, card, card, card)
				p.darkGang[card.Int32()] = p.ShortId
				execScoreFunc(opNtf, loseScores)
			}
		}

	case playCardType:
		gangType = GangType3
		ntf.GangType = 1 // 直杠（刮风）

		loseScores = loseScoreAnalyze(peer.seat) // 打牌的那个人，是输分者
		gangFunc = func(opNtf *outer.MahjongBTEOperaNtf) {
			s.Log().Infow("gang ok by play card")
			p.handCards, _, _ = p.handCards.Gang(card)
			p.lightGang[card.Int32()] = s.mahjongPlayers[peer.seat].ShortId
			execScoreFunc(opNtf, loseScores)
		}

		qiangGang = hasQiangGang()
		ntf.Card = card.Int32() // 杠的牌
	}

	// 没有人能抢杠，直接执行杠
	if !qiangGang {
		gangFunc(ntf)
		gangFunc = nil
	}

	s.appendPeerCard(gangType, card, seatIndex, gangFunc)
	return true, outer.ERROR_OK
}

// 获得排除了某些座位后，剩余的没胡座位
func (s *StatePlaying) allSeatsWithoutHu(ignoreSeat ...int) (result []int) {
	seatMap := map[int]struct{}{}
	for _, seat := range ignoreSeat {
		seatMap[seat] = struct{}{}
	}

	for seatIndex := 0; seatIndex < maxNum; seatIndex++ {
		if _, ignore := seatMap[seatIndex]; !ignore {
			result = append(result, seatIndex)
		}
	}

	return result
}
