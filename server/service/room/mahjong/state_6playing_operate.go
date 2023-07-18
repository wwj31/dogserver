package mahjong

import (
	"server/common/log"
	"server/proto/outermsg/outer"
)

// 碰杠胡过
func (s *StatePlaying) operate(player *mahjongPlayer, seatIndex int, op outer.ActionType, card Card) (ok bool, err outer.ERROR) {
	if op == outer.ActionType_ActionPlayCard {
		// 此函数不受理打牌
		return false, outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	if !s.currentAction.isValidAction(op) {
		return false, outer.ERROR_MAHJONG_ACTION_PLAYER_NOT_OPERA
	}

	ntf := &outer.MahjongBTEOperaNtf{
		OpShortId: player.ShortId,
		OpType:    op,
	}

	nextDrawShortIndex := s.nextSeatIndex(s.peerCards[len(s.peerCards)-1].seat) //提前计算下家摸牌的座位

	peer := s.peerCards[len(s.peerCards)-1]
	switch op {
	case outer.ActionType_ActionPass:
		// 检查抢杠胡的情况，所有人都过了，需要执行杠的行为
		if len(s.actionMap) == 0 && len(s.peerCards) > 0 {
			lastPeer := s.peerCards[len(s.peerCards)-1]
			if lastPeer.typ >= GangType1 && lastPeer.afterQiangPass != nil {
				lastPeer.afterQiangPass()
			}
		}
		ok = true
		log.Infow("pass", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId, "hand", player.handCards)

	case outer.ActionType_ActionPong:
		ok, err = s.operatePong(player, seatIndex)
		ntf.Card = peer.card.Int32() // 碰的牌

		log.Infow("pong!", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId,
			"peer", &peer, "hand", player.handCards, "pong cards", player.pong)

	case outer.ActionType_ActionGang:
		ok, err = s.operateGang(player, seatIndex, card, ntf)
		ntf.Card = card.Int32()        // 杠的牌
		nextDrawShortIndex = seatIndex // 杠的人自己摸一张

		log.Infow("gang!", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId,
			"peer", &peer, "hand", player.handCards, "lightGang cards", player.lightGang, "darkGang cards", player.darkGang)

	case outer.ActionType_ActionHu:
		ok, err = s.operateHu(player, seatIndex, ntf)
		nextDrawShortIndex = s.nextSeatIndex(seatIndex) // 胡牌的下家摸牌

		log.Infow("hu!", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId, "peer", &peer, "hand", player.handCards,
			"pong", player.pong, "lightGang cards", player.lightGang, "darkGang cards", player.darkGang, "hu", player.hu, "hu extra", player.huExtra)

	default:
		log.Errorw("unknown action op",
			"room", s.room.RoomId, "player", player.ShortId, "op", op)
		return false, outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	if !ok {
		return
	}

	// 除了过以外的操作都需要广播通知
	if op != outer.ActionType_ActionPass {
		s.room.Broadcast(ntf)
	}

	// 没有可行动的人，就摸牌
	if len(s.actionMap) == 0 {
		s.drawCard(nextDrawShortIndex)
	}
	s.nextAction()

	return true, outer.ERROR_OK
}

func (s *StatePlaying) gameOver() bool {
	huCount := 0
	for _, p := range s.mahjongPlayers {
		if p.hu != HuInvalid {
			huCount++
			if huCount >= 3 {
				return true
			}
		}
	}

	if s.cards.Len() == 0 {
		return true
	}
	return false
}

// 碰牌操作
func (s *StatePlaying) operatePong(p *mahjongPlayer, seatIndex int) (bool, outer.ERROR) {
	if len(s.peerCards) == 0 {
		log.Errorw("operate pong failed peerCards len = 0",
			"room", s.room.RoomId, "player", p.ShortId)
		return false, outer.ERROR_MSG_REQ_PARAM_INVALID
	}
	// 获得打出的那张牌
	peer := s.peerCards[len(s.peerCards)-1]
	if peer.typ != playCardType {
		log.Errorw("operate pong failed peer is drawCard",
			"room", s.room.RoomId, "player", p.ShortId, "peerCards", s.peerCards)
		return false, outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	// 检查,桌面牌最后一张是否和碰的牌一致
	tail := len(s.cardsInDesktop) - 1
	desktopCard := s.cardsInDesktop[tail]
	if desktopCard != peer.card {
		log.Errorw("operate pong logic error",
			"room", s.room.RoomId, "player", p.ShortId, "peerCards", s.peerCards, "desktopCard", desktopCard)
		return false, outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	// 检查,不可能自己碰自己的牌
	if peer.seat == seatIndex {
		log.Errorw("unexpected logic ",
			s.room.RoomId, "player", p.ShortId, "peer", peer, "seat", seatIndex)
		return false, outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	var err error
	p.handCards, _, err = p.handCards.Pong(peer.card)
	if err != nil {
		log.Errorw("unexpected logic pong failed",
			"room", s.room.RoomId, "seat", seatIndex, "player", p.ShortId, "peer", peer, "err", err)
		return false, outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	s.cardsInDesktop = s.cardsInDesktop[:tail]                                                            // 删除桌面牌
	p.pong[peer.card.Int32()] = s.mahjongPlayers[peer.seat].ShortId                                       // 加入自己的碰牌组
	s.actionMap[seatIndex] = &action{currentActions: []outer.ActionType{outer.ActionType_ActionPlayCard}} // 碰后新增出牌行为
	return true, outer.ERROR_OK
}

// 杠牌操作
func (s *StatePlaying) operateGang(p *mahjongPlayer, seatIndex int, card Card, ntf *outer.MahjongBTEOperaNtf) (ok bool, err outer.ERROR) {
	if len(s.peerCards) == 0 {
		log.Errorw("operate gang failed peerCards len = 0",
			"room", s.room.RoomId, "player", p.ShortId)
		return false, outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	hasQiangGang := func() bool {
		b := false
		for seat, other := range s.mahjongPlayers {
			if seatIndex == seat {
				continue
			}

			if hu := other.handCards.Insert(card).IsHu(other.lightGang, other.darkGang, other.pong, card); hu != HuInvalid {
				newAction := action{}
				newAction.currentActions = append(newAction.currentActions, outer.ActionType_ActionHu, outer.ActionType_ActionPass)
				newAction.currentHus = append(newAction.currentHus, hu.PB())
				s.actionMap[seat] = &newAction // 抢杠胡操作
				b = true
			}
		}
		return b
	}

	qiangGang := false
	// 获得最后一次操作的牌
	peer := s.peerCards[len(s.peerCards)-1]
	var (
		gangFunc func()
		gangType checkCardType
	)
	switch peer.typ {
	case drawCardType: // 摸牌
		if _, ok := p.pong[card.Int32()]; ok {
			ntf.GangType = 1 // 面下杠（刮风）
			gangType = GangType1

			gangFunc = func() {
				delete(p.pong, card.Int32())
				p.darkGang[card.Int32()] = p.ShortId
			}

			qiangGang = hasQiangGang()
		} else {
			ntf.GangType = 2
			gangType = GangType4

			// 暗杠（下雨）
			gangFunc = func() {
				for _, gang := range p.handCards.HasGang() {
					if gang == card {
						p.handCards = p.handCards.Remove(card, card, card, card)
						break
					}
				}
				p.darkGang[card.Int32()] = p.ShortId
			}
		}

	case playCardType:
		ntf.GangType = 1 // 直杠（刮风）
		gangType = GangType3
		gangFunc = func() {
			p.handCards, _, _ = p.handCards.Gang(card)
			p.lightGang[card.Int32()] = s.mahjongPlayers[peer.seat].ShortId
		}

		qiangGang = hasQiangGang()
	}

	// 没有人能抢杠，直接执行杠
	if !qiangGang {
		gangFunc()
		gangFunc = nil
	}

	s.appendPeerCard(gangType, card, seatIndex, gangFunc)
	return true, outer.ERROR_OK
}

// 胡牌操作
func (s *StatePlaying) operateHu(p *mahjongPlayer, seatIndex int, ntf *outer.MahjongBTEOperaNtf) (bool, outer.ERROR) {
	var paySeat []int // 需要赔前的座位
	// 获得最后一次操作的牌
	lastPeerIndex := len(s.peerCards) - 1
	peer := s.peerCards[lastPeerIndex]
	var hu HuType
	switch peer.typ {
	case drawCardType: // 自摸
		hu = p.handCards.IsHu(p.lightGang, p.darkGang, p.pong, peer.card)
		// 其余没胡的都要赔钱
		for seat, other := range s.mahjongPlayers {
			if other.hu != HuInvalid || other.ShortId == p.ShortId {
				continue
			}
			paySeat = append(paySeat, seat)
		}

	case playCardType, GangType1: // 点炮,抢杠
		hu = p.handCards.Insert(peer.card).IsHu(p.lightGang, p.darkGang, p.pong, peer.card)
		paySeat = append(paySeat, peer.seat) // 点炮的人陪钱
	}

	if hu == HuInvalid {
		log.Errorw("operate hu invalid", "room", s.room.RoomId, "seat", seatIndex, "player", p.ShortId, "hand", p.handCards)
		return false, outer.ERROR_MAHJONG_HU_INVALID
	}

	// 一炮多响,如果还有人胡了这个peer，那么就算一炮多响
	if s.checkMutilHu(lastPeerIndex) {
		s.mutilHuByIndex = peer.seat
	}

	if s.firstHuIndex == -1 {
		s.firstHuIndex = seatIndex
	}
	p.hu = hu
	p.huPeerIndex = len(s.peerCards) - 1
	ntf.HuType = hu.PB()
	if peer.typ == GangType1 || peer.typ == GangType3 {
		peer.afterQiangPass = nil // 抢杠成功，杠的人，杠失败
		ntf.QiangGangHuCard = peer.card.Int32()
	}

	// 计算额外加番
	p.huExtra = s.huExtra(seatIndex)
	p.huGen = int32(s.huGen(seatIndex))

	ntf.HuExtraType = p.huExtra.PB()

	// 胡成功后，删除Gang和Pong(可以一炮多响,但是有人胡了就不能再碰、杠)
	for seat, act := range s.actionMap {
		act.remove(outer.ActionType_ActionGang)
		act.remove(outer.ActionType_ActionPong)
		act.currentGang = []int32{}

		// 只剩下[过]操作，删掉
		if len(act.currentActions) == 1 && act.isValidAction(outer.ActionType_ActionPass) {
			act.remove(outer.ActionType_ActionPass)
		}

		if len(act.currentActions) == 0 {
			delete(s.actionMap, seat)
		}
	}

	// TODO 算番算分

	return true, outer.ERROR_OK
}

// 分析是否有额外加番
func (s *StatePlaying) huExtra(seatIndex int) ExtFanType {
	var extraFans []ExtFanType

	// 根据番数大到小，优先计算大番型
	if len(s.peerCards) == 1 {
		return TianHu
	}

	if len(s.peerCards) == 2 {
		return Dihu
	}

	lastPeerCard := s.peerCards[len(s.peerCards)-1]

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

	// 如果上上次是杠，那么上次一定是摸牌，判断是否杠上花
	if len(s.peerCards) >= 2 {
		beforeLastPeerCard := s.peerCards[len(s.peerCards)-2]
		if beforeLastPeerCard.typ >= GangType1 {
			extraFans = append(extraFans, GangShangHua) // 刚上花
		}
	}

	// 如果上上上次是杠，那么上次一定是出牌，判断是否杠上炮
	if len(s.peerCards) >= 3 {
		beforeBeforeLastPeerCard := s.peerCards[len(s.peerCards)-3]
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
