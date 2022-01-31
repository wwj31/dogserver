package player

import (
	"reflect"
	"server/db/table"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/tools"

	"server/common"
	"server/common/log"
	"server/service/game/iface"
	"server/service/game/logic/player/controller"
	"server/service/game/logic/player/models"
	"server/service/game/logic/player/models/item"
	"server/service/game/logic/player/models/mail"
	"server/service/game/logic/player/models/role"
)

// model作为功能聚合，player作为聚合根，roleId为聚合根ID
// 聚合之间通过聚合根关联引用，聚合之间相互访问需先访问聚合根，在导航到相关功能
// 解决玩家复杂的功能模块相互引用带来的混乱问题，功能模块化，模块间解耦

func New(roleId uint64, gamer iface.Gamer) *Player {
	p := &Player{
		roleId: roleId,
		gamer:  gamer,
	}
	return p
}

type (
	Player struct {
		actor.Base
		gamer iface.Gamer

		roleId   uint64
		gSession common.GSession // 网络session
		sender   common.SendTools

		models [all]iface.Modeler // 玩家所有功能模块

		saveTimerId string
		exitTimerId string
		liveTime    int64
	}
)

func (s *Player) OnInit() {
	s.sender = common.NewSendTools(s)

	s.models[modRole] = role.New(s.roleId, models.New(s)) // 角色
	s.models[modItem] = item.New(s.roleId, models.New(s)) // 道具
	s.models[modMail] = mail.New(s.roleId, models.New(s)) // 邮件

	s.saveTimerId = s.AddTimer(tools.UUID(), tools.NowTime()+int64(1*time.Minute), func(dt int64) {
		s.store()
		s.live()
	}, -1)
}

func (s *Player) OnHandleMessage(sourceId, targetId string, msg interface{}) {
	name := controller.MsgName(msg)
	handle, ok := controller.MsgRouter[name]
	if !ok {
		log.Errorw("player undefined route ", "name", name)
		return
	}
	handle(s, msg)
	s.liveTime = tools.NowTime()
}

func (s *Player) GateSession() common.GSession            { return s.gSession }
func (s *Player) SetGateSession(gSession common.GSession) { s.gSession = gSession }
func (s *Player) Send2Client(pb proto.Message) {
	if pb == nil || reflect.ValueOf(pb).IsNil() {
		return
	}
	if err := s.sender.Send2Client(s.gSession, pb); err != nil {
		log.Errorw("player send faild", "err", err)
	}
}

func (s *Player) Login() {
	for _, mod := range s.models {
		mod.OnLogin()
	}
	s.CancelTimer(s.exitTimerId)
}

func (s *Player) Logout() {
	for _, mod := range s.models {
		mod.OnLogout()
	}

	exitAt := tools.NowTime() + 3*time.Minute.Nanoseconds()
	s.exitTimerId = s.AddTimer(tools.UUID(), exitAt, func(dt int64) {
		s.store()
		s.Exit()
	})
}

func (s *Player) OnStop() bool {
	s.Logout()
	return true
}

func (s *Player) IsNewRole() bool    { return s.Role().IsNewRole() }
func (s *Player) Gamer() iface.Gamer { return s.gamer }
func (s *Player) Role() iface.Role   { return s.models[modRole].(iface.Role) }
func (s *Player) Item() iface.Item   { return s.models[modItem].(iface.Item) }
func (s *Player) Mail() iface.Mailer { return s.models[modMail].(iface.Mailer) }

// 回存数据
func (s *Player) store() {
	logFiled := []interface{}{"roleId", s.Role().RoleId()}
	var tabs []table.Tabler
	for _, mod := range s.models {
		tab := mod.Table()
		if tab == nil {
			continue
		}
		tabs = append(tabs, tab)
		mod.SetTable(nil)
		logFiled = append(logFiled, "table", tab.TableName())
	}

	err := s.Gamer().Save(tabs...)
	if err != nil {
		log.Errorw("player store err", "err", err)
	}

	log.Infow("player stored model", logFiled...)
}

func (s Player) live() {
	now := tools.NowTime()
	duration := now - s.liveTime
	if duration > int64(24*time.Hour) {
		s.Exit()
	}
}
