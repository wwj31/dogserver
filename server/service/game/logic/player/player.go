package player

import (
	"server/proto/outermsg/outer"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/tools"

	"server/common"
	"server/common/log"
	"server/service/game/iface"
	"server/service/game/logic/player/controller"
)

func New(roleId string, gamer iface.Gamer, firstLogin bool) *Player {
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
		gamer    iface.Gamer
		gSession common.GSession // 网络session
		sender   common.SendTools
		models   [allmod]iface.Modeler // 玩家所有功能模块

		roleId      string
		firstLogin  bool
		saveTimerId string
		exitTimerId string
		keepAlive   int64
	}
)

func (s *Player) OnInit() {
	s.sender = common.NewSendTools(s)
	if s.firstLogin {
		defer s.store()
	}

	// 初始化玩家所有功能模块
	s.initModule()

	// 定时回存
	randTime := tools.Now().Add(time.Second)
	s.saveTimerId = s.AddTimer(tools.XUID(), randTime, func(dt time.Duration) {
		s.store()
		s.checkAlive()
	}, -1)
	log.Infow("player actor activated", "id", s.ID(), "firstLogin", s.firstLogin)
}

func (s *Player) OnHandle(msg actor.Message) {
	rawMsg := msg.RawMsg()
	name := controller.MsgName(rawMsg)
	handle, ok := controller.MsgRouter[name]
	if !ok {
		log.Errorw("player undefined route ", "name", name)
		return
	}

	handle(s, rawMsg)

	pt, ok := rawMsg.(proto.Message)
	if ok {
		msgName := s.System().ProtoIndex().MsgName(pt)
		outer.Put(msgName, rawMsg)
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

// 回存数据
func (s *Player) store() {
	for _, mod := range s.models {
		mod.OnSave()
	}
	log.Infow("player stored model", "RID", s.roleId)
}

const aliveDuration = 5 * time.Second // 24*time.Hour
func (s Player) checkAlive() {
	now := tools.NowTime()
	duration := now - s.keepAlive
	if duration > int64(aliveDuration) && !s.Online() {
		//s.Exit()
	}
}
