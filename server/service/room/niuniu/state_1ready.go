package niuniu

import (
	"time"

	"github.com/wwj31/dogactor/tools"

	"server/proto/outermsg/outer"
)

// 准备状态

type StateReady struct {
	*NiuNiu
	timeId string
}

func (s *StateReady) State() int {
	return Ready
}

func (s *StateReady) Enter() {
	s.onPlayerEnter = s.playerEnter
	s.onPlayerLeave = s.playerLeave

	s.RangePartInPlayer(func(seat int, player *niuniuPlayer) {
		if player.Gold <= 0 {
			s.room.PlayerLeave(player.ShortId, true)
		}
	})
	s.Log().Infow("[NiuNiu] enter state ready ", "room", s.room.RoomId)
	s.room.Broadcast(&outer.NiuNiuReadyNtf{})
}

func (s *StateReady) Leave() {
	s.onPlayerEnter = nil
	s.onPlayerLeave = nil
	s.Log().Infow("[NiuNiu] leave state ready", "room", s.room.RoomId)
}

func (s *StateReady) Handle(shortId int64, v any) (result any) {
	return outer.ERROR_NIUNIU_STATE_MSG_INVALID
}

func (s *StateReady) playerEnter(player *niuniuPlayer) {
	if s.playerCount() >= s.gameParams().MinPlayPlayerCount && s.timeId == "" {
		s.timeId = tools.UUID()
		expireAt := tools.Now().Add(ReadyExpiration)
		s.room.AddTimer(s.timeId, expireAt, func(dt time.Duration) {
			s.SwitchTo(Deal)
		})
		s.room.Broadcast(&outer.NiuNiuStartCountDownNtf{ExpireAt: expireAt.UnixMilli()})
	}
}

func (s *StateReady) playerLeave(player *niuniuPlayer) {
	if s.playerCount() < s.gameParams().MinPlayPlayerCount && s.timeId != "" {
		s.room.CancelTimer(s.timeId)
		s.room.Broadcast(&outer.NiuNiuStopCountDownNtf{})
	}
}

func (s *StateReady) playerCount() int32 {
	var count int32

	for _, p := range s.niuniuPlayers {
		if p != nil {
			count++
		}
	}
	return count
}
