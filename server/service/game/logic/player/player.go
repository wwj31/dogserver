package player

import (
	"github.com/wwj31/dogactor/log"
	"github.com/wwj31/dogactor/tools"
	"reflect"
	"server/common"
	"server/db/table"
	"server/service/game/iface"
	"server/service/game/logic/model"
	"server/service/game/logic/player/item"
	"server/service/game/logic/player/role"
	"time"
)

type (
	Player struct {
		game        iface.Gamer
		gateSession common.GSession
		saveTimerId string

		models [all]iface.Modeler
	}
)

func New(roleId uint64, game iface.Gamer) *Player {
	p := &Player{game: game}
	p.models[modRole] = role.New(roleId, model.New(p)) // 角色
	p.models[modItem] = item.New(roleId, model.New(p)) // 道具

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
		s.Game().CancelTimer(s.saveTimerId)
	}

	s.saveTimerId = s.Game().AddTimer(tools.UUID(), 1*time.Minute, func(dt int64) {
		s.store()
	}, -1)
}

func (s *Player) Logout() {
	for _, mod := range s.models {
		mod.OnLogout()
	}
	s.store()
	s.Game().CancelTimer(s.saveTimerId)
}

// 停服回存全量数据
func (s *Player) Stop() {
	for _, mod := range s.models {
		mod.OnStop()
	}
	s.store()
}
func (s *Player) IsNewRole() bool {
	return s.Role().LoginAt() == 0
}

func (s *Player) Role() iface.Role {
	return s.models[modRole].(iface.Role)
}
func (s *Player) Item() iface.Item {
	return s.models[modItem].(iface.Item)
}

func (s *Player) SetTable(table.Tabler) {

}

// 回存功能模块
func (s *Player) store() {
	logFiled := log.Fields{"roleId": s.Role().RoleId()}
	for _, mod := range s.models {
		tab := mod.Table()
		if tab == nil || reflect.ValueOf(tab).IsNil() {
			continue
		}
		err := s.Game().Save(tab)
		mod.SetTable(nil)
		if err != nil {
			log.KV("err", err).Error("player store err")
		} else {
			logFiled["table"] = tab.TableName()
		}
	}
	log.KVs(logFiled).Info("player stored model")
}
