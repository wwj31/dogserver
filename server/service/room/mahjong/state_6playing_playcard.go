package mahjong

import (
	"github.com/wwj31/dogactor/logger"

	"server/proto/outermsg/outer"
)

// 打一张牌
func (s *StatePlaying) playCard(cardIndex, seatIndex int) (bool, outer.ERROR) {
	currentAction := s.getCurrentAction(seatIndex)
	if currentAction == nil || !currentAction.isValidAction(outer.ActionType_ActionPlayCard) {
		return false, outer.ERROR_MAHJONG_ACTION_PLAYER_NOT_OPERA
	}

	// 打牌的同时能胡牌，删除胡牌状态
	delete(s.canHus, seatIndex)

	player := s.mahjongPlayers[seatIndex]
	outCard := player.handCards[cardIndex]

	// 如果还有定缺花色，检查打出去的牌是否是定缺的花色
	if player.handCards.HasColorCard(player.ignoreColor) && outCard.Color() != player.ignoreColor {
		return false, outer.ERROR_MAHJONG_MUST_OUT_IGNORE_COLOR
	}

	// 把打的牌从手牌移除
	player.handCards = player.handCards.Remove(outCard)

	s.cardsInDesktop = append(s.cardsInDesktop, outCard)          // 按照打牌顺序加入桌面牌
	s.cardsPlayOrder = append(s.cardsPlayOrder, int32(seatIndex)) // 出牌座位
	s.appendPeerCard(playCardType, outCard, seatIndex, nil)

	// 先把打牌消息广播出去
	s.room.Broadcast(&outer.MahjongBTEOperaNtf{
		OpShortId:    player.ShortId,
		OpType:       outer.ActionType_ActionPlayCard,
		HandCardsNum: int32(player.handCards.Len()),
		Card:         outCard.Int32(),
		CardIndex:    int32(cardIndex),
	})

	s.Log().Color(logger.Yellow).Infow("play a card",
		"room", s.room.RoomId, "seat", seatIndex,
		"player", player.ShortId, "play", outCard,
		"hand", player.handCards, "pong", player.pong, "light gang", player.lightGang, "dark gang", player.darkGang)

	// 其余三家对这张牌依次做分析
	for otherSeat, other := range s.mahjongPlayers {
		// 跳过自己
		if seatIndex == otherSeat {
			continue
		}

		// 跳过胡牌的玩家
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

		if other.handCards.CanGangTo(outCard) && s.cards.Len() > 0 {
			newAction.acts = append(newAction.acts, outer.ActionType_ActionGang)
			newAction.gang = append(newAction.gang, outCard.Int32())
			pass = true
		}

		if other.handCards.CanPongTo(outCard) {
			newAction.acts = append(newAction.acts, outer.ActionType_ActionPong)
			pass = true
		}

		if hu := other.handCards.Insert(outCard).IsHu(other.lightGang, other.darkGang, other.pong, outCard, s.gameParams()); hu != HuInvalid {
			fan, gen, extra := s.fanGenExtra(hu, otherSeat)
			if fan > other.passHandHuFan {
				newAction.acts = append(newAction.acts, outer.ActionType_ActionHu)
				newAction.hus = append(newAction.hus, hu.PB())
				pass = true
			} else {
				s.Log().Infow("play a card trigger pass hand",
					"seat", otherSeat, "other", other.ShortId, "pass hand", other.passHandHuFan, "hu", hu, "gen", gen, "extra", extra)
			}
		}

		if pass {
			newAction.acts = append(newAction.acts, outer.ActionType_ActionPass)
		}

		if newAction.isActivated() {
			newAction.seat = otherSeat
			s.actionMap[otherSeat] = &newAction

			s.Log().Infow("a new action by play",
				"room", s.room.RoomId, "seat", otherSeat, "other", other.ShortId,
				"play", outCard, "hand", other.handCards, "action", &newAction)
		}
	}

	// 操作成功，删除行为
	delete(s.actionMap, seatIndex)
	s.removeCurrentAction(seatIndex) // 打牌操作移除当前行为

	if len(s.actionMap) == 0 {
		s.drawCard(s.nextSeatIndex(seatIndex))
	}
	s.nextAction() // 出牌后的下个行为
	return true, outer.ERROR_OK
}
