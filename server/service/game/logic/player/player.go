package player

import (
	"server/db/table"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player/models/chat"
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

// 引用DDD 实体、聚合概念
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
		models     [allmod]iface.Modeler // 玩家所有功能模块

		saveTimerId string
		exitTimerId string
		keepAlive   int64
	}
)

func (s *Player) OnInit() {
	s.sender = common.NewSendTools(s)
	s.playerData.RoleId = s.roleId
	data := table.Player{RoleId: s.roleId}

	// load data if not first login
	if !s.firstLogin {
		err := s.gamer.Load(&data)
		if err != nil {
			log.Errorw("load player data failed", "err", err)
			return
		}
	} else {
		defer s.store(true)
	}

	s.models[modRole] = role.New(models.New(s), data.RoleBytes) // 角色
	s.models[modItem] = item.New(models.New(s), data.ItemBytes) // 道具
	s.models[modMail] = mail.New(models.New(s), data.MailBytes) // 邮件
	s.models[modChat] = chat.New(models.New(s))                 // 聊天

	// 定时回存
	randTime := tools.Now().Add(time.Second)
	s.saveTimerId = s.AddTimer(tools.XUID(), randTime, func(dt time.Duration) {
		s.store()
		s.checkAlive()
	}, -1)
	log.Infow("player actor activated", "id", s.ID(), "firstLogin", s.firstLogin)
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
	s.keepAlive = tools.NowTime()
}

func (s *Player) Send2Client(pb proto.Message) {
	if pb == nil || !s.Online() {
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
}

func (s *Player) GateSession() common.GSession            { return s.gSession }
func (s *Player) SetGateSession(gSession common.GSession) { s.gSession = gSession }
func (s *Player) Online() bool                            { return s.Role().LoginAt() > s.Role().LogoutAt() }
func (s *Player) IsNewRole() bool                         { return s.firstLogin }
func (s *Player) Gamer() iface.Gamer                      { return s.gamer }
func (s *Player) Role() iface.Role                        { return s.models[modRole].(iface.Role) }
func (s *Player) Item() iface.Item                        { return s.models[modItem].(iface.Item) }
func (s *Player) Mail() iface.Mailer                      { return s.models[modMail].(iface.Mailer) }
func (s *Player) Chat() iface.Chat                        { return s.models[modChat].(iface.Chat) }

// 回存数据
func (s *Player) store(new ...bool) {
	data := table.Player{RoleId: s.roleId}
	for _, mod := range s.models {
		mod.OnSave(&data)
	}

	var insert bool
	if len(new) > 0 && new[0] {
		insert = true
	}

	err := s.Gamer().Store(insert, &data)
	if err != nil {
		log.Errorw("player store err", "err", err)
		return
	}

	log.Infow("player stored model", "RID", s.roleId)
}

const aliveDuration = 5 * time.Second // 24*time.Hour
func (s Player) checkAlive() {
	now := tools.NowTime()
	duration := now - s.keepAlive
	if duration > int64(aliveDuration) && !s.Online() {
		s.Exit()
	}
}
