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
	v, gSession, err := inner_message.UnwrapperGateMsg(msg)
	expect.Nil(err)

	switch msg := v.(type) {
	case *message.LoginReq:
		s.LoginReq(sourceId, gSession, msg)
	}
}

func (s *Login) LoginReq(sourceId, gSession string, msg *message.LoginReq) {
	log.Debug(msg.String())

	s.login(msg)
	// todo ....处理登录消息
	s.send(sourceId, &message.LoginRsp{})
}

func (s *Login) login(*message.LoginReq) {

}

func (s *Login) send(gSession string, pb proto.Message) {
	gateId, _ := common.SplitGateSession(gSession)
	wrap := inner_message.NewGateWrapperByPb(pb, gSession)
	expect.Nil(s.Send(gateId, wrap))
}
