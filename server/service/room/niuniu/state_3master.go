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
	// 每位玩家初始的抢庄状态为-1，表示没有操作
	l := s.playerNumber()
	s.timesSeats = make([]int32, l, l)
	for i := 0; i < l; i++ {
		s.timesSeats[i] = -1
	}

	s.timeout = tools.UUID()
	expireAt := tools.Now().Add(MasterExpiration)
	s.room.AddTimer(s.timeout, expireAt, func(dt time.Duration) {
		for i, seat := range s.timesSeats {
			if seat == -1 {
				s.timesSeats[i] = 0
			}
		}
		s.decideMaster()
	})

	s.room.Broadcast(&outer.NiuNiuMasterNtf{ExpireAt: expireAt.UnixMilli()})
	s.Log().Infow("[NiuNiu] enter state master ", "room", s.room.RoomId)
}

func (s *StateMaster) Leave() {
	s.Log().Infow("[NiuNiu] leave state master", "room", s.room.RoomId)
}

// 确定庄家
func (s *StateMaster) decideMaster() {
	sort.Slice(s.timesSeats, func(i, j int) bool {
		return s.timesSeats[i] > s.timesSeats[j]
	})

	selects := []int{0}
	for i := 1; i < len(s.room.Players); i++ {
		if s.timesSeats[i] == s.timesSeats[0] {
			selects = append(selects, i)
		}
	}

	s.masterIndex = selects[rand.Intn(len(selects))]
	s.SwitchTo(Betting)
}
func (s *StateMaster) Handle(shortId int64, v any) (result any) {
	player, _ := s.findNiuNiuPlayer(shortId)
	if player == nil {
		s.Log().Warnw("player not in room", "roomId", s.room.RoomId, "shortId", shortId)
		return outer.ERROR_PLAYER_NOT_IN_ROOM
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

		s.timesSeats[s.SeatIndex(shortId)] = req.Times

		// 如果所有人都选择完成，就确定庄家，并且进入下个状态
		allSelected := true
		for _, val := range s.timesSeats {
			if val == -1 {
				allSelected = false
			}
		}

		if allSelected {
			s.decideMaster()
		}
		return &outer.NiuNiuToBeMasterRsp{}
	default:
		s.Log().Warnw("ready state has received an unknown message", "msg", reflect.TypeOf(req).String())
	}
	return outer.ERROR_NIUNIU_STATE_MSG_INVALID
}
