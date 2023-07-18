package mahjong

import (
	"math/rand"
	"time"

	"github.com/wwj31/dogactor/tools"

	"server/common/log"
	"server/proto/outermsg/outer"
)

type checkCardType int32

const (
	drawCardType checkCardType = 1 // 摸牌
	playCardType checkCardType = 2 // 打牌
	GangType1    checkCardType = 3 // 明杠,自己摸牌，杠碰的牌(可抢杠胡)
	GangType3    checkCardType = 4 // 明杠,直杠，别人打牌出我手牌里有三张
	GangType4    checkCardType = 5 // 暗杠,自己摸牌，自己手牌有三张
)

type (
	peerCard struct {
		typ  checkCardType
		card Card
		seat int

		// 以下操作用于杠
		afterQiangPass func() // 主要用于抢杠胡 不抢的情况下，继续执行杠的行为
	}
)

type StatePlaying struct {
	*Mahjong
	actionTimerId string
	peerCards     []peerCard // 每次操作追加操作记录
}

func (s *StatePlaying) State() int {
	return Playing
}

func (s *StatePlaying) Enter() {
	s.peerCards = make([]peerCard, 0)
	s.actionMap = make(map[int]*action)
	s.actionTimerId = ""
	log.Infow("[Mahjong] enter state playing", "room", s.room.RoomId)
	s.drawCard(s.masterIndex)
	s.nextAction()
}

func (s *StatePlaying) Leave() {
	s.cancelActionTimer()
	log.Infow("[Mahjong] leave state playing", "room", s.room.RoomId)
}

func (s *StatePlaying) Handle(shortId int64, v any) (result any) {
	player, seatIndex, err := s.getPlayerAndSeatE(shortId)
	if err != outer.ERROR_OK {
		return err
	}

	if s.currentActionSeat != seatIndex {
		log.Warnw("illegal operation", "current seat", s.currentAction, "seat", seatIndex, "player", player.ShortId)
		return outer.ERROR_MAHJONG_ACTION_PLAYER_NOT_MATCH
	}

	switch msg := v.(type) {
	case *outer.MahjongBTEPlayCardReq: // 打牌
		if msg.Index < 0 || int(msg.Index) >= player.handCards.Len() {
			return outer.ERROR_MSG_REQ_PARAM_INVALID
		}

		// 进入打牌逻辑
		if ok, errCode := s.playCard(int(msg.Index), seatIndex); !ok {
			return errCode
		}
		return &outer.MahjongBTEPlayCardRsp{AllCards: player.allCardsToPB()}

	case *outer.MahjongBTEOperateReq: // 碰、杠、胡、过
		if ok, errCode := s.operate(player, seatIndex, msg.ActionType, Card(msg.Gang)); !ok {
			return errCode
		}
		return &outer.MahjongBTEOperateRsp{AllCards: player.allCardsToPB()}

	}
	return nil
}

func (s *StatePlaying) getPlayerAndSeatE(shortId int64) (*mahjongPlayer, int, outer.ERROR) {
	player, seatIndex := s.findMahjongPlayer(shortId)
	if player == nil {
		return nil, -1, outer.ERROR_PLAYER_NOT_IN_ROOM
	}

	if _, ok := s.actionMap[seatIndex]; !ok {
		return nil, -1, outer.ERROR_MAHJONG_ACTION_PLAYER_NOT_MATCH
	}
	return player, seatIndex, outer.ERROR_OK
}

// 取消行动倒计时
func (s *StatePlaying) cancelActionTimer() {
	s.room.CancelTimer(s.actionTimerId)
}

// 行动倒计时
func (s *StatePlaying) actionTimer(expireAt time.Time, seat int) {
	s.cancelActionTimer()
	player := s.mahjongPlayers[seat]
	act := s.actionMap[seat]
	s.actionTimerId = s.room.AddTimer(tools.XUID(), expireAt, func(dt time.Duration) {
		var (
			defaultOperaType outer.ActionType
			card             Card
		)

		// (碰杠胡过)行动者，优先打胡->杠->碰->打牌
		if act.isValidAction(outer.ActionType_ActionHu) {
			defaultOperaType = outer.ActionType_ActionHu
		} else if act.isValidAction(outer.ActionType_ActionGang) {
			defaultOperaType = outer.ActionType_ActionGang
			card = Card(act.currentGang[0])
		} else if act.isValidAction(outer.ActionType_ActionPong) {
			defaultOperaType = outer.ActionType_ActionPong
			card = s.cards[s.cards.Len()-1]
		} else if act.isValidAction(outer.ActionType_ActionPlayCard) {
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
					"room", s.room.RoomId, "player", player.ShortId, "act", act,
					"hand", player.handCards, "play", defaultPlayCard)
				return
			}
			s.playCard(playIndex, seat)
			return
		} else {
			log.Warnw("action exception",
				"room", s.room.RoomId, "seat", seat, "player", player.ShortId, "act", act)
			return
		}

		s.operate(player, seat, defaultOperaType, card)
	})
}

// 行动组有人，优先让能操作的人行动, 通知剩下不能操作的人，展示"有人正在操作中..."
func (s *StatePlaying) nextAction() {
	var (
		nextSeat    int
		canHu       []int
		actionEndAt time.Time
	)

	if len(s.actionMap) == 0 {
		log.Errorw("next action error no action", "room", s.room.RoomId)
		return
	}

	// 先把能胡的人找出来
	for seat, act := range s.actionMap {
		if act.isValidAction(outer.ActionType_ActionHu) {
			canHu = append(canHu, seat)
		}
		nextSeat = seat // 一边找就一边设置下个行动者，如果找不到胡的人，就直接用他,，默认就是碰杠
	}

	// 有人能胡，就随机找一个，让他先行动
	if len(canHu) > 0 {
		nextSeat = canHu[rand.Intn(len(canHu))]
	}

	nextPlayer := s.mahjongPlayers[nextSeat]
	nextAct := s.actionMap[nextSeat]

	// 广播协议
	notifyPlayerMsg := &outer.MahjongBTETurnNtf{
		TotalCards: int32(s.cards.Len()),
	}

	var expireDuration time.Duration
	if nextAct.isValidAction(outer.ActionType_ActionPlayCard) {
		expireDuration = playCardExpiration
		// 打牌操作，需要广播出牌人以及出牌行为
		notifyPlayerMsg.ActionShortId = nextPlayer.ShortId
		notifyPlayerMsg.ActionType = []outer.ActionType{outer.ActionType_ActionPlayCard}
	} else {
		expireDuration = pongGangHuGuoExpiration
	}

	actionEndAt = tools.Now().Add(expireDuration)
	notifyPlayerMsg.ActionEndAt = actionEndAt.UnixMilli()

	s.currentActionEndAt = actionEndAt
	s.currentAction = nextAct
	s.currentActionSeat = nextSeat
	s.actionTimer(actionEndAt, nextSeat) // 碰,杠,胡,过,行动倒计时
	delete(s.actionMap, nextSeat)        // 从行动者组中删除

	// 通知行动者
	s.room.SendToPlayer(nextPlayer.ShortId, &outer.MahjongBTETurnNtf{
		TotalCards:    int32(s.cards.Len()),
		ActionShortId: nextPlayer.ShortId,
		ActionEndAt:   actionEndAt.UnixMilli(),
		ActionType:    nextAct.currentActions,
		HuType:        nextAct.currentHus,
		GangCards:     nextAct.currentGang,
		NewCard:       nextAct.newCard.Int32(), // 客户端自己取桌面牌最后一张
	})

	s.room.Broadcast(notifyPlayerMsg, nextPlayer.ShortId)
}

func (s *StatePlaying) appendPeerCard(typ checkCardType, card Card, seat int, gangFn func()) {
	s.peerCards = append(s.peerCards, peerCard{
		typ:            typ,
		card:           card,
		seat:           seat,
		afterQiangPass: gangFn,
	})
}
