package mahjong

import (
	"server/proto/outermsg/outer"
)

// 碰杠胡过
func (s *StatePlaying) operate(player *mahjongPlayer, seatIndex int, op outer.ActionType, hu HuType, card Card) (ok bool, err outer.ERROR) {
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

	nextDrawShortIndex := s.nextSeatIndex(s.peerRecords[len(s.peerRecords)-1].seat) //提前计算下家摸牌的座位

	peer := s.peerRecords[len(s.peerRecords)-1]
	switch op {
	case outer.ActionType_ActionPass:
		// 检查抢杠胡的情况，所有人都过了，需要执行杠的行为
		if len(s.actionMap) == 0 && len(s.peerRecords) > 0 {
			lastPeer := s.peerRecords[len(s.peerRecords)-1]
			if lastPeer.typ >= GangType1 && lastPeer.afterQiangPass != nil {
				lastPeer.afterQiangPass(nil)
				lastPeer.afterQiangPass = nil
			}
		}
		ok = true
		s.Log().Infow("pass", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId, "hand", player.handCards)

	case outer.ActionType_ActionPong:
		ok, err = s.operatePong(player, seatIndex)
		ntf.Card = peer.card.Int32() // 碰的牌

		s.Log().Infow("pong!", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId,
			"peer", &peer, "hand", player.handCards, "pong cards", player.pong)

	case outer.ActionType_ActionGang:
		ok, err = s.operateGang(player, seatIndex, card, ntf)
		ntf.Card = card.Int32()        // 杠的牌
		nextDrawShortIndex = seatIndex // 杠的人自己摸一张

		s.Log().Infow("gang!", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId,
			"peer", s.peerRecords[len(s.peerRecords)-1], "action map", s.actionMap, "hand", player.handCards, "lightGang cards", player.lightGang, "darkGang cards", player.darkGang)

	case outer.ActionType_ActionHu:
		ok, err = s.operateHu(player, seatIndex, ntf)
		nextDrawShortIndex = s.nextSeatIndex(seatIndex) // 胡牌的下家摸牌

		s.Log().Infow("hu!", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId, "peer", &peer, "hand", player.handCards,
			"pong", player.pong, "lightGang cards", player.lightGang, "darkGang cards", player.darkGang, "hu", player.hu, "hu extra", player.huExtra)

	default:
		s.Log().Errorw("unknown action op",
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
	s.nextAction() // 碰、杠、胡、过 后的下个行为

	return true, outer.ERROR_OK
}
