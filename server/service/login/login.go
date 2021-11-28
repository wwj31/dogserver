package login

import (
	"fmt"
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
	common.Sender
	stored     iface.SaveLoader
	accountMgr *account.AccountMgr
}

func New(s iface.SaveLoader) *Login {
	return &Login{
		stored: s,
	}
}

func (s *Login) OnInit() {
	s.Sender = common.NewSendTools(s)
	s.accountMgr = account.NewAccountMgr()
	s.accountMgr.LoadAllAccount(s.stored)
	log.Debug("login OnInit")
}

func (s *Login) OnHandleMessage(sourceId, targetId string, msg interface{}) {
	v, gSession, err := inner_message.UnwrapperGateMsg(msg)
	expect.Nil(err)
	switch msg := v.(type) {
	case *message.LoginReq:
		err = s.LoginReq(sourceId, gSession, msg)
	default:
		err = fmt.Errorf("undefined msg type %v", msg)
	}

	if err != nil {
		log.KV("err", err).Error("handle message error")
	}
}

// 登录消息
func (s *Login) LoginReq(sourceId string, gSession common.GSession, msg *message.LoginReq) error {
	log.Debug(msg.String())

	acc, _ := s.accountMgr.Login(msg, s.stored)

	// 新、旧session相同，则同一连接多次登录，不做顶号处理
	if acc.GSession().Valid() && acc.GSession() != gSession {
		oldId, _ := acc.GSession().Split()
		_ = s.Send2Gate(oldId, &inner.L2GTSessionDisabled{GateSession: acc.GSession().String()})
	}
	acc.SetgSession(gSession)

	// 通知gate绑定角色服务器
	err := s.Send2Gate(sourceId, &inner.L2GTSessionAssignGame{
		GateSession:  gSession.String(),
		GameServerId: acc.ServerId(),
	})
	if err != nil {
		return err
	}

	// 通知玩家登录成功
	return s.Send2Client(gSession, &message.LoginRsp{
		UID: acc.UUId,
		RID: acc.LastRoleId,
	})
}
