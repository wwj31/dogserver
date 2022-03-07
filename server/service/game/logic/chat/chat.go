package chat

import (
	"server/common"
	"server/common/log"
	"server/proto/innermsg/inner"

	"github.com/wwj31/dogactor/actor"
)

func New() *ActorChat {
	return &ActorChat{}
}

type ActorChat struct {
	actor.Base

	sender   common.Sender
	channels map[common.CHANNEL_TYPE]*Channel
}

func (s *ActorChat) OnInit() {
	s.sender = common.NewSendTools(s)
	s.channels = make(map[common.CHANNEL_TYPE]*Channel)
	s.channels[common.WORLD] = NewChannel(s.sender)

	log.Debugf("chat server OnInit %v", s.ID())
}

func (s *ActorChat) OnStop() bool {
	log.Infof("stop chat server %v", s.ID())
	return true
}

func (s *ActorChat) OnHandleMessage(sourceId, targetId string, v interface{}) {
	switch msg := v.(type) {
	case *inner.LeaveChatChannelReq:
		if channel, ok := s.channels[msg.Channel]; ok {
			channel.Leave(msg.ActorId)
		}
	case *inner.MessageToChannel:
		if channel, ok := s.channels[msg.Channel]; ok {
			bmsg := common.ProtoUnmarshal(msg.Msgname, msg.Data)
			channel.Broadcast(bmsg)
		}
	}
}

func (s *ActorChat) OnHandleRequest(sourceId, targetId, requestId string, v interface{}) (respErr error) {
	switch msg := v.(type) {
	case *inner.JoinChatChannelReq:
		if channel, ok := s.channels[msg.Channel]; ok {
			channel.Join(msg.ActorId, common.GSession(msg.GSession))
			_ = s.Response(requestId, &inner.JoinChatChannelResp{Error: 0})
		}
	}
	return nil
}
