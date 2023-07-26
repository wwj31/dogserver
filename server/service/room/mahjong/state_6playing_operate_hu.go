package mahjong

import (
	"server/proto/outermsg/outer"
)

// 胡牌操作
func (s *StatePlaying) operateHu(p *mahjongPlayer, seatIndex int, ntf *outer.MahjongBTEOperaNtf) (bool, outer.ERROR) {
	var paySeat []int // 需要赔钱的座位
	// 获得最后一次操作的牌
	lastPeerIndex := len(s.peerRecords) - 1
	peer := s.peerRecords[lastPeerIndex]
	var hu HuType
	switch peer.typ {
	case drawCardType: // 自摸
		hu = p.handCards.IsHu(p.lightGang, p.darkGang, p.pong, peer.card, s.gameParams())
		// 其余没胡的都要赔钱
		for seat, other := range s.mahjongPlayers {
			if other.hu != HuInvalid || other.ShortId == p.ShortId {
				continue
			}
			paySeat = append(paySeat, seat)
		}

	case playCardType, GangType1, GangType3: // 点炮,抢杠
		hu = p.handCards.Insert(peer.card).IsHu(p.lightGang, p.darkGang, p.pong, peer.card, s.gameParams())
		paySeat = append(paySeat, peer.seat) // 点炮的人陪钱

	}

	if hu == HuInvalid {
		s.Log().Errorw("operate hu invalid", "room", s.room.RoomId, "seat", seatIndex, "player", p.ShortId, "hand", p.handCards)
		return false, outer.ERROR_MAHJONG_HU_INVALID
	}

	// 一炮多响检查，如果还有人胡了相同的peer，就算一炮多响
	if s.checkMutilHu(lastPeerIndex) {
		s.mutilHuByIndex = peer.seat
	}

	s.huSeat = append(s.huSeat, int32(seatIndex))
	p.hu = hu
	p.huPeerIndex = len(s.peerRecords) - 1
	ntf.HuType = hu.PB()
	if peer.typ == GangType1 || peer.typ == GangType3 {
		peer.afterQiangPass = nil // 抢杠成功，杠的人，杠失败

		s.room.Broadcast(&outer.MahjongBTEGangResultNtf{
			OpShortId:        s.mahjongPlayers[peer.seat].ShortId,
			QiangGangShortId: p.ShortId,
			Card:             peer.card.Int32(),
			LoseScores:       nil,
			CurrentScores:    nil,
		})
		s.Log().Infow("qiang gang success", "room", s.room.RoomId, "seat", seatIndex, "gang seat", peer.seat)
	}

	// 计算额外加番
	p.huExtra = s.huExtra(seatIndex)
	p.huGen = int32(s.huGen(seatIndex))

	ntf.HuExtraType = p.huExtra.PB()

	// 胡成功后，删除Gang和Pong(可以一炮多响,但是有人胡了就不能再碰、杠)
	for seat, act := range s.actionMap {
		act.remove(outer.ActionType_ActionGang)
		act.remove(outer.ActionType_ActionPong)
		act.gang = []int32{}

		// 只剩下[过]操作，删掉
		if len(act.acts) == 1 && act.isValidAction(outer.ActionType_ActionPass) {
			act.remove(outer.ActionType_ActionPass)
		}

		if len(act.acts) == 0 {
			delete(s.actionMap, seat)
		}
	}

	// TODO 算番算分
	if len(s.peerRecords) > 3 {
		last1 := s.peerRecords[len(s.peerRecords)-1]
		last2 := s.peerRecords[len(s.peerRecords)-2]
		last3 := s.peerRecords[len(s.peerRecords)-3]

		// 判断是否是杠上炮
		if (last3.typ == GangType1 || last3.typ == GangType3 || last3.typ == GangType4) &&
			last1.seat == last2.seat && last1.seat == last3.seat {
			// TODO 杠上炮
		}
	}

	return true, outer.ERROR_OK
}

// 分析是否有额外加番
func (s *StatePlaying) huExtra(seatIndex int) ExtFanType {
	var extraFans []ExtFanType

	// 根据番数大到小，优先计算大番型
	if len(s.peerRecords) == 0 && s.gameParams().TianDiHu {
		return TianHu
	}

	if len(s.peerRecords) == 1 && s.gameParams().TianDiHu {
		return Dihu
	}

	lastPeerCard := s.peerRecords[len(s.peerRecords)-1]

	// 没牌了，执行[扫底胡]和[海底炮]检测
	if s.cards.Len() == 0 {
		switch lastPeerCard.typ {
		case drawCardType:
			extraFans = append(extraFans, ShaoDiHu) // 最后一张牌，摸起来胡了，扫底胡
		case playCardType:
			extraFans = append(extraFans, HaiDiPao) // 最后一张牌，摸起来后出牌点炮了，海底炮
		}
	}

	p := s.mahjongPlayers[seatIndex]
	if p.handCards.Len() == 2 {
		extraFans = append(extraFans, JinGouGou) // 只剩2张牌做将，金钩胡
	}

	// 抢杠
	if lastPeerCard.typ == GangType1 {
		extraFans = append(extraFans, QiangGangHu) // 抢杠胡
	}

	// 如果上次是杠，这次一定是摸牌，判断是否杠上花
	if len(s.peerRecords) >= 2 {
		beforeLastPeerCard := s.peerRecords[len(s.peerRecords)-2]
		if beforeLastPeerCard.typ >= GangType1 {
			extraFans = append(extraFans, GangShangHua) // 刚上花
		}
	}

	// 如果上上次是杠，这次一定是出牌，判断是否杠上炮
	if len(s.peerRecords) >= 3 {
		beforeBeforeLastPeerCard := s.peerRecords[len(s.peerRecords)-3]
		if beforeBeforeLastPeerCard.typ >= GangType1 {
			extraFans = append(extraFans, GangShangPao) // 杠上炮
		}
	}

	return 0
}

// 算根
func (s *StatePlaying) huGen(seatIndex int) (count int) {
	p := s.mahjongPlayers[seatIndex]
	count = len(p.lightGang) + len(p.darkGang)
	for pongCard := range p.pong {
		p.handCards.Range(func(card Card) bool {
			if pongCard == card.Int32() {
				count++
				return true
			}
			return false
		})
	}

	cardsStat := p.handCards.ConvertStruct()
	for _, num := range cardsStat {
		if num == 4 {
			count++
		}
	}
	return
}
