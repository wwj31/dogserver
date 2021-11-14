package login

import (
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/iniconfig"
	"github.com/wwj31/dogactor/log"
	"server/proto/inner_message"
	"server/proto/message"
)

type saveLoader interface {
	Save(key string, data interface{}) error
	Load(key string) interface{}
}
type Login struct {
	actor.Base
	Config iniconfig.Config

	storage saveLoader
}

func New(s saveLoader, conf iniconfig.Config) *Login {
	return &Login{
		storage: s,
		Config:  conf,
	}
}

func (s *Login) OnInit() {
	log.Debug("login OnInit")
}

func (s *Login) OnHandleMessage(sourceId, targetId string, msg interface{}) {
	v, gateway, err := inner_message.UnwrapperGateMsg(msg)
	expect.Nil(err)

	switch msg := v.(type) {
	case *message.LoginReq:
		s.LoginReq(sourceId, gateway, msg)
	}
}

func (s *Login) LoginReq(sourceId, gateway string, msg *message.LoginReq) {
	log.Debug(msg.String())

	// todo ....处理登录消息

	wrap := inner_message.NewGateWrapperByPb(&message.LoginRsp{}, gateway)
	expect.Nil(s.Send(sourceId, wrap))
}
