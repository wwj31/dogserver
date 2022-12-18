package login

import (
	"fmt"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"server/common"
	"server/common/log"
	"server/proto/outermsg/outer"
)

type Login struct {
	actor.Base
}

func New() *Login {
	return &Login{}
}

func (s *Login) OnInit() {
	log.Infow("login OnInit")
}

func (s *Login) OnStop() bool {
	log.Debugw("login stop", "id", s.ID())
	return true
}

func (s *Login) OnHandle(m actor.Message) {
	rawMsg := m.RawMsg()
	v, _, gSession, err := common.UnwrappedGateMsg(rawMsg)

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
