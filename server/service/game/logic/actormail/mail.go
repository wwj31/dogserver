package actormail

import (
	"github.com/wwj31/dogactor/actor"
	"server/common/log"
	"server/service/game/iface"
)

func New(serverId uint16, storage iface.SaveLoader) *ActorMail {
	return &ActorMail{sid: serverId, storage: storage}
}

type ActorMail struct {
	actor.Base
	storage iface.SaveLoader
	sid     uint16 // serverId
}

func (s *ActorMail) OnInit() {

	log.Debugf("ActorMail OnInit %v", s.ID())
}

func (s *ActorMail) OnStop() bool {
	log.Infof("stop ActorMail %v", s.ID())
	return true
}

func (s *ActorMail) OnHandleMessage(sourceId, targetId string, msg interface{}) {

}
func (s *ActorMail) OnHandleRequest(sourceId, targetId, requestId string, msg interface{}) (respErr error) {
	return nil
}
