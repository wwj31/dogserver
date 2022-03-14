package chat

import (
	"server/common"
	"server/proto/innermsg/inner"
	"server/service/game/logic/player/models"

	"github.com/wwj31/dogactor/expect"

	gogo "github.com/gogo/protobuf/proto"
)

type Chat struct {
	models.Model
}

func New(base models.Model) *Chat {
	mod := &Chat{
		Model: base,
	}

	return mod
}

func (s *Chat) OnLogin() {
	s.joinWorld()
}

func (s *Chat) OnLogout() {
	s.leaveWorld()
}

func (s *Chat) SendToChannel(channel string, msg gogo.Message) {
	msg2Chan := &inner.MessageToChannel{
		Channel: channel,
		Msgname: common.ProtoType(msg),
		Data:    common.ProtoMarshal(msg),
	}
	err := s.Player.Send(common.ChatName(s.Player.Gamer().SID()), msg2Chan)
	expect.Nil(err)
}

func (s *Chat) joinWorld() bool {
	msg := &inner.JoinChannelReq{
		Channel:  common.WORLD,
		ActorId:  s.Player.ID(),
		GSession: s.Player.GateSession().String(),
	}
	result, err := s.Player.RequestWait(common.ChatName(s.Player.Gamer().SID()), msg)
	expect.Nil(err)

	resp := result.(*inner.JoinChannelResp)
	return resp.Error == 0
}

func (s *Chat) leaveWorld() {
	msg := &inner.LeaveChannelReq{
		Channel: common.WORLD,
		ActorId: s.Player.ID(),
	}
	err := s.Player.Send(common.ChatName(s.Player.Gamer().SID()), msg)
	expect.Nil(err)

}
