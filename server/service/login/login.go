package login

import (
	"github.com/wwj31/dogactor/actor"
	"server/proto/message"
)

type Login struct {
	actor.Base
}

func (s *Login) OnInit() {

}

func (s *Login) OnHandleMessage(sourceId, targetId string, msg interface{}) {
	switch msg.(type) {
	case *message.LoginReq:
	}
}
