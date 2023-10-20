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
	s.shows = make(map[int32]int32)

	s.timeout = tools.UUID()
	expireAt := tools.Now().Add(ShowCardsExpiration)
	s.room.AddTimer(s.timeout, expireAt, func(dt time.Duration) {
		for seat, v := range s.shows {
			if v == 0 {
				s.room.Broadcast(&outer.NiuNiuFinishShowCardsNtf{
					ShortId:   s.niuniuPlayers[seat].ShortId,
					HandCards: s.niuniuPlayers[seat].handCards.ToPB(),
				})
			}
		}
		s.SwitchTo(Settlement)
	})

	for _, player := range s.niuniuPlayers {
		s.room.SendToPlayer(player.ShortId, &outer.NiuNiuShowCardsNtf{
			ExpireAt:  expireAt.UnixMilli(),
			HandCards: player.handCards.ToPB(),
		})
	}

	s.Log().Infow("[NiuNiu] enter state Show ", "room", s.room.RoomId)
}

func (s *StateShow) Leave() {
	s.Log().Infow("[NiuNiu] leave state Show", "room", s.room.RoomId)
}

func (s *StateShow) Handle(shortId int64, v any) (result any) {
	player, _ := s.findNiuNiuPlayer(shortId)
	if player == nil {
		s.Log().Warnw("player not in room", "roomId", s.room.RoomId, "shortId", shortId)
		return outer.ERROR_NIUNIU_NOT_IN_GAMING
	}

	switch req := v.(type) {
	case *outer.NiuNiuShowCardsReq: // 亮牌
		s.shows[int32(s.SeatIndex(shortId))] = 1
		s.room.Broadcast(&outer.NiuNiuFinishShowCardsNtf{
			ShortId:   shortId,
			HandCards: player.handCards.ToPB(),
		})

		if len(s.shows) == len(s.niuniuPlayers) {
			s.SwitchTo(Settlement)
		}

		return &outer.NiuNiuShowCardsRsp{}
	default:
		s.Log().Warnw("ready state has received an unknown message", "msg", reflect.TypeOf(req).String())
	}
	return outer.ERROR_NIUNIU_STATE_MSG_INVALID
}
