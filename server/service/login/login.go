package login

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/iniconfig"
	"github.com/wwj31/dogactor/log"
	"server/common"
	"server/proto/inner_message"
	"server/proto/inner_message/inner"
	"server/proto/message"
	"server/service/login/account"
	"server/service/login/iface"
)

type Login struct {
	actor.Base
	Config iniconfig.Config

	stored iface.SaveLoader

	accountMgr *account.AccountMgr
}

func New(s iface.SaveLoader, conf iniconfig.Config) *Login {
	return &Login{
		stored: s,
		Config: conf,
	}
}

func (s *Login) OnInit() {
	s.accountMgr = account.NewAccountMgr()
	s.accountMgr.LoadAllAccount(s.stored)
	log.Debug("login OnInit")
}

func (s *Login) OnHandleMessage(sourceId, targetId string, msg interface{}) {
	v, gSession, err := inner_message.UnwrapperGateMsg(msg)
	expect.Nil(err)

	switch msg := v.(type) {
	case *message.LoginReq:
		s.LoginReq(gSession, msg)
	}
}

func (s *Login) LoginReq(gSession string, msg *message.LoginReq) {
	log.Debug(msg.String())

	acc, _ := s.accountMgr.Login(msg, s.stored)

	// 通知顶号
	if acc.GSession != "" {
		s.send(gSession, &inner.L2GTSessionDisabled{GateSession: gSession})
	}
	acc.GSession = gSession

	// 通知gate绑定角色服务器
	s.send(gSession, &inner.L2GTSessionAssignGame{
		GateSession:  gSession,
		GameServerId: acc.ServerId,
	})

	// 通知玩家登录成功
	s.send(gSession, &message.LoginRsp{
		UID: acc.UUId,
		RID: acc.LastRoleId,
	})
}

func (s *Login) send(gSession string, pb proto.Message) {
	if gSession == "" {
		return
	}
	gateId, _ := common.SplitGateSession(gSession)
	wrap := inner_message.NewGateWrapperByPb(pb, gSession)
	expect.Nil(s.Send(gateId, wrap))
}
