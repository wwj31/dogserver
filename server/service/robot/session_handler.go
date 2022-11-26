package robot

import (
	"server/common/log"

	"github.com/wwj31/dogactor/network"
)

type SessionHandler struct {
	client *Robot
}

func (s SessionHandler) OnSessionCreated(session network.Session) {
}

func (s SessionHandler) OnSessionClosed() {
}

func (s SessionHandler) OnRecv(bytes []byte) {
	msgType := int32(network.Byte4ToUint32(bytes[:4]))
	pb := s.client.System().ProtoIndex().UnmarshalPbMsg(msgType, bytes[4:])
	log.Debugw("recv msg", "msgType", msgType, "pb", pb.String())
}
