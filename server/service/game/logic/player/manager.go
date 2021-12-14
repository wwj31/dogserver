package player

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/log"
	"server/common"
	"server/service/game/iface"
)

type Manager struct {
	game             iface.Gamer
	playerByGsession map[common.GSession]common.ActorId
	gSessionByPlayer map[common.ActorId]common.GSession
}

func NewMgr(g iface.Gamer) *Manager {
	return &Manager{
		game:             g,
		playerByGsession: make(map[common.GSession]common.ActorId, 100),
		gSessionByPlayer: make(map[common.ActorId]common.GSession, 100),
	}
}

func (s *Manager) SetPlayer(gSession common.GSession, id common.ActorId) {
	s.playerByGsession[gSession] = id
	s.gSessionByPlayer[id] = gSession
}

func (s *Manager) PlayerBySession(gateSession common.GSession) (common.ActorId, bool) {
	p, ok := s.playerByGsession[gateSession]
	return p, ok
}
func (s *Manager) GSessionByPlayer(id common.ActorId) (common.GSession, bool) {
	p, ok := s.gSessionByPlayer[id]
	return p, ok
}
func (s *Manager) DelGSession(gateSession common.GSession) {
	id, ok := s.playerByGsession[gateSession]
	if ok {
		delete(s.gSessionByPlayer, id)
	}
	delete(s.playerByGsession, gateSession)

}

func (s *Manager) RangeOnline(f func(common.GSession, common.ActorId)) {
	for id, s := range s.gSessionByPlayer {
		f(s, id)
	}
}

func (s *Manager) Broadcast(msg proto.Message) {
	s.RangeOnline(func(gs common.GSession, id common.ActorId) {
		if err := s.game.Send2Client(gs, msg); err != nil {
			log.KVs(log.Fields{"err": err, "gSession": gs, "Player": id}).
				ErrorStack(2, "broadcast msg error")
		}
	})
}
