package mahjong

import (
	"server/common/log"
	"server/proto/outermsg/outer"
	"server/service/room"
)

type StateReady struct {
	*Mahjong
}

func (s StateReady) State() int {
	return Ready
}

func (s StateReady) Enter(fsm *room.FSM) {
	log.Infow("Mahjong enter ready ", "room", s.room.RoomId)
}

func (s StateReady) Leave(fsm *room.FSM) {
	log.Infow("Mahjong leave ready", "room", s.room.RoomId)
}

func (s StateReady) Handle(fsm *room.FSM, i ...any) (result any) {
	if len(i) <= 0 {
		log.Errorw("i len = 0", "room", s.room.RoomId)
		return
	}

	v := i[0]
	switch msg := v.(type) {
	case *outer.JoinRoomReq:
		_ = msg
	}
	return nil
}
