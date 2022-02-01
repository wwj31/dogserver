package player

import (
	"server/common"
	"server/common/log"
	"server/service/game/iface"

	"github.com/gogo/protobuf/proto"
)

type Manager struct {
	game             iface.Gamer
	playerByGSession map[common.GSession]common.ActorId
	gSessionByPlayer map[common.ActorId]common.GSession
}

func NewMgr(g iface.Gamer) *Manager {
	return &Manager{
		game:             g,
		playerByGSession: make(map[common.GSession]common.ActorId, 100),
		gSessionByPlayer: make(map[common.ActorId]common.GSession, 100),
	}
}

func (s *Manager) AssociateSession(id common.ActorId, gSession common.GSession) {
	s.playerByGSession[gSession] = id
	s.gSessionByPlayer[id] = gSession
}

func (s *Manager) PlayerBySession(gateSession common.GSession) (common.ActorId, bool) {
	p, ok := s.playerByGSession[gateSession]
	return p, ok
}
func (s *Manager) GSessionByPlayer(id common.ActorId) (common.GSession, bool) {
	p, ok := s.gSessionByPlayer[id]
	return p, ok
}
func (s *Manager) DelGSession(gateSession common.GSession) {
	id, ok := s.playerByGSession[gateSession]
	if ok {
		delete(s.gSessionByPlayer, id)
	}
	delete(s.playerByGSession, gateSession)

}

func (s *Manager) RangeOnline(f func(common.GSession, common.ActorId)) {
	for id, s := range s.gSessionByPlayer {
		f(s, id)
	}
}

func (s *Manager) Broadcast(msg proto.Message) {
	sender := common.NewSendTools(s.game)
	s.RangeOnline(func(gs common.GSession, id common.ActorId) {
		if err := sender.Send2Client(gs, msg); err != nil {
			log.Errorw("broadcast localmsg error", "err", err, "gSession", gs, "Player", id)
		}
	})
}
