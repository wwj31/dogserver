package actorchat

import (
	"server/common"
	"server/common/log"
	"server/proto/innermsg/inner"

	"github.com/wwj31/dogactor/actor"
)

func New() *ActorChat {
	return &ActorChat{}
}

type CHANNEL_TYPE = string

const (
	WORLD    CHANNEL_TYPE = "world"    // 世界频道
	ALLIANCE              = "alliance" // 联盟频道
	PRIVATE               = "private"  // 私聊频道
)

type ActorChat struct {
	actor.Base

	channels map[CHANNEL_TYPE]*Channel
}

func (s *ActorChat) OnInit() {
	s.channels = make(map[CHANNEL_TYPE]*Channel)
	s.channels[WORLD] = NewChannel()
	log.Debugf("chat server OnInit %v", s.ID())
}

func (s *ActorChat) OnStop() bool {
	log.Infof("stop chat server %v", s.ID())
	return true
}

func (s *ActorChat) OnHandleMessage(sourceId, targetId string, v interface{}) {
	switch msg := v.(type) {
	case *inner.JoinChatChannel:
		if channel, ok := s.channels[msg.Channel]; ok {
			channel.Join(msg.ActorId, common.GSession(msg.GSession))
		}
	case *inner.LeaveChatChannel:
		if channel, ok := s.channels[msg.Channel]; ok {
			channel.Leave(msg.ActorId)
		}
	}
}
func (s *ActorChat) OnHandleRequest(sourceId, targetId, requestId string, msg interface{}) (respErr error) {
	return nil
}
