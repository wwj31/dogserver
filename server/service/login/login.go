package login

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
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

	stored iface.SaveLoader

	accountMgr *account.AccountMgr
}

func New(s iface.SaveLoader) *Login {
	return &Login{
		stored: s,
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
		s.LoginReq(sourceId, gSession, msg)
	}
}

func (s *Login) LoginReq(sourceId string, gSession common.GSession, msg *message.LoginReq) {
	log.Debug(msg.String())

	acc, _ := s.accountMgr.Login(msg, s.stored)

	// 新、旧session相同，则同一连接多次登录，不做顶号处理
	if acc.GSession.Valid() && acc.GSession != gSession {
		oldId, _ := acc.GSession.Split()
		s.send2Gate(oldId, &inner.L2GTSessionDisabled{GateSession: acc.GSession.String()})
	}
	acc.GSession = gSession

	// 通知gate绑定角色服务器
	s.send2Gate(sourceId, &inner.L2GTSessionAssignGame{
		GateSession:  gSession.String(),
		GameServerId: acc.ServerId,
	})

	// 通知玩家登录成功
	s.send(gSession, &message.LoginRsp{
		UID: acc.UUId,
		RID: acc.LastRoleId,
	})
}

func (s *Login) send(gSession common.GSession, pb proto.Message) {
	if gSession.Invalid() {
		return
	}
	gateId, _ := gSession.Split()
	wrap := inner_message.NewGateWrapperByPb(pb, gSession)
	expect.Nil(s.Send(gateId, wrap))
}

func (s *Login) send2Gate(id common.ActorId, pb proto.Message) {
	if err := s.Send(id, pb); err != nil {
		log.KV("err", err).Error(" send to gateway error")
	}
}
