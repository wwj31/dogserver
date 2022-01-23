package actorchat

import (
	"github.com/wwj31/dogactor/actor"
	"server/common/log"
)

func New(serverId uint16) *ActorChat {
	return &ActorChat{sid: serverId}
}

type ActorChat struct {
	actor.Base
	sid uint16 // serverId
}

func (s *ActorChat) OnInit() {

	log.Debugf("ChatServer OnInit %v", s.ID())
}

func (s *ActorChat) OnStop() bool {
	log.Infof("stop ChatServer %v", s.ID())
	return true
}

func (s *ActorChat) OnHandleMessage(sourceId, targetId string, msg interface{}) {

}
func (s *ActorChat) OnHandleRequest(sourceId, targetId, requestId string, msg interface{}) (respErr error) {
	return nil
}
