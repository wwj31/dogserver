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
	l := s.playerNumber()
	s.betSeats = make([]int32, l, l)

	s.timeout = tools.UUID()
	expireAt := tools.Now().Add(BettingExpiration)
	s.room.AddTimer(s.timeout, expireAt, func(dt time.Duration) {
		for seat, times := range s.betSeats {
			if times == 0 {
				if s.timesSeats[seat] > 0 {
					s.betSeats[seat] = 2
				} else {
					s.betSeats[seat] = 1
				}

				s.room.Broadcast(&outer.NiuNiuSelectBettingNtf{
					ShortId: s.niuniuPlayers[seat].ShortId,
					Times:   s.betSeats[seat],
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
		s.Log().Warnw("player not in room", "roomId", s.room.RoomId, "shortId", shortId)
		return outer.ERROR_PLAYER_NOT_IN_ROOM
	}

	switch req := v.(type) {
	case *outer.NiuNiuToBettingReq: // 押注
		if req.Times < 0 || req.Times > 5 {
			return outer.ERROR_NIUNIU_BETTING_OUT_OF_RANGE
		}

		seat := s.SeatIndex(shortId)
		if s.timesSeats[seat] > 0 && req.Times == 1 {
			return outer.ERROR_NIUNIU_BETTING_HAS_BE_MASTER
		}

		s.betSeats[seat] = req.Times

		allSelected := true
		for _, betVal := range s.betSeats {
			if betVal == 0 {
				allSelected = false
			}
		}

		if allSelected {
			s.SwitchTo(Settlement)
		}

		return &outer.NiuNiuToBettingRsp{}
	default:
		s.Log().Warnw("ready state has received an unknown message", "msg", reflect.TypeOf(req).String())
	}
	return outer.ERROR_NIUNIU_STATE_MSG_INVALID
}
