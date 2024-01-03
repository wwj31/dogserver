package niuniu

import (
	"time"

	"github.com/wwj31/dogactor/tools"

	"server/proto/outermsg/outer"
)

// 准备状态

type StateReady struct {
	*NiuNiu
	timeout string
}

func (s *StateReady) State() int {
	return Ready
}

func (s *StateReady) Enter() {
	s.onPlayerEnter = s.playerEnter
	s.onPlayerLeave = s.playerLeave
	s.timeout = ""

	s.RangePartInPlayer(func(seat int, player *niuniuPlayer) {
		if player.Gold <= s.baseScore() || player.trusteeshipCount >= s.gameParams().TrusteeshipCount {
			s.room.PlayerLeave(player.ShortId, true)
		}
	})
	s.Log().Infow("[NiuNiu] enter state ready ", "room", s.room.RoomId)
	s.room.Broadcast(&outer.NiuNiuReadyNtf{})

	s.checkAndSwitchNext()
}

func (s *StateReady) Leave() {
	s.room.CancelTimer(s.timeout)
	s.onPlayerEnter = nil
	s.onPlayerLeave = nil
	s.Log().Infow("[NiuNiu] leave state ready", "room", s.room.RoomId)
}

func (s *StateReady) Handle(shortId int64, v any) (result any) {
	return outer.ERROR_NIUNIU_STATE_MSG_INVALID
}

func (s *StateReady) checkAndSwitchNext() {
	if s.playerCount() >= s.gameParams().MinPlayPlayerCount && s.timeout == "" {
		s.timeout = tools.UUID()
		expireAt := tools.Now().Add(ReadyExpiration)
		s.room.AddTimer(s.timeout, expireAt, func(dt time.Duration) {
			s.SwitchTo(Deal)
		})
		s.room.Broadcast(&outer.NiuNiuStartCountDownNtf{ExpireAt: expireAt.UnixMilli()})
	}
}

func (s *StateReady) playerEnter(player *niuniuPlayer) {
	s.checkAndSwitchNext()
}

func (s *StateReady) playerLeave(player *niuniuPlayer) {
	if s.playerCount() < s.gameParams().MinPlayPlayerCount && s.timeout != "" {
		s.room.CancelTimer(s.timeout)
		s.room.Broadcast(&outer.NiuNiuStopCountDownNtf{})
		s.timeout = ""
	}
}
