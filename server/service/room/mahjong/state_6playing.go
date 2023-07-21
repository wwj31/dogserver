package mahjong

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/wwj31/dogactor/tools"

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
		afterQiangPass func()          // 主要用于抢杠胡 不抢的情况下，继续执行杠的行为
		loseScores     map[int64]int64 // 杠的赔分
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
	s.currentStateEnterAt = time.Time{}

	// 判断能否胡牌
	newAct := &action{acts: []outer.ActionType{outer.ActionType_ActionPlayCard}}
	master := s.mahjongPlayers[s.masterIndex]
	hu := master.handCards.IsHu(master.lightGang, master.darkGang, master.pong, master.handCards[master.handCards.Len()-1])
	if hu != HuInvalid {
		newAct.hus = append(newAct.hus, outer.HuType(hu))
		newAct.acts = append(newAct.acts, outer.ActionType_ActionHu)
	}
	s.currentAction = newAct
	s.currentActionSeat = s.masterIndex
	s.actionMap[s.masterIndex] = s.currentAction
	s.Log().Infow("[Mahjong] enter state playing", "room", s.room.RoomId)
	s.nextAction() // 庄家出牌
}

func (s *StatePlaying) Leave() {
	s.cancelActionTimer()
	s.Log().Infow("[Mahjong] leave state playing", "room", s.room.RoomId)
}

func (s *StatePlaying) Handle(shortId int64, v any) (result any) {
	player, seatIndex, err := s.getPlayerAndSeatE(shortId)
	if err != outer.ERROR_OK {
		return err
	}

	if s.currentActionSeat != seatIndex {
		s.Log().Warnw("illegal operation", "current seat", s.currentAction, "seat", seatIndex, "player", player.ShortId)
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
		if ok, errCode := s.operate(player, seatIndex, msg.ActionType, HuType(msg.Hu), Card(msg.Gang)); !ok {
			return errCode
		}
		return &outer.MahjongBTEOperateRsp{AllCards: player.allCardsToPB()}

	default:
		s.Log().Warnw("playing status has received an unknown message", "msg", reflect.TypeOf(msg).String())
	}
	return outer.ERROR_MAHJONG_STATE_MSG_INVALID
}

func (s *StatePlaying) getPlayerAndSeatE(shortId int64) (*mahjongPlayer, int, outer.ERROR) {
	player, seatIndex := s.findMahjongPlayer(shortId)
	if player == nil {
		return nil, -1, outer.ERROR_PLAYER_NOT_IN_ROOM
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
			card = Card(act.gang[0])
		} else if act.isValidAction(outer.ActionType_ActionPong) {
			defaultOperaType = outer.ActionType_ActionPong
			card = s.cardsInDesktop[s.cardsInDesktop.Len()-1]
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
				s.Log().Errorw("action timeout, playIndex == -1",
					"room", s.room.RoomId, "player", player.ShortId, "act", act,
					"hand", player.handCards, "play", defaultPlayCard)
				return
			}
			s.playCard(playIndex, seat)
			return
		} else {
			s.Log().Warnw("action exception",
				"room", s.room.RoomId, "seat", seat, "player", player.ShortId, "act", act)
			return
		}

		var hu HuType
		if len(act.hus) > 0 {
			hu = HuType(act.hus[0])
		}

		s.operate(player, seat, defaultOperaType, hu, card)
	})
}

// 行动组有人，优先让能操作的人行动, 通知剩下不能操作的人，展示"有人正在操作中..."
func (s *StatePlaying) nextAction() {
	if len(s.actionMap) == 0 {
		return
	}

	var (
		nextSeat    int
		canHu       []int
		actionEndAt time.Time
	)

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

	s.Log().Infow("next action", "room", s.room.RoomId, "act player", nextPlayer.ShortId, "seat", nextSeat,
		"current action", s.currentAction, "action map", s.actionMap)

	// 通知行动者
	s.room.SendToPlayer(nextPlayer.ShortId, &outer.MahjongBTETurnNtf{
		TotalCards:    int32(s.cards.Len()),
		ActionShortId: nextPlayer.ShortId,
		ActionEndAt:   actionEndAt.UnixMilli(),
		ActionType:    nextAct.acts,
		HuType:        nextAct.hus,
		GangCards:     nextAct.gang,
		NewCard:       nextAct.newCard.Int32(), // 客户端自己取桌面牌最后一张
	})

	s.room.Broadcast(notifyPlayerMsg, nextPlayer.ShortId)
}

func (s *StatePlaying) appendPeerCard(typ checkCardType, card Card, seat int, gangFn func(), loseScore map[int64]int64) {
	s.peerCards = append(s.peerCards, peerCard{
		typ:            typ,
		card:           card,
		seat:           seat,
		afterQiangPass: gangFn,
		loseScores:     loseScore,
	})
}

// 检查是否一炮多响
func (s *StatePlaying) checkMutilHu(huPeerIndex int) bool {
	if huPeerIndex == 0 {
		return false
	}
	for _, player := range s.mahjongPlayers {
		if player.huPeerIndex != 0 && player.huPeerIndex == huPeerIndex {
			return true
		}
	}
	return false
}

// 获得排除了某些座位后，剩余的座位
func (s *StatePlaying) allSeat(ignoreSeat ...int) (result []int) {
	seatMap := map[int]struct{}{}
	for _, seat := range ignoreSeat {
		seatMap[seat] = struct{}{}
	}

	for seatIndex := 0; seatIndex < 4; seatIndex++ {
		if _, ignore := seatMap[seatIndex]; !ignore {
			result = append(result, seatIndex)
		}
	}

	return result
}
