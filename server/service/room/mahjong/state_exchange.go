package mahjong

import (
	"server/common/log"
	"server/proto/outermsg/outer"
	"server/service/room"
)

// 换三张

type StateExchange3 struct {
	*Mahjong
}

func (s *StateExchange3) State() int {
	return Exchange3
}

func (s *StateExchange3) Enter(fsm *room.FSM) {
	s.room.Broadcast(&outer.MahjongBTEExchange3Ntf{})
	log.Infow("Mahjong enter exchange3", "room", s.room.RoomId)
}

func (s *StateExchange3) Leave(fsm *room.FSM) {
	log.Infow("Mahjong leave exchange3", "room", s.room.RoomId)
}

func (s *StateExchange3) Handle(fsm *room.FSM, v any, shortId int64) (result any) {
	return nil
}
