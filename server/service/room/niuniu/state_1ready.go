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
	s.timeId = ""

	s.RangePartInPlayer(func(seat int, player *niuniuPlayer) {
		if player.Gold <= s.baseScore() {
			s.room.PlayerLeave(player.ShortId, true)
			return
		}
	})
	s.Log().Infow("[NiuNiu] enter state ready ", "room", s.room.RoomId)
	s.room.Broadcast(&outer.NiuNiuReadyNtf{})

	s.checkAndSwitchNext()
}

func (s *StateReady) Leave() {
	s.onPlayerEnter = nil
	s.onPlayerLeave = nil
	s.Log().Infow("[NiuNiu] leave state ready", "room", s.room.RoomId)
}

func (s *StateReady) Handle(shortId int64, v any) (result any) {
	return outer.ERROR_NIUNIU_STATE_MSG_INVALID
}

func (s *StateReady) checkAndSwitchNext() {
	if s.playerCount() >= s.gameParams().MinPlayPlayerCount && s.timeId == "" {
		s.timeId = tools.UUID()
		expireAt := tools.Now().Add(ReadyExpiration)
		s.room.AddTimer(s.timeId, expireAt, func(dt time.Duration) {
			s.SwitchTo(Deal)
		})
		s.room.Broadcast(&outer.NiuNiuStartCountDownNtf{ExpireAt: expireAt.UnixMilli()})
	}
}

func (s *StateReady) playerEnter(player *niuniuPlayer) {
	s.checkAndSwitchNext()
}

func (s *StateReady) playerLeave(player *niuniuPlayer) {
	if s.playerCount() < s.gameParams().MinPlayPlayerCount && s.timeId != "" {
		s.room.CancelTimer(s.timeId)
		s.room.Broadcast(&outer.NiuNiuStopCountDownNtf{})
		s.timeId = ""
	}
}
