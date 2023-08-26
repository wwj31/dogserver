package mahjong

import (
	"reflect"
	"time"

	"github.com/wwj31/dogactor/logger"

	"github.com/wwj31/dogactor/tools"

	"server/proto/outermsg/outer"
)

type checkCardType int32

const (
	drawCardType checkCardType = 1 // 摸牌
	playCardType checkCardType = 2 // 打牌
	GangType1    checkCardType = 3 // 明杠(弯杠),自己摸牌，杠碰的牌(可抢杠胡)
	GangType3    checkCardType = 4 // 明杠(直杠)，别人打牌，刚好我手牌里有三张(可抢杠胡)
	GangType4    checkCardType = 5 // 暗杠,自己摸牌，自己手牌有三张
)

type (
	peerRecords struct {
		typ  checkCardType
		card Card
		seat int

		// 以下操作用于可抢杠胡并且不抢的情况下，延续执行杠操作
		afterQiangPass func(ntf *outer.MahjongBTEOperaNtf)
	}
)

type StatePlaying struct {
	*Mahjong
	actionTimerId string

	// 1.false表示可胡，但还没胡的玩家
	// 2.true表示明确操作了胡的玩家
	// 3.玩家操作过，或者碰杠，会把false的kv删掉
	canHus map[int]bool

	HusPongGang func() // 表示一炮多响期间选择碰杠的那个人，如果所有能胡的都过了，执行原本的操作
}

func (s *StatePlaying) State() int {
	return Playing
}

func (s *StatePlaying) Enter() {
	s.peerRecords = make([]peerRecords, 0)
	s.actionMap = make(map[int]*action)
	s.actionTimerId = ""
	s.currentStateEnterAt = time.Time{}
	s.canHus = make(map[int]bool)
	s.HusPongGang = nil

	// 开局默认庄家摸了一张
	s.appendPeerCard(drawCardType, s.masterCard14, s.masterIndex, nil)

	// 判断能否胡牌
	newAct := &action{seat: s.masterIndex, acts: []outer.ActionType{outer.ActionType_ActionPlayCard}}
	master := s.mahjongPlayers[s.masterIndex]
	hu := master.handCards.IsHu(master.lightGang, master.darkGang, master.pong, master.handCards[master.handCards.Len()-1], s.gameParams())

	var pass bool
	if hu != HuInvalid {
		newAct.hus = append(newAct.hus, outer.HuType(hu))
		newAct.acts = append(newAct.acts, outer.ActionType_ActionHu)
		pass = true
	}

	// 判断能否暗杠
	gangs := master.handCards.HasGang() // 检查手牌
	newAct.gang = gangs.ToSlice()
	if len(newAct.gang) > 0 {
		newAct.acts = append(newAct.acts, outer.ActionType_ActionGang)
		pass = true
	}

	if pass {
		newAct.acts = append(newAct.acts, outer.ActionType_ActionPass)
	}

	s.actionMap[s.masterIndex] = newAct
	s.Log().Infow("[Mahjong] enter state playing", "room", s.room.RoomId, "params", *s.gameParams())
	s.nextAction() // 庄家出牌
}

func (s *StatePlaying) Leave() {
	s.cancelActionTimer()
	s.Log().Infow("[Mahjong] leave state playing", "room", s.room.RoomId)
}

func (s *StatePlaying) getCurrentAction(seat int) *action {
	for _, a := range s.currentAction {
		if a.seat == seat {
			return a
		}
	}
	return nil
}

func (s *StatePlaying) removeCurrentAction(seat int) {
	for i, a := range s.currentAction {
		if a.seat == seat {
			s.currentAction = append(s.currentAction[:i], s.currentAction[i+1:]...)
			return
		}
	}
}

func (s *StatePlaying) Handle(shortId int64, v any) (result any) {
	player, seatIndex, err := s.getPlayerAndSeatE(shortId)
	if err != outer.ERROR_OK {
		return err
	}

	currentAction := s.getCurrentAction(seatIndex)
	if currentAction == nil {
		s.Log().Warnw("illegal operation", "current seat", s.currentAction, "seat", seatIndex, "player", player.ShortId)
		return outer.ERROR_MAHJONG_ACTION_PLAYER_NOT_MATCH
	}
	s.Log().Infow("playing handle msg", "shortId", shortId, reflect.TypeOf(v).String(), v)

	switch msg := v.(type) {
	case *outer.MahjongBTEPlayCardReq: // 打牌
		if msg.Index < 0 || int(msg.Index) >= player.handCards.Len() {
			return outer.ERROR_MSG_REQ_PARAM_INVALID
		}

		// 进入打牌逻辑
		if ok, errCode := s.playCard(int(msg.Index), seatIndex); !ok {
			return errCode
		}
		return &outer.MahjongBTEPlayCardRsp{AllCards: player.allCardsToPB(s.gameParams(), player.ShortId, false)}

	case *outer.MahjongBTEOperateReq: // 碰、杠、胡、过
		if errCode := s.operate(player, seatIndex, msg.ActionType, HuType(msg.Hu), Card(msg.Gang)); errCode != outer.ERROR_OK {
			return errCode
		}
		return &outer.MahjongBTEOperateRsp{AllCards: player.allCardsToPB(s.gameParams(), player.ShortId, false)}

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
func (s *StatePlaying) actionTimer(expireAt time.Time, seats ...int) {
	s.cancelActionTimer()
	s.actionTimerId = s.room.AddTimer(tools.XUID(), expireAt, func(dt time.Duration) {
		var (
			defaultOperaType outer.ActionType
			card             Card
		)

		for _, seat := range seats {
			player := s.mahjongPlayers[seat]
			act := s.actionMap[seat]
			if act == nil {
				break
			}

			// (碰杠胡过)行动者，优先打胡->杠->碰->打牌
			if act.isValidAction(outer.ActionType_ActionHu) {
				defaultOperaType = outer.ActionType_ActionHu
			} else if act.isValidAction(outer.ActionType_ActionGang) && s.cards.Len() > 0 {
				defaultOperaType = outer.ActionType_ActionGang
				card = Card(act.gang[0])
			} else if act.isValidAction(outer.ActionType_ActionPong) {
				defaultOperaType = outer.ActionType_ActionPong
				card = s.cardsInDesktop[s.cardsInDesktop.Len()-1]
			} else if act.isValidAction(outer.ActionType_ActionPlayCard) {
				var defaultPlayCard Card
				// 检查是否有定缺的牌，优先打定缺的牌,没有定缺的牌，就选手牌
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
			}

			var hu HuType
			if len(act.hus) > 0 {
				hu = HuType(act.hus[0])
			}

			s.operate(player, seat, defaultOperaType, hu, card)
		}
	})
}

func (s *StatePlaying) gameOver() bool {
	huCount := 0
	for _, p := range s.mahjongPlayers {
		if p.hu != HuInvalid {
			huCount++
			if huCount >= 3 {
				return true
			}
		}
	}

	if s.cards.Len() == 0 {
		return true
	}

	// 只要有一位玩家分<=警戒值就结束
	for _, player := range s.mahjongPlayers {
		// NOTE: 玩家每把结算后，会更新playerInfo，所以每把的GoldLine是固定的
		if player.score <= player.GetGoldLine() {
			s.scoreZeroOver = true
			return true
		}
	}

	return false
}

// 行动组有人，优先让能操作的人行动, 通知剩下不能操作的人，展示"有人正在操作中..."
func (s *StatePlaying) nextAction() {
	if len(s.actionMap) == 0 {
		return
	}

	var (
		nextActionSeats    []int
		nextActionShortIds []int64
		canHu              []int
		actionEndAt        time.Time
		defaultSeat        int
	)

	// 先把能胡的人找出来
	for seat, act := range s.actionMap {
		if act.isValidAction(outer.ActionType_ActionHu) {
			canHu = append(canHu, seat)
			s.canHus[seat] = false
		}
		defaultSeat = seat // 一边找就一边设置默认行动者，如果找不到胡的人，就直接用他,，默认就是碰杠
	}

	// 有多个人能胡，就触发一炮多响
	if len(canHu) > 0 {
		nextActionSeats = canHu
	} else {
		nextActionSeats = append(nextActionSeats, defaultSeat)
	}

	var expireDuration time.Duration
	// 广播协议
	notifyPlayerMsg := &outer.MahjongBTETurnNtf{
		TotalCards: int32(s.cards.Len()),
	}

	for _, actionSeat := range nextActionSeats {
		nextPlayer := s.mahjongPlayers[actionSeat]
		nextAct := s.actionMap[actionSeat]
		nextActionShortIds = append(nextActionShortIds, nextPlayer.ShortId)

		if nextAct.isValidAction(outer.ActionType_ActionPlayCard) {
			expireDuration = playCardExpiration
			// 打牌操作，需要广播出牌人以及出牌行为

			notifyPlayerMsg.ActionShortId = nextPlayer.ShortId
			notifyPlayerMsg.ActionType = []outer.ActionType{outer.ActionType_ActionPlayCard}
			notifyPlayerMsg.Tips = s.tips(nextPlayer)
		} else {
			expireDuration = pongGangHuGuoExpiration
		}

		s.currentAction = append(s.currentAction, nextAct)
		actionEndAt = tools.Now().Add(expireDuration)

		// 通知行动者
		s.room.SendToPlayer(nextPlayer.ShortId, &outer.MahjongBTETurnNtf{
			TotalCards:    int32(s.cards.Len()),
			ActionShortId: nextPlayer.ShortId,
			ActionEndAt:   actionEndAt.UnixMilli(),
			ActionType:    nextAct.acts,
			HuType:        nextAct.hus,
			GangCards:     nextAct.gang,
			NewCard:       nextAct.newCard.Int32(), // 客户端自己取桌面牌最后一张
			HandCards:     nextPlayer.handCards.ToSlice(),
		})

		s.Log().Color(logger.White).Infow("new action", "room", s.room.RoomId, "next seats", nextActionSeats,
			"current action", nextAct, "action map", s.actionMap, "actionEndAt", actionEndAt)
	}

	s.currentActionEndAt = actionEndAt
	notifyPlayerMsg.ActionEndAt = s.currentActionEndAt.UnixMilli()

	s.actionTimer(actionEndAt, nextActionSeats...) // 碰,杠,胡,过,行动倒计时

	s.room.Broadcast(notifyPlayerMsg, nextActionShortIds...)
}

func (s *StatePlaying) appendPeerCard(typ checkCardType, card Card, seat int, gangFn func(ntf *outer.MahjongBTEOperaNtf)) {
	s.peerRecords = append(s.peerRecords, peerRecords{
		typ:            typ,
		card:           card,
		seat:           seat,
		afterQiangPass: gangFn,
	})
}

// 胡牌了，计算总共几番，其中多少个根，是否有额外番
func (s *StatePlaying) fanGenExtra(hu HuType, seat int) (fan, gen int, extra ExtFanType) {
	extra = s.huExtra(seat)
	gen = s.huGen(seat)
	fan = huFan[hu] + extraFan[extra] + gen
	return
}

// 出牌tips
func (s *StatePlaying) tips(p *mahjongPlayer) (result []*outer.PlayCardTips) {
	// 分析每一张牌如果打出去，会叫哪些牌
	for _, playCard := range p.handCards {
		newHand := p.handCards.Remove(playCard)
		tingCards, err := newHand.ting(p.ignoreColor, p.lightGang, p.darkGang, p.pong, s.gameParams())
		if err != nil {
			s.Log().Warnw("tips ting err", "err", err)
			return nil
		}

		// 找到能胡的最大番牌
		var (
			fan    int
			huType HuType
			cards  []int32
		)

		for c, t := range tingCards {
			if huFan[t] > fan {
				fan = huFan[t]
				huType = t
				cards = nil
			}

			if t == huType {
				cards = append(cards, c.Int32())
			}
		}

		result = append(result, &outer.PlayCardTips{
			Card:      playCard.Int32(),
			HuType:    huType.PB(),
			TingCards: cards,
		})
	}
	return result
}
