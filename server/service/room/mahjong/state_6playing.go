package mahjong

import (
	"time"

	"github.com/wwj31/dogactor/tools"

	"server/common/log"
	"server/proto/outermsg/outer"
)

// 游戏状态

type StatePlaying struct {
	*Mahjong
}

func (s *StatePlaying) State() int {
	return Playing
}

func (s *StatePlaying) Enter() {
	s.currentActionIndex = s.masterIndex
	log.Infow("[Mahjong] enter state playing", "room", s.room.RoomId)
}

func (s *StatePlaying) Leave() {
	log.Infow("[Mahjong] leave state playing", "room", s.room.RoomId)
}

func (s *StatePlaying) Handle(shortId int64, v any) (result any) {
	return nil
}

func (s *StatePlaying) currentIndex() int {
	if s.currentActionIndex < 0 || s.currentActionIndex >= maxNum {
		log.Errorw("current action index out of range",
			"roomId", s.room.RoomId, "index", s.currentActionIndex)
		s.currentActionIndex = 0
	}

	return s.currentActionIndex
}

// 摸一张牌
func (s *StatePlaying) drawCard() {
	s.latestDrawIndex = s.currentIndex()
	player := s.mahjongPlayers[s.latestDrawIndex] // 当前摸牌者
	newCard := s.cards[0]
	s.cards = s.cards.Remove(newCard)
	player.handCards.Insert(newCard)
	s.currentActionEndAt = tools.Now().Add(15 * time.Second)

	// 客户端根据总牌数量是否少一张，来判断是否播摸牌动画
	notifyMsg := &outer.MahjongBTETurnNtf{
		TotalCards:    int32(s.cards.Len()),
		ActionShortId: player.ShortId,
		ActionEndAt:   s.currentActionEndAt.UnixMilli(),
	}

	// 广播通知当前行动者(排除行动者自己)
	s.room.Broadcast(notifyMsg, player.ShortId)

	// 以下分析玩家可行的操作方式

	// 摸牌就要打牌，先加入出牌操作
	s.currentActions = []outer.ActionType{outer.ActionType_ActionPlayCard}
	// 判断能否杠
	gangs := player.handCards.HasGang()
	s.currentGang = gangs.ToSlice()
	if len(s.currentGang) > 0 {
		s.currentActions = append(s.currentActions, outer.ActionType_ActionGang)
		notifyMsg.GangCards = s.currentGang
	}
	// 判断能否胡牌
	hu := player.handCards.IsHu(player.lightGang, player.darkGang, player.pong)
	if hu != HuInvalid {
		s.currentActions = append(s.currentActions, outer.ActionType_ActionHu)
		notifyMsg.HuType = []outer.HuType{hu.PB()}
	}

	// 通知行动者自己
	notifyMsg.ActionType = s.currentActions
	notifyMsg.NewCard = newCard.Int32() // 摸到的新牌
	s.room.SendToPlayer(player.ShortId, notifyMsg)
}

// 根据打出去的牌，决定下一个行动者
func (s *StatePlaying) nextActionIndex(playedCard Card) int64 {
	// 如果是碰，那么自己摸牌
	if playedCard == 0 {

	}
	return 0
}
