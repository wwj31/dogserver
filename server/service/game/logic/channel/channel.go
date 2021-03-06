package channel

import (
	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/proto/innermsg/inner"

	"github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/expect"

	"github.com/wwj31/dogactor/actor"
)

type id2SessionMap map[actortype.ActorId]common.GSession

func New() *Channel {
	return &Channel{}
}

type Channel struct {
	actor.Base

	sender   common.Sender
	channels map[common.CHANNEL_TYPE]id2SessionMap
}

func (s *Channel) OnInit() {
	s.sender = common.NewSendTools(s)
	s.channels = make(map[common.CHANNEL_TYPE]id2SessionMap)
	s.channels[common.WORLD] = id2SessionMap{}

	log.Debugf("chat server OnInit %v", s.ID())
}

func (s *Channel) OnStop() bool {
	log.Infof("stop chat server %v", s.ID())
	return true
}

func (s *Channel) OnHandleMessage(sourceId, targetId string, v interface{}) {
	switch msg := v.(type) {
	case *inner.LeaveChannelReq:
		s.leave(msg.Channel, msg.ActorId)
	case *inner.MessageToChannel:
		bmsg := common.ProtoUnmarshal(msg.Msgname, msg.Data)
		s.broadcast(msg.Channel, bmsg)
	}
}

func (s *Channel) OnHandleRequest(sourceId, targetId, requestId string, v interface{}) (respErr error) {
	switch msg := v.(type) {
	case *inner.JoinChannelReq:
		s.join(msg.Channel, msg.ActorId, common.GSession(msg.GSession))
		_ = s.Response(requestId, &inner.JoinChannelResp{Error: 0})
	}
	return nil
}

func (s *Channel) join(typ common.CHANNEL_TYPE, playerId actortype.ActorId, gs common.GSession) bool {
	var (
		channel id2SessionMap
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

func (s *Channel) leave(typ common.CHANNEL_TYPE, playerId actortype.ActorId) {
	var (
		smap id2SessionMap
		ok   bool
	)

	defer log.Infow("smap leave ", "player", playerId)

	smap, ok = s.channels[typ]
	if !ok {
		return
	}

	delete(smap, playerId)
}

func (s *Channel) broadcast(typ common.CHANNEL_TYPE, msg proto.Message) {
	var (
		smap id2SessionMap
		ok   bool
	)

	defer log.Infow("smap broadcast ", "msg", msg.String())

	smap, ok = s.channels[typ]
	if !ok {
		return
	}

	for _, gSession := range smap {
		err := s.sender.Send2Client(gSession, msg)
		expect.Nil(err)
	}
}
