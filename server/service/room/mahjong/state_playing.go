package mahjong

import (
	"server/common/log"
	"server/service/room"
)

// 游戏状态

type StatePlaying struct {
	*Mahjong
}

func (s *StatePlaying) State() int {
	return Playing
}

func (s *StatePlaying) Enter(fsm *room.FSM) {
	log.Infow("Mahjong enter playing", "room", s.room.RoomId)
}

func (s *StatePlaying) Leave(fsm *room.FSM) {
	log.Infow("Mahjong leave playing", "room", s.room.RoomId)
}

func (s *StatePlaying) Handle(fsm *room.FSM, v any, shortId int64) (result any) {
	return nil
}
