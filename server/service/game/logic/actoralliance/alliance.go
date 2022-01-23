package actoralliance

import (
	"github.com/wwj31/dogactor/actor"
	"server/common/log"
)

func New(serverId uint16) *ActorAlliance {
	return &ActorAlliance{sid: serverId}
}

type ActorAlliance struct {
	actor.Base
	sid uint16 // serverId
}

func (s *ActorAlliance) OnInit() {

	log.Debugf("ActorAlliance OnInit %v", s.ID())
}

func (s *ActorAlliance) OnStop() bool {
	log.Infof("stop ActorAlliance %v", s.ID())
	return true
}

func (s *ActorAlliance) OnHandleMessage(sourceId, targetId string, msg interface{}) {

}
func (s *ActorAlliance) OnHandleRequest(sourceId, targetId, requestId string, msg interface{}) (respErr error) {
	return nil
}
