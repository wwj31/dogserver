package player

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/log"
	"github.com/wwj31/dogactor/tools"
	"reflect"
	"server/common"
	"server/common/toml"
	"server/db"
	"server/service/game/iface"
	"server/service/game/logic/model"
	"server/service/game/logic/player/controller"
	"server/service/game/logic/player/item"
	"server/service/game/logic/player/role"
	"time"
)

func New(roleId uint64) *Player {
	p := &Player{roleId: roleId}
	return p
}

type (
	Player struct {
		actor.Base
		iface.SaveLoader

		gSession common.GSession
		sender   common.SendTools

		roleId      uint64
		saveTimerId string

		models [all]iface.Modeler
	}
)

func (s *Player) OnInit() {
	s.SaveLoader = db.New(toml.Get("mysql"), toml.Get("database"))
	s.sender = common.NewSendTools(s)

	s.models[modRole] = role.New(s.roleId, model.New(s)) // 角色
	s.models[modItem] = item.New(s.roleId, model.New(s)) // 道具

	s.saveTimerId = s.AddTimer(tools.UUID(), 1*time.Minute, func(dt int64) {
		s.store()
	}, -1)
}
func (s *Player) OnHandleMessage(sourceId, targetId string, msg interface{}) {
	name := controller.MsgName(msg)
	handle, ok := controller.MsgRouter[name]
	if !ok {
		log.KV("name", name).Error("player undefined route ")
		return
	}
	handle(s, msg)
}

func (s *Player) GateSession() common.GSession            { return s.gSession }
func (s *Player) SetGateSession(gSession common.GSession) { s.gSession = gSession }
func (s *Player) Send2Client(pb proto.Message) {
	if err := s.sender.Send2Client(s.gSession, pb); err != nil {
		log.KV("err", err).Error("player send faild")
	}
}

func (s *Player) Login() {
	for _, mod := range s.models {
		mod.OnLogin()
	}

	if s.saveTimerId != "" {
		s.CancelTimer(s.saveTimerId)
	}
}

func (s *Player) Logout() {
	for _, mod := range s.models {
		mod.OnLogout()
	}
	s.store()
}

func (s *Player) OnStop() bool {
	s.Logout()
	return true
}

func (s *Player) IsNewRole() bool  { return s.Role().LoginAt() == 0 }
func (s *Player) Role() iface.Role { return s.models[modRole].(iface.Role) }
func (s *Player) Item() iface.Item { return s.models[modItem].(iface.Item) }

// 回存功能模块
func (s *Player) store() {
	logFiled := log.Fields{"roleId": s.Role().RoleId()}
	for _, mod := range s.models {
		tab := mod.Table()
		if tab == nil || reflect.ValueOf(tab).IsNil() {
			continue
		}
		err := s.Save(tab)
		mod.SetTable(nil)
		if err != nil {
			log.KV("err", err).Error("player store err")
		} else {
			logFiled["table"] = tab.TableName()
		}
	}
	log.KVs(logFiled).Info("player stored model")
}
