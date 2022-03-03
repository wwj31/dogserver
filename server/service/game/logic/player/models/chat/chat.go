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

func (s *Chat) SendToChannel(channel string, msg gogo.Message) {
	msg2Chan := &inner.MessageToChannel{
		Channel: channel,
		Msgname: common.ProtoType(msg),
		Data:    common.ProtoMarshal(msg),
	}
	err := s.Player.Send(common.ChatName(0), msg2Chan)
	expect.Nil(err)
}
