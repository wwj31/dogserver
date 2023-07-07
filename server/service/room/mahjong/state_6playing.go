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
	log.Infow("[Mahjong] enter state playing", "room", s.room.RoomId)
	s.drawCard(s.masterIndex)
}

func (s *StatePlaying) Leave() {
	log.Infow("[Mahjong] leave state playing", "room", s.room.RoomId)
}

func (s *StatePlaying) Handle(shortId int64, v any) (result any) {
	switch msg := v.(type) {
	case *outer.MahjongBTEPlayCardReq: // 打牌
		player, seatIndex := s.findMahjongPlayer(shortId)
		if player == nil {
			return outer.ERROR_PLAYER_NOT_IN_ROOM
		}

		if _, ok := s.actionMap[seatIndex]; !ok {
			return outer.ERROR_MAHJONG_ACTION_PLAYER_NOT_MATCH
		}

		if msg.Index < 0 || int(msg.Index) >= player.handCards.Len() {
			return outer.ERROR_MSG_REQ_PARAM_INVALID
		}

		// 把打的牌从手牌移除
		card := player.handCards[msg.Index]
		player.handCards = player.handCards.Remove(card)

		// 进入打牌逻辑
		if ok, errCode := s.playCard(card, player, seatIndex); !ok {
			return errCode
		}
		return &outer.MahjongBTEPlayCardRsp{AllCards: player.allCardsToPB()}

	case *outer.MahjongBTEOperateReq: // 碰、杠、胡、过
		// TODO
	}
	return nil
}

// 摸一张牌,产生一个行动者
func (s *StatePlaying) drawCard(seatIndex int) {
	player := s.mahjongPlayers[seatIndex] // 当前摸牌者

	newCard := s.cards[0]
	s.cards = s.cards.Remove(newCard)
	player.handCards.Insert(newCard)

	// 本次行为结束时间
	s.currentActionEndAt = tools.Now().Add(15 * time.Second)

	// 为摸牌者创建一个action
	newAction := &action{}

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
	newAction.currentActions = []outer.ActionType{outer.ActionType_ActionPlayCard}
	// 判断能否杠
	gangs := player.handCards.HasGang()
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
	s.actionMap[seatIndex] = newAction

	// 通知行动者自己
	notifyMsg.ActionType = newAction.currentActions
	notifyMsg.NewCard = newCard.Int32() // 摸到的新牌
	s.room.SendToPlayer(player.ShortId, notifyMsg)
}

// 打一张牌,消费一个行动者
func (s *StatePlaying) playCard(card Card, player *mahjongPlayer, seatIndex int) (bool, outer.ERROR) {
	// 检查是否是行动者,以及行为是否有效
	if act, ok := s.actionMap[seatIndex]; !ok {
		return false, outer.ERROR_MAHJONG_ACTION_PLAYER_NOT_MATCH
	} else if !act.isValidAction(outer.ActionType_ActionPlayCard) {
		return false, outer.ERROR_MAHJONG_ACTION_PLAYER_NOT_OPERA
	}

	s.latestPlayIndex = seatIndex
	// 先把打牌消息广播出去
	s.room.Broadcast(&outer.MahjongBTEOperaNtf{
		OpShortId: player.ShortId,
		OpType:    outer.ActionType_ActionPlayCard,
		HuType:    outer.HuType_HuTypeUnknown,
		Card:      card.Int32(),
	})

	// TODO 依次其余三家对这张牌做分析

	return true, outer.ERROR_OK
}

// 根据打出去的牌，决定下一个行动者
func (s *StatePlaying) nextActionIndex(playedCard Card) int64 {
	// 如果是碰，那么自己摸牌
	if playedCard == 0 {

	}
	return 0
}
