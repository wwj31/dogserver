package niuniu

import (
	"server/proto/outermsg/outer"
)

// 压注状态

type StateBetting struct {
	*NiuNiu
}

func (s *StateBetting) State() int {
	return Betting
}

func (s *StateBetting) Enter() {
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

	return outer.ERROR_NIUNIU_STATE_MSG_INVALID
}
