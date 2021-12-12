package player

import (
	"server/common"
	"server/service/game/iface"
)

type Manager struct {
	game            iface.Gamer
	playerbySession map[common.GSession]iface.Player
	playerbyUID     map[uint64]iface.Player
	playerbyRID     map[uint64]iface.Player
}

func NewMgr(g iface.Gamer) *Manager {
	return &Manager{
		game:            g,
		playerbySession: make(map[common.GSession]iface.Player, 100),
		playerbyUID:     make(map[uint64]iface.Player, 100),
		playerbyRID:     make(map[uint64]iface.Player, 100),
	}
}

func (s *Manager) PlayerBySession(gateSession common.GSession) (iface.Player, bool) {
	p, ok := s.playerbySession[gateSession]
	return p, ok
}
func (s *Manager) SetPlayer(p iface.Player) {
	s.playerbySession[p.GateSession()] = p
	s.playerbyUID[p.Role().UUId()] = p
	s.playerbyRID[p.Role().RoleId()] = p
}

func (s *Manager) PlayerByUID(uid uint64) (iface.Player, bool) {
	p, ok := s.playerbyUID[uid]
	return p, ok
}

func (s *Manager) PlayerByRID(rid uint64) (iface.Player, bool) {
	p, ok := s.playerbyRID[rid]
	return p, ok
}

func (s *Manager) RangeOnline(f func(player iface.Player), except ...uint64) {
	e := map[uint64]struct{}{}
	for _, id := range except {
		e[id] = struct{}{}
	}
	for _, p := range s.playerbySession {
		if _, exist := e[p.Role().UUId()]; exist {
			continue
		}
		f(p)
	}
}

func (s *Manager) Stop() {
	for _, p := range s.playerbyUID {
		p.Stop()
	}
}
