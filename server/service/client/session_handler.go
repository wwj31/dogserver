package client

import (
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/network"
	"server/common/log"
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
	msgType := int32(network.Byte4ToUint32(bytes[:4]))
	pb := s.client.System().ProtoIndex().UnmarshalPbMsg(msgType, bytes[4:])

	err := s.client.Send(s.client.ID(), pb)
	expect.Nil(err)
}
