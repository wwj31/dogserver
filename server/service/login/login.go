package login

import (
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/log"
	"server/proto/inner_message"
	"server/proto/message"
)

type Login struct {
	actor.Base
}

func (s *Login) OnInit() {
	log.Debug("login OnInit")
}

func (s *Login) OnHandleMessage(sourceId, targetId string, msg interface{}) {
	v, gateway, err := inner_message.UnwrapperGateMsg(msg)
	expect.Nil(err)

	switch m := v.(type) {
	case *message.LoginReq:
		s.LoginReq(sourceId, gateway, m)
	}
}

func (s *Login) LoginReq(sourceId, gateway string, msg *message.LoginReq) {
	log.Debug(msg.String())
	wrap := inner_message.NewGateWrapperByPb(&message.LoginRsp{}, gateway)
	expect.Nil(s.Send(sourceId, wrap))
}
