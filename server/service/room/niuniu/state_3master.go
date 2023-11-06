package niuniu

import (
	"math/rand"
	"reflect"
	"sort"
	"time"

	"github.com/wwj31/dogactor/tools"

	"server/proto/outermsg/outer"
)

// 抢庄状态

type StateMaster struct {
	*NiuNiu
	timeout string
}

func (s *StateMaster) State() int {
	return DecideMaster
}

func (s *StateMaster) Enter() {
	s.masterTimesSeats = make(map[int32]int32)
	s.RangePartInPlayer(func(seat int, player *niuniuPlayer) { s.masterTimesSeats[int32(seat)] = -1 })

	s.timeout = tools.UUID()
	expireAt := tools.Now().Add(MasterExpiration)
	s.room.Broadcast(&outer.NiuNiuMasterNtf{ExpireAt: expireAt.UnixMilli()})

	s.room.AddTimer(s.timeout, expireAt, func(dt time.Duration) {
		for seat, v := range s.masterTimesSeats {
			if v == -1 {
				s.masterTimesSeats[seat] = 0
				s.room.Broadcast(&outer.NiuNiuSelectMasterNtf{
					ShortId: s.niuniuPlayers[seat].ShortId,
					Times:   0,
				})
			}
		}
		s.decideMaster()
	})

	s.Log().Infow("[NiuNiu] enter state master ", "room", s.room.RoomId)
}

func (s *StateMaster) Leave() {
	s.room.CancelTimer(s.timeout)
	s.Log().Infow("[NiuNiu] leave state master",
		"room", s.room.RoomId,
		"master", s.niuniuPlayers[s.masterIndex].ShortId,
		"master times", s.masterTimesSeats,
		"push bet index", s.pushBetIndex,
	)
}

// 确定庄家
func (s *StateMaster) decideMaster() {
	var arr []struct {
		Seat  int32
		Times int32
	}
	for seat, times := range s.masterTimesSeats {
		arr = append(arr, struct {
			Seat  int32
			Times int32
		}{Seat: seat, Times: times})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Times > arr[j].Times
	})

	selects := []int{int(arr[0].Seat)}
	for i := 1; i < len(arr); i++ {
		if arr[i].Times == arr[0].Times {
			selects = append(selects, int(arr[i].Seat))
		}
	}

	randSeat := rand.Intn(len(selects))
	s.masterIndex = selects[randSeat]
	if arr[0].Times > 0 {
		s.pushBetIndex = append(selects[:randSeat], selects[randSeat+1:]...)
	}
	s.SwitchTo(Betting)
}

func (s *StateMaster) Handle(shortId int64, v any) (result any) {
	player, _ := s.findNiuNiuPlayer(shortId)
	if player == nil {
		s.Log().Warnw("player not in game", "roomId", s.room.RoomId, "shortId", shortId)
		return outer.ERROR_NIUNIU_NOT_IN_GAMING
	}

	switch req := v.(type) {
	case *outer.NiuNiuToBeMasterReq:
		minGoldRequired := s.baseScore() * 200
		if player.Gold < minGoldRequired {
			return outer.ERROR_NIUNIU_GOLD_NOT_ENOUGH_MASTER
		}

		if req.Times < 0 || req.Times > 4 {
			return outer.ERROR_NIUNIU_MASTER_OUT_OF_RANGE
		}

		seat := int32(s.SeatIndex(shortId))
		if s.masterTimesSeats[seat] != -1 {
			return outer.ERROR_NIUNIU_HAS_BE_MASTER
		}

		s.masterTimesSeats[seat] = req.Times
		s.room.Broadcast(&outer.NiuNiuSelectMasterNtf{
			ShortId: player.ShortId,
			Times:   req.Times,
		})

		switchToNext := true
		for _, val := range s.masterTimesSeats {
			if val == -1 {
				switchToNext = false
			}
		}
		s.Log().Infow("NiuNiuToBeMasterReq", "shortId", shortId, "seat", seat, "req", req.String(), "switch", switchToNext)

		// 如果所有人都选择完成，就确定庄家，并且进入下个状态
		if switchToNext {
			s.decideMaster()
		}
		return &outer.NiuNiuToBeMasterRsp{}
	default:
		s.Log().Warnw("ready state has received an unknown message", "msg", reflect.TypeOf(req).String())
	}
	return outer.ERROR_NIUNIU_STATE_MSG_INVALID
}
