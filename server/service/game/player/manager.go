package player

import (
	"server/common"
	"server/service/game/iface"
)

type Manager struct {
	game            iface.Gamer
	playerbySession map[common.GSession]*Player
	playerbyUID     map[uint64]*Player
	playerbyRID     map[uint64]*Player
}

func NewMgr(g iface.Gamer) *Manager {
	return &Manager{
		game:            g,
		playerbySession: make(map[common.GSession]*Player, 100),
		playerbyUID:     make(map[uint64]*Player, 100),
		playerbyRID:     make(map[uint64]*Player, 100),
	}
}

func (s *Manager) PlayerBySession(gateSession common.GSession) (iface.Player, bool) {
	p, ok := s.playerbySession[gateSession]
	return p, ok
}

func (s *Manager) PlayerByUID(uid uint64) (iface.Player, bool) {
	p, ok := s.playerbyUID[uid]
	return p, ok
}

func (s *Manager) PlayerByRID(rid uint64) (iface.Player, bool) {
	p, ok := s.playerbyRID[rid]
	return p, ok
}
