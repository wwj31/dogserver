package mahjong

import (
	"time"

	"github.com/wwj31/dogactor/tools"

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
	s.AppendPeerCard(playCardType, outCard, seatIndex)

	// 先把打牌消息广播出去
	s.room.Broadcast(&outer.MahjongBTEOperaNtf{
		OpShortId: player.ShortId,
		OpType:    outer.ActionType_ActionPlayCard,
		HuType:    outer.HuType_HuTypeUnknown,
		Card:      outCard.Int32(),
	})
	log.Infow("play a card",
		"roomId", s.room.RoomId, "player", player.ShortId, "play", outCard, "hand", player.handCards)

	var (
		actionEndAt    time.Time // 通过此时间是否为Zero，可以判断是否有人需要碰、杠、胡
		actionShortIds []int64   // 能操作的玩家加入集合
	)
	// 其余三家对这张牌依次做分析
	for idx, other := range s.mahjongPlayers {
		if seatIndex == idx { // 跳过自己
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
		if hu := other.handCards.Insert(outCard).IsHu(other.lightGang, other.darkGang, other.pong); hu != HuInvalid {
			newAction.currentActions = append(newAction.currentActions, outer.ActionType_ActionHu)
			newAction.currentHus = append(newAction.currentHus, hu.PB())
			pass = true
		}
		if pass {
			newAction.currentActions = append(newAction.currentActions, outer.ActionType_ActionPass)
		}
		if newAction.isActivated() {
			if actionEndAt.IsZero() {
				actionEndAt = tools.Now().Add(pongGangHuGuoExpire)
			}
			s.actionMap[idx] = &newAction // 碰杠胡的玩家加入行动组
			s.room.SendToPlayer(other.ShortId, &outer.MahjongBTETurnNtf{
				TotalCards:    int32(s.cards.Len()),
				ActionShortId: other.ShortId,
				ActionEndAt:   actionEndAt.UnixMilli(),
				ActionType:    newAction.currentActions,
				HuType:        newAction.currentHus,
				GangCards:     newAction.currentGang,
				NewCard:       -1, // 客户端自己取桌面牌最后一张
			})

			log.Infow("active a new action by play a card",
				"roomId", s.room.RoomId, "seat", idx, "other", other.ShortId,
				"play", outCard, "hand", other.handCards, "new action", newAction)
		}
	}

	// 行动组有人，优先让能操作的人行动, 通知剩下不能操作的人，展示"有人正在操作中..."
	if len(s.actionMap) > 0 {
		s.actionTimer(actionEndAt) // 碰,杠,胡,过,行动倒计时

		notifyPlayerMsg := &outer.MahjongBTETurnNtf{
			TotalCards:  int32(s.cards.Len()),
			ActionEndAt: actionEndAt.UnixMilli(),
		}
		s.room.Broadcast(notifyPlayerMsg, actionShortIds...)
		return true, outer.ERROR_OK
	}

	// 检查是否需要结算
	if s.cards.Len() == 0 {
		s.SwitchTo(Settlement)
		return true, outer.ERROR_OK
	}

	// 到这里，说明出的牌没有任何人能碰杠胡，正常轮动到下家出牌，统一广播下个摸牌的人
	s.drawCard(s.nextSeatIndex(seatIndex))
	return true, outer.ERROR_OK
}
