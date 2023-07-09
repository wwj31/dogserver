package mahjong

import (
	"time"

	"github.com/wwj31/dogactor/tools"

	"server/common/log"
	"server/proto/outermsg/outer"
)

// 游戏状态
const (
	pongGangHuGuoExpire = 10 * time.Second // 碰、杠、胡、过持续时间
	playCardExpire      = 15 * time.Second // 摸牌后的行为持续时间(出牌，杠，胡)
)

type StatePlaying struct {
	*Mahjong
	actionTimerId string
}

func (s *StatePlaying) State() int {
	return Playing
}

func (s *StatePlaying) Enter() {
	log.Infow("[Mahjong] enter state playing", "room", s.room.RoomId)
	s.drawCard(s.masterIndex)
}

func (s *StatePlaying) Leave() {
	s.cancelActionTimer()
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

		// 进入打牌逻辑
		if ok, errCode := s.playCard(int(msg.Index), seatIndex); !ok {
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
	// 摸牌的时候，行动者必须是nil
	if len(s.actionMap) > 0 {
		log.Errorw("draw a card exception", "roomId", s.room.RoomId, s.actionMap)
		return
	}

	player := s.mahjongPlayers[seatIndex] // 当前摸牌者

	newCard := s.cards[0]
	s.cards = s.cards.Remove(newCard)
	player.handCards.Insert(newCard)

	// 摸牌后的行为持续时间
	s.currentActionEndAt = tools.Now().Add(playCardExpire)

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

	// 摸牌后必须出牌，所以先加入出牌操作
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
	s.actionMap[seatIndex] = newAction // 摸牌者加入行动组

	// 通知行动者自己
	notifyMsg.ActionType = newAction.currentActions
	notifyMsg.NewCard = newCard.Int32() // 摸到的新牌
	s.room.SendToPlayer(player.ShortId, notifyMsg)

	s.actionTimer() // 出牌行动倒计时
	log.Infow("draw a card", "roomId", s.room.RoomId, "seatIndex", seatIndex,
		"newAction", newAction, "newCard", newCard, "current hand", player.handCards)
}

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
	playCard := player.handCards[cardIndex]
	player.handCards = player.handCards.Remove(playCard)

	s.latestPlayIndex = seatIndex                         // 设置最后一次出牌的玩家
	delete(s.actionMap, seatIndex)                        // 提前将打牌的人从行动者中删除
	s.cardsInDesktop = append(s.cardsInDesktop, playCard) // 按照打牌顺序加入桌面牌

	// 先把打牌消息广播出去
	s.room.Broadcast(&outer.MahjongBTEOperaNtf{
		OpShortId: player.ShortId,
		OpType:    outer.ActionType_ActionPlayCard,
		HuType:    outer.HuType_HuTypeUnknown,
		Card:      playCard.Int32(),
	})
	log.Infow("play a card",
		"roomId", s.room.RoomId, "player", player.ShortId, "play", playCard, "hand", player.handCards)

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
		if other.handCards.CanGangTo(playCard) {
			newAction.currentActions = append(newAction.currentActions, outer.ActionType_ActionGang)
			newAction.currentGang = append(newAction.currentGang, playCard.Int32())
			pass = true
		}
		if other.handCards.CanPongTo(playCard) {
			newAction.currentActions = append(newAction.currentActions, outer.ActionType_ActionPong)
			pass = true
		}
		if hu := other.handCards.IsHu(other.lightGang, other.darkGang, other.pong); hu != HuInvalid {
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

			log.Infow("active a new action by play",
				"roomId", s.room.RoomId, "seat", idx, "other", other.ShortId,
				"play", playCard, "hand", other.handCards, "new action", newAction)
		}
	}

	// 行动组有人，优先让能操作的人行动, 通知剩下不能操作的人，展示"有人正在操作中..."
	if len(s.actionMap) > 0 {
		s.currentActionEndAt = actionEndAt
		notifyPlayerMsg := &outer.MahjongBTETurnNtf{
			TotalCards:  int32(s.cards.Len()),
			ActionEndAt: s.currentActionEndAt.UnixMilli(),
		}
		s.room.Broadcast(notifyPlayerMsg, actionShortIds...)
		s.actionTimer() // 碰,杠,胡,过,行动倒计时
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

// 取消行动倒计时
func (s *StatePlaying) cancelActionTimer() {
	s.room.CancelTimer(s.actionTimerId)
}

// 行动倒计时
func (s *StatePlaying) actionTimer() {
	s.cancelActionTimer()
	s.actionTimerId = s.room.AddTimer(tools.XUID(), s.currentActionEndAt, func(dt time.Duration) {
		if len(s.actionMap) == 0 {
			// 所有行动计时器都会被正常释放，无行动者超时属于异常
			log.Warnw("action timer timeout len == 0")
			return
		}

		// 根据能执行的行为，默认为玩家操作
		for seatIndex, act := range s.actionMap {
			player := s.mahjongPlayers[seatIndex]
			// 出牌人，只可能有一个行动者
			if act.isValidAction(outer.ActionType_ActionPlayCard) {
				var defaultPlayCard Card

				// 优先打定缺花色,没有定缺花色的牌，就选手牌
				ignoreCards := player.handCards.colorCards(player.ignoreColor)
				if ignoreCards.Len() != 0 {
					defaultPlayCard = ignoreCards.Random()
				} else {
					defaultPlayCard = player.handCards.Random()
				}

				// 把超时后随机选的牌打出去
				playIndex := player.handCards.CardIndex(defaultPlayCard)
				if playIndex == -1 {
					log.Errorw("action timeout, playIndex == -1",
						"roomId", s.room.RoomId, "player", player.ShortId, "act", act,
						"hand", player.handCards, "play", defaultPlayCard)
					return
				}
				s.playCard(playIndex, seatIndex)
				break
			} else {
				// 碰杠胡过行动者，优先打胡->杠->碰
				// TODO ...
			}
		}
	})
}
