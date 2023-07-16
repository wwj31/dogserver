package mahjong

import (
	"github.com/wwj31/dogactor/tools"

	"server/common/log"
	"server/proto/outermsg/outer"
)

// 摸一张牌,产生一个行动者
func (s *StatePlaying) drawCard(seatIndex int) {
	// 摸牌的时候，行动者必须是nil
	if len(s.actionMap) > 0 {
		log.Errorw("draw a card exception", "roomId", s.room.RoomId, s.actionMap)
		return
	}

	player := s.mahjongPlayers[seatIndex] // 当前摸牌者

	newCard := s.cards[0]
	s.cards = s.cards.Remove(newCard)
	player.handCards = player.handCards.Insert(newCard)
	s.AppendPeerCard(drawCardType, newCard, seatIndex)

	// 摸牌后的行为持续时间
	actionExpireAt := tools.Now().Add(playCardExpiration)
	s.actionTimer(actionExpireAt) // 出牌行动倒计时

	// 为摸牌者创建一个action
	newAction := &action{}

	// 客户端根据总牌数量是否少一张，来判断是否播摸牌动画
	notifyMsg := &outer.MahjongBTETurnNtf{
		TotalCards:    int32(s.cards.Len()),
		ActionShortId: player.ShortId,
		ActionEndAt:   actionExpireAt.UnixMilli(),
	}

	// 广播通知当前行动者(排除行动者自己)
	s.room.Broadcast(notifyMsg, player.ShortId)

	// 以下分析玩家可行的操作方式

	// 摸牌后必须出牌，所以先加入出牌操作
	newAction.currentActions = []outer.ActionType{outer.ActionType_ActionPlayCard}

	// 判断能否杠
	var gangs Cards
	gangs = player.handCards.HasGang()                   // 检查手牌
	if _, exist := player.pong[newCard.Int32()]; exist { // 检查碰牌组
		gangs = gangs.Insert(newCard)
	}
	newAction.currentGang = gangs.ToSlice()
	if len(newAction.currentGang) > 0 {
		newAction.currentActions = append(newAction.currentActions, outer.ActionType_ActionGang)
		notifyMsg.GangCards = newAction.currentGang
	}

	// 判断能否胡牌
	hu := player.handCards.IsHu(player.lightGang, player.darkGang, player.pong)
	if hu != HuInvalid {
		newAction.currentActions = append(newAction.currentActions, outer.ActionType_ActionHu)
		notifyMsg.HuType = []outer.HuType{hu.PB()}
	}
	s.actionMap[seatIndex] = newAction // 摸牌者加入行动组

	// 通知行动者自己
	notifyMsg.ActionType = newAction.currentActions
	notifyMsg.NewCard = newCard.Int32() // 摸到的新牌
	s.room.SendToPlayer(player.ShortId, notifyMsg)

	log.Infow("draw a card", "roomId", s.room.RoomId, "seat", seatIndex, "player", player.ShortId,
		"newCard", newCard, "action", newAction, "totalCards", s.cards.Len(), "hand", player.handCards)
}
