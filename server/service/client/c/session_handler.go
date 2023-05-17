package c

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/network"
	"server/common/log"
	"server/proto/outermsg/outer"
)

type SessionHandler struct {
	client *Client
}

func (s SessionHandler) OnSessionCreated(session network.Session) {
	log.Infof("session OnSessionCreated!")
}

func (s SessionHandler) OnSessionClosed() {
	log.Infof("session OnSessionClosed!")
}

func (s SessionHandler) OnRecv(bytes []byte) {
	var (
		base = outer.Base{}
		err  error
	)

	err = proto.Unmarshal(bytes, &base)
	expect.Nil(err)

	pb := s.client.System().ProtoIndex().UnmarshalPbMsg(base.MsgId, base.Data)
	err = s.client.Send(s.client.ID(), pb)
	expect.Nil(err)
}
