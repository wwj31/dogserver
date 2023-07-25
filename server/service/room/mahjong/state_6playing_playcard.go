package mahjong

import (
	"server/proto/outermsg/outer"
)

// 打一张牌
func (s *StatePlaying) playCard(cardIndex, seatIndex int) (bool, outer.ERROR) {
	if !s.currentAction.isValidAction(outer.ActionType_ActionPlayCard) {
		return false, outer.ERROR_MAHJONG_ACTION_PLAYER_NOT_OPERA
	}

	// 把打的牌从手牌移除
	player := s.mahjongPlayers[seatIndex]
	outCard := player.handCards[cardIndex]

	// 如果还有定缺花色，检查打出去的牌是否是定缺的花色
	if player.handCards.HasColorCard(player.ignoreColor) && outCard.Color() != player.ignoreColor {
		return false, outer.ERROR_MAHJONG_MUST_OUT_IGNORE_COLOR
	}
	player.handCards = player.handCards.Remove(outCard)

	s.cardsInDesktop = append(s.cardsInDesktop, outCard)          // 按照打牌顺序加入桌面牌
	s.cardsPlayOrder = append(s.cardsPlayOrder, int32(seatIndex)) // 出牌座位
	s.appendPeerCard(playCardType, outCard, seatIndex, nil)

	// 先把打牌消息广播出去
	s.room.Broadcast(&outer.MahjongBTEOperaNtf{
		OpShortId: player.ShortId,
		OpType:    outer.ActionType_ActionPlayCard,
		Card:      outCard.Int32(),
	})
	s.Log().Infow("play a card",
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

		// 定缺花色的牌，直接跳过
		if outCard.Color() == other.ignoreColor {
			continue
		}

		var (
			newAction action
			pass      bool
		)

		if other.handCards.CanGangTo(outCard) {
			newAction.acts = append(newAction.acts, outer.ActionType_ActionGang)
			newAction.gang = append(newAction.gang, outCard.Int32())
			pass = true
		}

		if other.handCards.CanPongTo(outCard) {
			newAction.acts = append(newAction.acts, outer.ActionType_ActionPong)
			pass = true
		}

		if hu := other.handCards.Insert(outCard).IsHu(other.lightGang, other.darkGang, other.pong, outCard, s.gameParams()); hu != HuInvalid {
			newAction.acts = append(newAction.acts, outer.ActionType_ActionHu)
			newAction.hus = append(newAction.hus, hu.PB())
			pass = true
		}

		if pass {
			newAction.acts = append(newAction.acts, outer.ActionType_ActionPass)
		}

		if newAction.isActivated() {
			s.actionMap[idx] = &newAction // 碰杠胡的玩家加入行动组

			s.Log().Infow("a new action",
				"room", s.room.RoomId, "seat", idx, "other", other.ShortId,
				"play", outCard, "hand", other.handCards, "action", &newAction)
		}
	}

	if len(s.actionMap) == 0 {
		s.drawCard(s.nextSeatIndex(seatIndex))
	}
	s.nextAction() // 出牌后的下个行为

	return true, outer.ERROR_OK
}
