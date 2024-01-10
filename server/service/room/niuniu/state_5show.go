package niuniu

import (
	"reflect"
	"time"

	"github.com/wwj31/dogactor/tools"

	"server/proto/outermsg/outer"
)

// 亮牌状态

type StateShow struct {
	*NiuNiu
	timeout string
}

func (s *StateShow) State() int {
	return ShowCards
}

func (s *StateShow) Enter() {
	s.showCards = make(map[int32]int32)

	s.timeout = tools.UUID()
	expireAt := tools.Now().Add(ShowCardsExpiration)
	s.room.AddTimer(s.timeout, expireAt, func(dt time.Duration) {
		s.RangePartInPlayer(func(seat int, player *niuniuPlayer) {
			if s.showCards[int32(seat)] == 0 {
				player.cardsGroup = player.handCards.AnalyzeCards(s.gameParams())
				s.Log().Infow("timeout show cards", "shortId", player.ShortId, "hand cards", player.handCards, "cards group", player.cardsGroup)
				s.room.Broadcast(&outer.NiuNiuFinishShowCardsNtf{
					ShortId:   player.ShortId,
					HandCards: player.handCards.ToPB(),
					CardsType: insteadJoker(player.handCards, player.cardsGroup).ToPB(),
				})
				player.checkTrusteeship(s.room)
			}
			s.showCards[int32(seat)] = 1
		})

		s.SwitchTo(Settlement)
	})

	for _, player := range s.niuniuPlayers {
		if player != nil {
			s.room.SendToPlayer(player.ShortId, &outer.NiuNiuShowCardsNtf{
				ExpireAt:  expireAt.UnixMilli(),
				HandCards: player.handCards.ToPB(),
			})
		}
	}

	s.Log().Infow("[NiuNiu] enter state Show ", "room", s.room.RoomId)
}

func (s *StateShow) Leave() {
	s.room.CancelTimer(s.timeout)
	s.Log().Infow("[NiuNiu] leave state Show", "room", s.room.RoomId)
}

func (s *StateShow) Handle(shortId int64, v any) (result any) {
	player, _ := s.findNiuNiuPlayer(shortId)
	if player == nil || !player.ready {
		s.Log().Warnw("player not in room", "roomId", s.room.RoomId, "shortId", shortId)
		return outer.ERROR_NIUNIU_NOT_IN_GAMING
	}

	switch req := v.(type) {
	case *outer.NiuNiuShowCardsReq: // 亮牌
		if player.trusteeship {
			return outer.ERROR_ROOM_NEED_CANCEL_TRUSTEESHIP
		}

		s.showCards[int32(s.SeatIndex(shortId))] = 1
		player.cardsGroup = player.handCards.AnalyzeCards(s.gameParams())
		player.timeoutTrusteeshipCount = 0
		s.room.Broadcast(&outer.NiuNiuFinishShowCardsNtf{
			ShortId:   shortId,
			HandCards: player.handCards.ToPB(),
			CardsType: insteadJoker(player.handCards, player.cardsGroup).ToPB(),
		})
		s.Log().Infow("player show cards", "shortId", shortId, "hand cards", player.handCards, "cards group", player.cardsGroup)

		if len(s.showCards) == s.participantCount() {
			s.SwitchTo(Settlement)
		}

		return &outer.NiuNiuShowCardsRsp{}
	default:
		s.Log().Warnw("ready state has received an unknown message", "msg", reflect.TypeOf(req).String())
	}
	return outer.ERROR_NIUNIU_STATE_MSG_INVALID
}

func insteadJoker(handCards PokerCards, cardsGroup CardsGroup) (result CardsGroup) {
	result = cardsGroup
	var jokers PokerCards
	for _, card := range handCards {
		if card.Color() == Joker {
			jokers = append(jokers, card)
		}
	}

	if len(jokers) == 0 {
		return
	}

	var dmpCards PokerCards
	// 找到dmpCards中存在的一张牌，并且删除
	dmpDelete := func(targetCard PokerCard) bool {
		for i, card := range dmpCards {
			if card == targetCard {
				dmpCards = append(dmpCards[:i], dmpCards[i+1:]...)
				return true
			}
		}
		return false
	}

	copy(dmpCards, handCards)
	for i, card := range result.Cards {
		if !dmpDelete(card) {
			result.Cards[i] = jokers[0]
			jokers = jokers[1:]
			if len(jokers) == 0 {
				return
			}
		}
	}

	copy(dmpCards, handCards)
	for i, card := range result.SideCards {
		if !dmpDelete(card) {
			result.SideCards[i] = jokers[0]
			jokers = jokers[1:]
			if len(jokers) == 0 {
				return
			}
		}
	}

	return
}
