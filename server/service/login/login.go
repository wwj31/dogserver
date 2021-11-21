package login

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/iniconfig"
	"github.com/wwj31/dogactor/log"
	"server/common"
	"server/proto/inner_message"
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
	s.accountMgr = account.NewAccountMgr(s.stored)
	s.accountMgr.LoadAllAccount()
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

	acc, _ := s.accountMgr.Login(msg)
	s.send(gSession, &message.LoginRsp{
		UID:      int64(acc.UUId),
		RID:      int64(acc.LastRoleId),
		ServerId: acc.Roles[acc.LastRoleId].SId,
	})
}

func (s *Login) send(gSession string, pb proto.Message) {
	gateId, _ := common.SplitGateSession(gSession)
	wrap := inner_message.NewGateWrapperByPb(pb, gSession)
	expect.Nil(s.Send(gateId, wrap))
}
