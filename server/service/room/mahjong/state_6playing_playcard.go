package mahjong

import (
	"server/common/log"
	"server/proto/outermsg/outer"
)

// 打一张牌
func (s *StatePlaying) playCard(cardIndex, seatIndex int) (bool, outer.ERROR) {
	var (
		act   *action
		exist bool
	)

	// 检查是否是行动者,以及行为是否有效
	if act, exist = s.actionMap[seatIndex]; !exist {
		return false, outer.ERROR_MAHJONG_ACTION_PLAYER_NOT_MATCH
	} else if !act.isValidAction(outer.ActionType_ActionPlayCard) {
		return false, outer.ERROR_MAHJONG_ACTION_PLAYER_NOT_OPERA
	}

	// 把打的牌从手牌移除
	player := s.mahjongPlayers[seatIndex]
	outCard := player.handCards[cardIndex]
	player.handCards = player.handCards.Remove(outCard)

	delete(s.actionMap, seatIndex)                       // 提前将打牌的人从行动者中删除
	s.cardsInDesktop = append(s.cardsInDesktop, outCard) // 按照打牌顺序加入桌面牌
	s.appendPeerCard(playCardType, outCard, seatIndex)

	// 先把打牌消息广播出去
	s.room.Broadcast(&outer.MahjongBTEOperaNtf{
		OpShortId: player.ShortId,
		OpType:    outer.ActionType_ActionPlayCard,
		HuType:    outer.HuType_HuTypeUnknown,
		Card:      outCard.Int32(),
	})
	log.Infow("play a card",
		"room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId, "play", outCard, "hand", player.handCards)

	// 其余三家对这张牌依次做分析
	for idx, other := range s.mahjongPlayers {
		// 跳过自己
		if seatIndex == idx {
			continue
		}

		// 提过胡牌的玩家
		if other.hu != HuInvalid {
			continue
		}

		var (
			newAction action
			pass      bool
		)
		if other.handCards.CanGangTo(outCard) {
			newAction.currentActions = append(newAction.currentActions, outer.ActionType_ActionGang)
			newAction.currentGang = append(newAction.currentGang, outCard.Int32())
			pass = true
		}
		if other.handCards.CanPongTo(outCard) {
			newAction.currentActions = append(newAction.currentActions, outer.ActionType_ActionPong)
			pass = true
		}
		if hu := other.handCards.Insert(outCard).IsHu(other.lightGang, other.darkGang, other.pong, outCard); hu != HuInvalid {
			newAction.currentActions = append(newAction.currentActions, outer.ActionType_ActionHu)
			newAction.currentHus = append(newAction.currentHus, hu.PB())
			pass = true
		}
		if pass {
			newAction.currentActions = append(newAction.currentActions, outer.ActionType_ActionPass)
		}
		if newAction.isActivated() {
			s.actionMap[idx] = &newAction // 碰杠胡的玩家加入行动组

			log.Infow("a new action",
				"room", s.room.RoomId, "seat", idx, "other", other.ShortId,
				"play", outCard, "hand", other.handCards, "action", &newAction)
		}
	}

	if len(s.actionMap) > 0 {
		s.nextAction()
		return true, outer.ERROR_OK
	}

	// 检查是否需要结算
	if s.gameOver() {
		s.SwitchTo(Settlement)
		return true, outer.ERROR_OK
	}

	// 到这里，说明出的牌没有任何人能碰杠胡，正常轮动到下家出牌，统一广播下个摸牌的人
	s.drawCard(s.nextSeatIndex(seatIndex))
	s.nextAction()
	return true, outer.ERROR_OK
}
