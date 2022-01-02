package login

import (
	"fmt"
	"server/common"
	"server/common/log"
	"server/proto/inner_message/inner"
	"server/proto/message"
	"server/service/game/iface"
	"server/service/login/account"

	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
)

type Login struct {
	actor.Base
	common.SendTools
	storer     iface.SaveLoader
	accountMgr account.AccountMgr
}

func New(s iface.SaveLoader) *Login {
	return &Login{
		storer: s,
	}
}

func (s *Login) OnInit() {
	s.SendTools = common.NewSendTools(s)
	s.accountMgr = account.NewAccountMgr()
	s.accountMgr.LoadAllAccount(s.storer)
	log.Debugf("login OnInit")
}

func (s *Login) OnHandleMessage(sourceId, targetId string, msg interface{}) {
	v, gSession, err := common.UnwrapperGateMsg(msg)
	expect.Nil(err)
	switch msg := v.(type) {
	case *message.LoginReq:
		err = s.LoginReq(sourceId, gSession, msg)
	default:
		err = fmt.Errorf("undefined localmsg type %v", msg)
	}

	if err != nil {
		log.Errorw("handle message error", "err", err)
	}
}

func (s *Login) LoginReq(sourceId string, gSession common.GSession, msg *message.LoginReq) error {
	log.Debugf(msg.String())

	if common.LoginChecksum(msg) != msg.Checksum {
		return fmt.Errorf("login req checksum failed msg:%v", msg.String())
	}

	acc, _ := s.accountMgr.Login(msg, s.storer)

	// 新、旧sessio不相同做顶号处理
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
		UID: acc.UUId(),
		RID: acc.LastRoleId(),
	})
}
