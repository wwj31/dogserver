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
	s.betGoldSeats = make(map[int32]int64)

	s.timeout = tools.UUID()
	expireAt := tools.Now().Add(BettingExpiration)
	pushBetSeat := make([]int32, 0)
	for seat, player := range s.niuniuPlayers {
		if player != nil && s.canPushBet(player.ShortId) == outer.ERROR_OK {
			pushBetSeat = append(pushBetSeat, int32(seat))
		}
	}

	s.room.Broadcast(&outer.NiuNiuBettingNtf{
		ExpireAt:     expireAt.UnixMilli(),
		MasterSeat:   int32(s.masterIndex),
		CanPushSeats: pushBetSeat, // 能推注的位置
	})

	s.room.AddTimer(s.timeout, expireAt, func(dt time.Duration) {
		s.RangePartInPlayer(func(seat int, player *niuniuPlayer) {
			if s.betGoldSeats[int32(seat)] == 0 {
				// 抢过庄的闲家默认是2倍,没抢过的默认是1倍
				if s.masterTimesSeats[int32(seat)] > 0 {
					s.betGoldSeats[int32(seat)] = s.baseScore() * 2
				} else {
					s.betGoldSeats[int32(seat)] = s.baseScore()
				}

				s.room.Broadcast(&outer.NiuNiuSelectBettingNtf{
					ShortId: player.ShortId,
					Gold:    s.betGoldSeats[int32(seat)],
				})
			}
		})
		s.SwitchTo(Settlement)
	})
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
		seat := int32(s.SeatIndex(shortId))
		// 发的分大于底分的5倍，算推注
		if req.Gold > s.baseScore()*5 {
			if err := s.canPushBet(shortId); err != outer.ERROR_OK {
				return err
			}
		}

		if req.Gold < player.Gold {
			return outer.ERROR_GOLD_NOT_ENOUGH
		}

		if req.Gold > s.baseScore()*int64(s.gameParams().PushBetTimes) {
			return outer.ERROR_NIUNIU_BETTING_OUT_OF_RANGE
		}

		// 抢过庄的，不能选1倍数
		if s.masterTimesSeats[seat] > 0 && req.Gold <= s.baseScore() {
			return outer.ERROR_NIUNIU_BETTING_HAS_BE_MASTER
		}

		if _, ok := s.betGoldSeats[seat]; ok {
			return outer.ERROR_NIUNIU_HAS_BE_BET
		}

		s.betGoldSeats[seat] = req.Gold
		s.room.Broadcast(&outer.NiuNiuSelectBettingNtf{
			ShortId: shortId,
			Gold:    s.betGoldSeats[seat],
		})

		if len(s.betGoldSeats) == s.participantCount() {
			s.SwitchTo(Settlement)
		}
		return &outer.NiuNiuToBettingRsp{}
	default:
		s.Log().Warnw("ready state has received an unknown message", "msg", reflect.TypeOf(req).String())
	}
	return outer.ERROR_NIUNIU_STATE_MSG_INVALID
}
