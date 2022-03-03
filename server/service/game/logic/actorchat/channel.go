package actorchat

import (
	"server/common"

	"github.com/wwj31/dogactor/expect"

	"github.com/gogo/protobuf/proto"
)

func NewChannel(sender common.Sender) *Channel {
	return &Channel{
		Sender:     sender,
		id2Session: make(map[common.ActorId]common.GSession, 10),
	}
}

type Channel struct {
	common.Sender
	id2Session map[common.ActorId]common.GSession
}

func (s *Channel) Join(playerId common.ActorId, gs common.GSession) {
	s.id2Session[playerId] = gs
}

func (s *Channel) Leave(playerId common.ActorId) {
	delete(s.id2Session, playerId)
}

func (s *Channel) Broadcast(msg proto.Message) {
	for _, gSession := range s.id2Session {
		err := s.Send2Client(gSession, msg)
		expect.Nil(err)
	}
}
