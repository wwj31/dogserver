package mahjong

import (
	"server/proto/outermsg/outer"
)

// 摸一张牌,产生一个行动者
func (s *StatePlaying) drawCard(seatIndex int) {
	// 摸牌的时候，行动者必须是nil
	if len(s.actionMap) > 0 {
		s.Log().Errorw("draw a card exception", "room", s.room.RoomId, s.actionMap)
		return
	}

	// 每次摸牌先判断是否结算
	if s.gameOver() {
		s.SwitchTo(Settlement)
		return
	}

	newCard := s.cards[0]
	s.cards = s.cards.Remove(newCard)
	player := s.mahjongPlayers[seatIndex] // 当前摸牌者
	player.handCards = player.handCards.Insert(newCard)
	s.appendPeerCard(drawCardType, newCard, seatIndex, nil)

	// 摸牌后必须出牌，所以先加入出牌操作
	newAction := &action{seat: seatIndex}
	newAction.acts = []outer.ActionType{outer.ActionType_ActionPlayCard}
	newAction.newCard = newCard

	// 判断能否杠
	var gangs Cards
	gangs = player.handCards.HasGang()                   // 检查手牌
	if _, exist := player.pong[newCard.Int32()]; exist { // 检查碰牌组
		gangs = gangs.Insert(newCard)
	}

	newAction.gang = gangs.ToSlice()
	if len(newAction.gang) > 0 {
		newAction.acts = append(newAction.acts, outer.ActionType_ActionGang)
	}

	// 判断能否胡牌
	hu := player.handCards.IsHu(player.lightGang, player.darkGang, player.pong, newCard, s.gameParams())
	if hu != HuInvalid {
		newAction.acts = append(newAction.acts, outer.ActionType_ActionHu)
	}
	s.actionMap[seatIndex] = newAction // 摸牌者加入行动组

	s.Log().Infow("draw a card", "room", s.room.RoomId, "seat", seatIndex, "player", player.ShortId,
		"newCard", newCard, "action", newAction, "totalCards", s.cards.Len(), "hand", player.handCards)
}
