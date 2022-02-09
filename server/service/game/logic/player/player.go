package player

import (
	"reflect"
	"server/db/table"
	"server/proto/outermsg/outer"
	"time"

	"github.com/gogo/protobuf/proto"
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

func New(roleId uint64, gamer iface.Gamer, firstLogin bool) *Player {
	p := &Player{
		roleId:     roleId,
		gamer:      gamer,
		firstLogin: firstLogin,
	}
	return p
}

type (
	Player struct {
		actor.Base
		gamer iface.Gamer

		roleId     uint64
		firstLogin bool
		gSession   common.GSession // 网络session
		sender     common.SendTools

		playerData table.Player
		models     [all]iface.Modeler // 玩家所有功能模块

		saveTimerId string
		exitTimerId string
		liveTime    int64
	}
)

func (s *Player) OnInit() {
	s.sender = common.NewSendTools(s)
	s.playerData.RoleId = s.roleId
	// 不是首次登录，加载数据
	if !s.firstLogin {
		err := s.gamer.Load(&s.playerData)
		if err != nil {
			log.Errorw("load player data failed", "err", err)
			return
		}
	} else {
		defer s.store()
	}

	s.models[modRole] = role.New(models.New(s)) // 角色
	s.models[modItem] = item.New(models.New(s)) // 道具
	s.models[modMail] = mail.New(models.New(s)) // 邮件

}

func (s *Player) OnHandleMessage(sourceId, targetId string, msg interface{}) {
	name := controller.MsgName(msg)
	handle, ok := controller.MsgRouter[name]
	if !ok {
		log.Errorw("player undefined route ", "name", name)
		return
	}
	handle(s, msg)
	pt, ok := msg.(proto.Message)
	if ok {
		msgName := s.System().ProtoIndex().MsgName(pt)
		outer.Put(msgName, msg)
	}
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

	// 定时回存
	s.saveTimerId = s.AddTimer(tools.UUID(), tools.NowTime()+int64(1*time.Minute), func(dt int64) {
		s.store()
		s.live()
	}, -1)

	s.CancelTimer(s.exitTimerId)
}

func (s *Player) Logout() {
	for _, mod := range s.models {
		mod.OnLogout()
	}
}
func (s *Player) Online() bool {
	return s.Role().LoginAt() > s.Role().LogoutAt()
}

func (s *Player) OnStop() bool {
	return true
}

func (s *Player) IsNewRole() bool           { return s.firstLogin }
func (s *Player) Gamer() iface.Gamer        { return s.gamer }
func (s *Player) PlayerData() *table.Player { return &s.playerData }
func (s *Player) Role() iface.Role          { return s.models[modRole].(iface.Role) }
func (s *Player) Item() iface.Item          { return s.models[modItem].(iface.Item) }
func (s *Player) Mail() iface.Mailer        { return s.models[modMail].(iface.Mailer) }

// 回存数据
func (s *Player) store() {
	for _, mod := range s.models {
		mod.OnSave()
	}

	err := s.Gamer().Save(&s.playerData)
	if err != nil {
		log.Errorw("player store err", "err", err)
		return
	}

	log.Infow("player stored model", "RID", s.roleId)
}

func (s Player) live() {
	now := tools.NowTime()
	duration := now - s.liveTime
	if duration > int64(24*time.Hour) && !s.Online() {
		s.Exit()
	}
}
