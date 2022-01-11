package client

import (
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/network"
)

type SessionHandler struct {
	client *Client
}

func (s SessionHandler) OnSessionCreated(session network.NetSession) {
	logger.Infof("session OnSessionCreated!")
}

func (s SessionHandler) OnSessionClosed() {
	logger.Infof("session OnSessionClosed!")
}

func (s SessionHandler) OnRecv(bytes []byte) {
	msgType := int32(network.Byte4ToUint32(bytes[:4]))
	pb := s.client.msgParser.UnmarshalPbMsg(msgType, bytes[4:])
	err := s.client.Send(s.client.ID(), pb)
	expect.Nil(err)
}
