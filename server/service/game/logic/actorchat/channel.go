package actorchat

import (
	"server/common"

	"github.com/gogo/protobuf/proto"
)

func NewChannel() *Channel {
	return &Channel{
		id2Session: make(map[common.ActorId]common.GSession, 10),
	}
}

type Channel struct {
	id2Session map[common.ActorId]common.GSession
}

func (s *Channel) Join(playerId common.ActorId, gs common.GSession) {
	s.id2Session[playerId] = gs
}

func (s *Channel) Leave(playerId common.ActorId) {
	delete(s.id2Session, playerId)
}

func (s *Channel) Broadcast(msg proto.Message) {
	for _,
	delete(s.id2Session, playerId)
}
