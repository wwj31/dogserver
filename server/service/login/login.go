package login

import (
	"fmt"
	"server/common"
	"server/common/log"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/iface"
	"server/service/login/account"

	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
)

type Login struct {
	actor.Base
	common.SendTools
	storage    iface.SaveLoader
	accountMgr account.AccountMgr
}

func New(s iface.SaveLoader) *Login {
	return &Login{
		storage: s,
	}
}

func (s *Login) OnInit() {
	s.SendTools = common.NewSendTools(s)
	s.accountMgr = account.NewAccountMgr()
	s.accountMgr.LoadAllAccount(s.storage)
	log.Debugf("login OnInit")
}

func (s *Login) OnHandleMessage(sourceId, targetId string, msg interface{}) {
	v, _, gSession, err := common.UnwrapperGateMsg(msg)

	expect.Nil(err)
	switch msg := v.(type) {
	case *outer.LoginReq:
		err = s.LoginReq(sourceId, gSession, msg)
	default:
		err = fmt.Errorf("undefined localmsg type %v", msg)
	}

	if err != nil {
		log.Errorw("handle outer error", "err", err)
	}
}

func (s *Login) LoginReq(sourceId string, gSession common.GSession, msg *outer.LoginReq) error {
	log.Debugf(msg.String())

	if common.LoginChecksum(msg) != msg.Checksum {
		return fmt.Errorf("login req checksum failed msg:%v", msg.String())
	}

	acc, newPlayer := s.accountMgr.Login(msg, s.storage)

	// 新、旧session不相同做顶号处理
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

	md5 := common.LoginMD5(acc.UUId(), acc.LastRoleId(), newPlayer)
	// 通知玩家登录成功
	return s.Send2Client(gSession, &outer.LoginResp{
		UID:       acc.UUId(),
		RID:       acc.LastRoleId(),
		NewPlayer: newPlayer,
		Checksum:  md5,
	})
}
