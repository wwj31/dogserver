package player

import (
	"server/common"
	"server/service/game/iface"
)

type Player struct {
	game        iface.Gamer
	gateSession common.GSession

	iface.Role
}

func (s *Player) GateSession() common.GSession            { return s.gateSession }
func (s *Player) SetGateSession(gSession common.GSession) { s.gateSession = gSession }
