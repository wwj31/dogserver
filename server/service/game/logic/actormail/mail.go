package actormail

import (
	"server/common/log"
	"server/service/game/iface"

	"github.com/wwj31/dogactor/actor"
)

func New(serverId uint16, storage iface.StoreLoader) *ActorMail {
	return &ActorMail{sid: serverId, storage: storage}
}

type ActorMail struct {
	actor.Base
	storage iface.StoreLoader
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
