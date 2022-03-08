package channel

import (
	"server/common"
	"server/common/log"
	"server/proto/innermsg/inner"

	"github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/expect"

	"github.com/wwj31/dogactor/actor"
)

type ChannelMap map[common.ActorId]common.GSession

func New() *Channel {
	return &Channel{}
}

type Channel struct {
	actor.Base

	sender   common.Sender
	channels map[common.CHANNEL_TYPE]ChannelMap
}

func (s *Channel) OnInit() {
	s.sender = common.NewSendTools(s)
	s.channels = make(map[common.CHANNEL_TYPE]ChannelMap)
	s.channels[common.WORLD] = ChannelMap{}

	log.Debugf("chat server OnInit %v", s.ID())
}

func (s *Channel) OnStop() bool {
	log.Infof("stop chat server %v", s.ID())
	return true
}

func (s *Channel) OnHandleMessage(sourceId, targetId string, v interface{}) {
	switch msg := v.(type) {
	case *inner.LeaveChatChannelReq:
		s.leave(msg.Channel, msg.ActorId)
	case *inner.MessageToChannel:
		bmsg := common.ProtoUnmarshal(msg.Msgname, msg.Data)
		s.broadcast(msg.Channel, bmsg)
	}
}

func (s *Channel) OnHandleRequest(sourceId, targetId, requestId string, v interface{}) (respErr error) {
	switch msg := v.(type) {
	case *inner.JoinChatChannelReq:
		s.join(msg.Channel, msg.ActorId, common.GSession(msg.GSession))
		_ = s.Response(requestId, &inner.JoinChatChannelResp{Error: 0})
	}
	return nil
}

func (s *Channel) join(typ common.CHANNEL_TYPE, playerId common.ActorId, gs common.GSession) bool {
	var (
		channel ChannelMap
		ok      bool
	)

	defer log.Infow("channel join ", "channelType", typ, "player", playerId, "gs", gs)

	channel, ok = s.channels[typ]
	if !ok {
		return false
	}
	channel[playerId] = gs
	return true
}

func (s *Channel) leave(typ common.CHANNEL_TYPE, playerId common.ActorId) {
	var (
		channel ChannelMap
		ok      bool
	)

	defer log.Infow("channel leave ", "player", playerId)

	channel, ok = s.channels[typ]
	if !ok {
		return
	}

	delete(channel, playerId)
}

func (s *Channel) broadcast(typ common.CHANNEL_TYPE, msg proto.Message) {
	var (
		channel ChannelMap
		ok      bool
	)

	defer log.Infow("channel broadcast ", "msg", msg.String())

	channel, ok = s.channels[typ]
	if !ok {
		return
	}

	for _, gSession := range channel {
		err := s.sender.Send2Client(gSession, msg)
		expect.Nil(err)
	}
}
