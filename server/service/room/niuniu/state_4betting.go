package niuniu

import (
	"reflect"
	"time"

	"github.com/wwj31/dogactor/tools"

	"server/proto/outermsg/outer"
)

// 压注状态

type StateBetting struct {
	*NiuNiu
	timeout string
}

func (s *StateBetting) State() int {
	return Betting
}

func (s *StateBetting) Enter() {
	s.betTimesSeats = make(map[int32]int32)

	s.timeout = tools.UUID()
	expireAt := tools.Now().Add(BettingExpiration)
	s.room.AddTimer(s.timeout, expireAt, func(dt time.Duration) {
		for seat, player := range s.niuniuPlayers {
			if player == nil || !player.ready {
				continue
			}

			times := s.betTimesSeats[int32(seat)]
			if times == 0 {
				// 抢过庄的闲家默认是2倍,没抢过的默认是1倍
				if s.masterTimesSeats[int32(seat)] > 0 {
					s.betTimesSeats[int32(seat)] = 2
				} else {
					s.betTimesSeats[int32(seat)] = 1
				}

				s.room.Broadcast(&outer.NiuNiuSelectBettingNtf{
					ShortId: player.ShortId,
					Times:   s.betTimesSeats[int32(seat)],
				})
			}
		}
		s.SwitchTo(Settlement)
	})
	s.room.Broadcast(&outer.NiuNiuBettingNtf{ExpireAt: expireAt.UnixMilli()})
	s.Log().Infow("[NiuNiu] enter state Betting ", "room", s.room.RoomId)
}

func (s *StateBetting) Leave() {
	s.Log().Infow("[NiuNiu] leave state Betting", "room", s.room.RoomId)
}

func (s *StateBetting) Handle(shortId int64, v any) (result any) {
	player, _ := s.findNiuNiuPlayer(shortId)
	if player == nil {
		s.Log().Warnw("player not in game", "roomId", s.room.RoomId, "shortId", shortId)
		return outer.ERROR_NIUNIU_NOT_IN_GAMING
	}

	switch req := v.(type) {
	case *outer.NiuNiuToBettingReq: // 押注
		if req.Times < 0 || req.Times > 5 {
			return outer.ERROR_NIUNIU_BETTING_OUT_OF_RANGE
		}

		// 抢过庄的，不能选1倍数
		seat := int32(s.SeatIndex(shortId))
		if s.masterTimesSeats[seat] > 0 && req.Times == 1 {
			return outer.ERROR_NIUNIU_BETTING_HAS_BE_MASTER
		}

		if _, ok := s.betTimesSeats[seat]; ok {
			return outer.ERROR_NIUNIU_HAS_BE_BET
		}

		s.betTimesSeats[seat] = req.Times
		s.room.Broadcast(&outer.NiuNiuSelectBettingNtf{
			ShortId: shortId,
			Times:   s.betTimesSeats[seat],
		})

		if len(s.betTimesSeats) == s.participantCount() {
			s.SwitchTo(Settlement)
		}
		return &outer.NiuNiuToBettingRsp{}
	default:
		s.Log().Warnw("ready state has received an unknown message", "msg", reflect.TypeOf(req).String())
	}
	return outer.ERROR_NIUNIU_STATE_MSG_INVALID
}
