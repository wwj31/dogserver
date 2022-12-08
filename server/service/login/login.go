package login

import (
	"fmt"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"server/common"
	"server/common/log"
	"server/proto/outermsg/outer"
	"server/service/game/iface"
)

type Login struct {
	actor.Base
	common.SendTools
	storage iface.StoreLoader
}

func New(s iface.StoreLoader) *Login {
	return &Login{
		storage: s,
	}
}

func (s *Login) OnInit() {
	s.SendTools = common.NewSendTools(s)
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

	s.Login(gSession, msg)
	return nil
}
