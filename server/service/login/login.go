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
	storage    iface.StoreLoader
	accountMgr account.Mgr
}

func New(s iface.StoreLoader) *Login {
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

func (s *Login) OnHandle(m actor.Message) {
	rawMsg := m.RawMsg()
	v, _, gSession, err := common.UnwrapperGateMsg(rawMsg)

	expect.Nil(err)
	switch msg := v.(type) {
	case *outer.LoginReq:
		err = s.LoginReq(m.GetSourceId(), gSession, msg)
	default:
		err = fmt.Errorf("undefined localmsg type %v", msg)
	}

	if err != nil {
		log.Errorw("handle outer error", "err", err)
	}
}

func (s *Login) LoginReq(sourceId string, gSession common.GSession, msg *outer.LoginReq) error {
	log.Debugf(msg.String())

	if common.LoginToken(msg) != msg.Token {
		return fmt.Errorf("login req token failed msg:%v", msg.String())
	}

	acc, newPlayer := s.accountMgr.Login(msg, s.storage)

	// 通知game拉起player
	err := s.Send(acc.Game(), &inner.PullPlayer{RID: acc.LastRoleId()})
	if err != nil {
		return err
	}

	md5 := common.LoginMD5(acc.UUId(), acc.LastRoleId(), newPlayer)
	// 通知玩家登录成功
	return s.Send2Client(gSession, &outer.LoginResp{
		UID:       acc.UUId(),
		RID:       acc.LastRoleId(),
		NewPlayer: newPlayer,
		Token:     md5,
	})
}
