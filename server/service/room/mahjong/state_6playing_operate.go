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

	ntf := &outer.MahjongBTEOperaNtf{
		OpShortId: player.ShortId,
		OpType:    op,
	}

	nextDrawSeatIndex := s.nextSeatIndex(s.peerRecords[len(s.peerRecords)-1].seat)

	switch op {
	case outer.ActionType_ActionPass:
		s.Log().Infow("pass", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId, "hand", player.handCards)

		// 检查抢杠胡的情况，所有人都过了，需要执行杠的行为
		if s.husWasAllPass() {
			lastPeer := s.peerRecords[len(s.peerRecords)-1]
			if lastPeer.typ >= GangType1 && lastPeer.afterQiangPass != nil {
				lastPeer.afterQiangPass(nil)
				lastPeer.afterQiangPass = nil
			}
		}
		ok = true
		delete(s.Hus, seatIndex)
		delete(s.actionMap, seatIndex)

		if s.husWasAllDo() {
			s.huSettlement(nil) // 传nil，表示ntf单独推送
		}

	case outer.ActionType_ActionPong:
		ok, err = s.operatePong(player, seatIndex)
		peer := s.peerRecords[len(s.peerRecords)-1]
		ntf.Card = peer.card.Int32() // 碰的牌

		s.Log().Color(logger.Green).Infow("pong!", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId,
			"peer", s.peerRecordsLog(), "hand", player.handCards, "pong cards", player.pong)

	case outer.ActionType_ActionGang:
		ok, err = s.operateGang(player, seatIndex, card, ntf)
		nextDrawSeatIndex = seatIndex // 杠的人自己摸一张
		delete(s.actionMap, seatIndex)

		s.Log().Color(logger.Green).Infow("gang!", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId,
			"peer", s.peerRecordsLog(), "action map", s.actionMap, "hand", player.handCards, "lightGang cards", player.lightGang, "darkGang cards", player.darkGang)

	case outer.ActionType_ActionHu:
		ok, err = s.operateHu(player, seatIndex, ntf)
		nextDrawSeatIndex = s.nextSeatIndex(seatIndex) // 胡牌的下家摸牌
		delete(s.actionMap, seatIndex)

		s.Log().Color(logger.Red).Infow("hu!", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId, "peer", s.peerRecordsLog(), "hand", player.handCards,
			"pong", player.pong, "lightGang cards", player.lightGang, "darkGang cards", player.darkGang, "hu", player.hu, "hu extra", player.huExtra)

	default:
		s.Log().Errorw("unknown action op",
			"room", s.room.RoomId, "player", player.ShortId, "op", op)
		return false, outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	if !ok {
		return
	}

	// 操作成功，删除行为
	s.removeCurrentAction(seatIndex)

	// 除了过以外的操作都需要广播通知
	if op != outer.ActionType_ActionPass {
		ntf.HandCardsNum = int32(player.handCards.Len())
		s.room.Broadcast(ntf)
	}

	// 没有可行动的人，就摸牌
	if len(s.actionMap) == 0 {
		s.Hus = make(map[int]bool) // 每次摸牌清一次胡牌状态数据
		s.drawCard(nextDrawSeatIndex)
	}
	s.nextAction() // 碰、杠、胡、过 后的下个行为

	return true, outer.ERROR_OK
}
