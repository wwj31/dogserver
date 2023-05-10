package common

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"strings"

	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/actor"

	"server/common/actortype"
	"server/common/log"
)

type GSession string

func GateSession(gateId actortype.ActorId, sessionId uint64) GSession {
	return GSession(fmt.Sprintf("%v:%v", gateId, sessionId))
}

func (s GSession) Split() (gateId actortype.ActorId, sessionId uint64) {
	str := strings.Split(string(s), ":")
	if len(str) != 2 {
		log.Errorw("split failed", "gateSession", s)
		panic(nil)
	}
	gateId = str[0]
	n, e := cast.ToUint64E(str[1])
	if e != nil {
		log.Errorw("split failed", "gateSession", s)
		panic(nil)
	}
	sessionId = n
	return
}

func (s GSession) String() string {
	return string(s)
}

func (s GSession) Invalid() bool {
	return s == ""
}
func (s GSession) Valid() bool {
	return s != ""
}

func (s GSession) SendToClient(sender actor.Messenger, pb proto.Message) {
	if s.Invalid() {
		return
	}

	gateway, _ := s.Split()
	wrap := NewGateWrapperByPb(pb, s)
	if err := sender.Send(gateway, wrap); err != nil {
		log.Errorw("gsession send to client failed", "err", err)
	}
	log.Debugf("send %T", pb)
}
