package mahjong

import (
	"github.com/wwj31/dogactor/logger"

	"server/proto/outermsg/outer"
)

// 碰杠胡过
func (s *StatePlaying) operate(player *mahjongPlayer, seatIndex int, op outer.ActionType, hu HuType, card Card) (err outer.ERROR) {
	if op == outer.ActionType_ActionPlayCard {
		// 此函数不受理打牌
		return outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	currentAction := s.getCurrentAction(seatIndex)
	if currentAction == nil || !currentAction.isValidAction(op) {
		return outer.ERROR_MAHJONG_ACTION_PLAYER_NOT_OPERA
	}

	// 打牌的时候，选过 把其他操作都删了
	if currentAction.isValidAction(outer.ActionType_ActionPlayCard) && op == outer.ActionType_ActionPass {
		currentAction.acts = []outer.ActionType{outer.ActionType_ActionPlayCard}
		currentAction.hus = nil
		currentAction.gang = nil
		s.Log().Infow("operate pass with play card", "shortId", player.ShortId, "seat", seatIndex, "card", card)
		return outer.ERROR_OK
	}

	delete(s.actionMap, seatIndex)
	if op != outer.ActionType_ActionHu {
		// 如果有胡操作，碰杠过都把胡状态删除
		delete(s.canHus, seatIndex)
	}

	// 针对一炮多响情况，提前对碰杠操作做一次检查
	switch op {
	case outer.ActionType_ActionPong, outer.ActionType_ActionGang: // 碰杠，需要单独判断是否在一炮多响的情况下
		if len(s.canHus) > 0 {
			// 一炮多响的情况下,
			// 如果还没有人胡,需要保留碰杠操作,能胡的人选过再执行碰杠,
			// 否则，碰杠操作统统视为过.

			if !s.husWasDo() {
				s.Log().Infow("delay this pong/gang operator",
					"seat", seatIndex, "short", player.ShortId, "op", op, "card", card)
				s.HusPongGang = func() { s.operate(player, seatIndex, op, hu, card) }

				return outer.ERROR_OK
			} else {
				player.passHandHuFan = 0 // 这里需要清除过手胡
				s.Log().Infow("replace pong/gang with pass",
					"seat", seatIndex, "short", player.ShortId, "op", op, "card", card)
				op = outer.ActionType_ActionPass
			}
		}
	}

	ntf := &outer.MahjongBTEOperaNtf{
		OpShortId: player.ShortId,
		OpType:    op,
	}

	nextDrawSeatIndex := s.nextSeatIndex(s.peerRecords[len(s.peerRecords)-1].seat)

	switch op {
	case outer.ActionType_ActionPass:
		s.Log().Infow("pass", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId, "hand", player.handCards)
		s.removeCurrentAction(seatIndex) // 删除行为

		if s.husWasAllPass() { // 一炮多响，所有人都过了
			// 检查抢杠胡的情况，需要执行杠的行为
			lastPeer := s.peerRecords[len(s.peerRecords)-1]
			if lastPeer.typ == GangType1 && lastPeer.afterQiangPass != nil {
				lastPeer.afterQiangPass(nil)
				lastPeer.afterQiangPass = nil
			}

			// 如果有人在一炮多响期间，点了碰杠，需要延续执行操作
			if s.HusPongGang != nil {
				s.HusPongGang()
				s.HusPongGang = nil
				return outer.ERROR_OK
			}

		} else if !s.needWaiting4Hu() { // 一炮多响，不用等了
			s.HusPongGang = nil
			s.huSettlement(nil) // 传nil，表示ntf单独推送
		}

	case outer.ActionType_ActionPong:
		if err = s.operatePong(player, seatIndex); err != outer.ERROR_OK {
			return err
		}
		s.removeCurrentAction(seatIndex) // 删除行为

		peer := s.peerRecords[len(s.peerRecords)-1]
		ntf.Card = peer.card.Int32() // 碰的牌
		s.Log().Color(logger.Green).Infow("pong!", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId,
			"peer", s.peerRecordsLog(), "hand", player.handCards, "pong cards", player.pong)

	case outer.ActionType_ActionGang:
		if err = s.operateGang(player, seatIndex, card, ntf); err != outer.ERROR_OK {
			return err
		}
		s.removeCurrentAction(seatIndex) // 删除行为

		nextDrawSeatIndex = seatIndex // 杠的人自己摸一张
		s.Log().Color(logger.Green).Infow("gang!", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId,
			"peer", s.peerRecordsLog(), "action map", s.actionMap, "hand", player.handCards, "lightGang cards", player.lightGang, "darkGang cards", player.darkGang)

	case outer.ActionType_ActionHu:
		if err = s.operateHu(player, seatIndex, ntf); err != outer.ERROR_OK {
			return err
		}
		s.removeCurrentAction(seatIndex) // 删除行为

		nextDrawSeatIndex = s.nextSeatIndex(seatIndex) // 胡牌的下家摸牌
		s.Log().Color(logger.Red).Infow("hu!", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId, "peer", s.peerRecordsLog(), "hand", player.handCards,
			"pong", player.pong, "lightGang cards", player.lightGang, "darkGang cards", player.darkGang, "hu", player.hu, "hu extra", player.huExtra)

	default:
		s.Log().Errorw("unknown action op",
			"room", s.room.RoomId, "player", player.ShortId, "op", op)
		return outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	// 还有人可以胡，但没胡，先等待其他人操作
	if s.needWaiting4Hu() {
		s.Log().Infow("YiPaoDuoXiang waiter other hu", "hus", s.canHus)
		return
	}

	// 1.过操作，始终不广播.
	// 2.碰杠胡操作，不在一炮多响的情况下要广播，
	//   在一炮多响中，如果不用等待其他人操作，要广播.
	if op != outer.ActionType_ActionPass && !s.needWaiting4Hu() {
		ntf.HandCardsNum = int32(player.handCards.Len())
		s.room.Broadcast(ntf)
	}

	// 如果触发了一炮多响，并且至少2个人胡了，就点炮的人摸牌
	if len(s.canHus) >= 2 {
		nextDrawSeatIndex = s.multiHuByIndex
	} else if len(s.canHus) == 1 { // 如果一炮多响的时候，只有一个人胡了，那么下次出来就是他的下家
		for seat := range s.canHus {
			nextDrawSeatIndex = s.nextSeatIndex(seat) // 胡牌的下家摸牌
		}
	}

	s.afterOperate(nextDrawSeatIndex)
	return outer.ERROR_OK
}

func (s *StatePlaying) afterOperate(nextDrawCardSeat int) {
	// 没有可行动的人，就摸牌
	if len(s.actionMap) == 0 {
		s.drawCard(nextDrawCardSeat)
	}
	s.nextAction() // 碰、杠、胡、过 后的下个行为
}
