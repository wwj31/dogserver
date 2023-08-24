package mahjong

import (
	"github.com/wwj31/dogactor/logger"

	"server/proto/outermsg/outer"
)

// 碰杠胡过
func (s *StatePlaying) operate(player *mahjongPlayer, seatIndex int, op outer.ActionType, hu HuType, card Card) (ok bool, err outer.ERROR) {
	if op == outer.ActionType_ActionPlayCard {
		// 此函数不受理打牌
		return false, outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	currentAction := s.getCurrentAction(seatIndex)
	if currentAction == nil || !currentAction.isValidAction(op) {
		return false, outer.ERROR_MAHJONG_ACTION_PLAYER_NOT_OPERA
	}

	// 打牌的时候，选过 把其他操作都删了
	if currentAction.isValidAction(outer.ActionType_ActionPlayCard) && op == outer.ActionType_ActionPass {
		currentAction.acts = []outer.ActionType{outer.ActionType_ActionPlayCard}
		currentAction.hus = nil
		currentAction.gang = nil
		s.Log().Infof("pass play card")
		return true, outer.ERROR_OK
	}

	// 如果有胡操作，碰杠过都把胡状态删除
	if op != outer.ActionType_ActionHu {
		delete(s.canHus, seatIndex)
	}

	// 碰杠，需要单独判断是否在一炮多响的情况下
	if op == outer.ActionType_ActionPong || op == outer.ActionType_ActionGang {
		// 一炮多响,如果还有人能胡，但是没胡，需要保留操作
		if s.needWaiting4Hu() {
			s.Log().Infow(" YiPaoDuoXiang PongGang before other hu",
				"seat", seatIndex, "short", player.ShortId, "op", op, "hu", hu, "card", card)
			s.HusPongGang = func() { s.operate(player, seatIndex, op, hu, card) }
			return true, outer.ERROR_OK
		}

	}

	ntf := &outer.MahjongBTEOperaNtf{
		OpShortId: player.ShortId,
		OpType:    op,
	}

	nextDrawSeatIndex := s.nextSeatIndex(s.peerRecords[len(s.peerRecords)-1].seat)

	delete(s.actionMap, seatIndex)

	switch op {
	case outer.ActionType_ActionPass:
		s.Log().Infow("pass", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId, "hand", player.handCards)
		ok = true

		// 一炮多响，所有人都过了
		if s.husWasAllPass() {
			// 检查抢杠胡的情况，需要执行杠的行为
			lastPeer := s.peerRecords[len(s.peerRecords)-1]
			if lastPeer.typ >= GangType1 && lastPeer.afterQiangPass != nil {
				lastPeer.afterQiangPass(nil)
				lastPeer.afterQiangPass = nil
			}

			// 如果有人在一炮多响期间，点了胡碰操作，需要还原操作
			if s.HusPongGang != nil {
				s.HusPongGang()
				s.HusPongGang = nil
			}
		}

		if s.husWasAllDo() {
			s.HusPongGang = nil
			s.huSettlement(nil) // 传nil，表示ntf单独推送
		}

	case outer.ActionType_ActionPong:
		// 一炮多响，已经有人胡了，直接判断结算
		if s.husWasAllDo() {
			s.Log().Infow("pong trigger settlement", "hus", s.canHus)
			s.HusPongGang = nil
			s.huSettlement(nil) // 传nil，表示ntf单独推送
			ok = true
			break
		}

		ok, err = s.operatePong(player, seatIndex)
		peer := s.peerRecords[len(s.peerRecords)-1]
		ntf.Card = peer.card.Int32() // 碰的牌

		s.Log().Color(logger.Green).Infow("pong!", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId,
			"peer", s.peerRecordsLog(), "hand", player.handCards, "pong cards", player.pong)

	case outer.ActionType_ActionGang:
		// 一炮多响，已经有人胡了，直接判断结算
		if s.husWasAllDo() {
			s.Log().Infow("gang trigger settlement", "hus", s.canHus)
			s.HusPongGang = nil
			s.huSettlement(nil) // 传nil，表示ntf单独推送
			ok = true
			break
		}

		ok, err = s.operateGang(player, seatIndex, card, ntf)
		nextDrawSeatIndex = seatIndex // 杠的人自己摸一张

		s.Log().Color(logger.Green).Infow("gang!", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId,
			"peer", s.peerRecordsLog(), "action map", s.actionMap, "hand", player.handCards, "lightGang cards", player.lightGang, "darkGang cards", player.darkGang)

	case outer.ActionType_ActionHu:
		// 一炮多响，已经有人胡了，直接判断能否结算
		if s.husWasAllDo() {
			s.HusPongGang = nil
			s.huSettlement(nil) // 传nil，表示ntf单独推送
			ok = true
			break
		}

		ok, err = s.operateHu(player, seatIndex, ntf)
		nextDrawSeatIndex = s.nextSeatIndex(seatIndex) // 胡牌的下家摸牌

		s.Log().Color(logger.Red).Infow("hu!", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId, "peer", s.peerRecordsLog(), "hand", player.handCards,
			"pong", player.pong, "lightGang cards", player.lightGang, "darkGang cards", player.darkGang, "hu", player.hu, "hu extra", player.huExtra)

	default:
		s.Log().Errorw("unknown action op",
			"room", s.room.RoomId, "player", player.ShortId, "op", op)
		return false, outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	s.removeCurrentAction(seatIndex) // 删除行为

	if !ok {
		return
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
	return true, outer.ERROR_OK
}

func (s *StatePlaying) afterOperate(nextDrawCardSeat int) {
	// 没有可行动的人，就摸牌
	if len(s.actionMap) == 0 {
		s.drawCard(nextDrawCardSeat)
	}
	s.nextAction() // 碰、杠、胡、过 后的下个行为
}
