package mahjong

import (
	"server/common/log"
	"server/proto/outermsg/outer"
	"server/service/room"
)

type StateDeal struct {
	*mahjong
}

func (s StateDeal) State() int {
	return Ready
}

func (s StateDeal) Enter(fsm *room.FSM) {
	log.Infow("mahjong enter deal ", "room", s.room.RoomId)
}

func (s StateDeal) Leave(fsm *room.FSM) {
	log.Infow("mahjong leave ready", "room", s.room.RoomId)
}

func (s StateDeal) Handle(fsm *room.FSM, i ...any) (result any) {
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
