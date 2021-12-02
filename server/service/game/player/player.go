package player

import (
	"server/common"
	"server/service/game/iface"
	"server/service/game/player/model"
	"server/service/game/player/role"
)

type (
	Player struct {
		game        iface.Gamer
		gateSession common.GSession
		models      [All]iface.Modeler
	}
)

func New(roleId uint64, game iface.Gamer) *Player {
	p := &Player{game: game}
	p.models[Role] = role.New(roleId, model.New(p)) // 角色

	return p
}

func (s *Player) Game() iface.Gamer                       { return s.game }
func (s *Player) GateSession() common.GSession            { return s.gateSession }
func (s *Player) SetGateSession(gSession common.GSession) { s.gateSession = gSession }

func (s *Player) Login() {
	for _, model := range s.models {
		model.OnLogin()
	}
}

func (s *Player) Logout() {
	for _, model := range s.models {
		model.OnLogout()
	}
}

func (s *Player) Role() iface.Role {
	return s.models[Role].(iface.Role)
}
