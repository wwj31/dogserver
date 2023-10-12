package niuniu

import (
	"server/proto/outermsg/outer"
)

// 抢庄状态

type StateMaster struct {
	*NiuNiu
}

func (s *StateMaster) State() int {
	return DecideMaster
}

func (s *StateMaster) Enter() {
	s.Log().Infow("[NiuNiu] enter state master ", "room", s.room.RoomId)
}

func (s *StateMaster) Leave() {
	s.Log().Infow("[NiuNiu] leave state master", "room", s.room.RoomId)
}

func (s *StateMaster) Handle(shortId int64, v any) (result any) {
	player, _ := s.findNiuNiuPlayer(shortId)
	if player == nil {
		s.Log().Warnw("player not in room", "roomId", s.room.RoomId, "shortId", shortId)
		return outer.ERROR_PLAYER_NOT_IN_ROOM
	}

	return outer.ERROR_NIUNIU_STATE_MSG_INVALID
}
