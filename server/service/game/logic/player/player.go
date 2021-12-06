package player

import (
	"github.com/wwj31/dogactor/log"
	"github.com/wwj31/dogactor/tools"
	"server/common"
	"server/service/game/iface"
	"server/service/game/logic/model"
	"server/service/game/logic/role"
	"time"
)

type (
	Player struct {
		game        iface.Gamer
		gateSession common.GSession
		models      [all]iface.Modeler

		saveTimerId string
	}
)

func New(roleId uint64, game iface.Gamer) *Player {
	p := &Player{game: game}
	p.models[modRole] = role.New(roleId, model.New(p)) // 角色

	return p
}

func (s *Player) Game() iface.Gamer                       { return s.game }
func (s *Player) GateSession() common.GSession            { return s.gateSession }
func (s *Player) SetGateSession(gSession common.GSession) { s.gateSession = gSession }

func (s *Player) Login() {
	for _, mod := range s.models {
		mod.OnLogin()
	}

	if s.saveTimerId != "" {
		s.game.CancelTimer(s.saveTimerId)
	}

	s.saveTimerId = s.game.AddTimer(tools.UUID(), 1*time.Second, func(dt int64) {
		log.KV("roleId", s.Role().RoleId()).Info("player save")
		s.store()
	}, -1)
}

func (s *Player) Logout() {
	for _, mod := range s.models {
		mod.OnLogout()
	}
	s.store()
	s.game.CancelTimer(s.saveTimerId)
}

func (s *Player) Role() iface.Role {
	return s.models[modRole].(iface.Role)
}
func (s *Player) store() {
	for _, mod := range s.models {
		err := s.game.Save(mod.Table())
		if err != nil {
			log.KV("err", err).Error("player store err")
		}
	}
}
