package common

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/log"
	"strings"
)

type GSession string

func (s GSession) Split() (actorId ActorId, sessionId uint32) {
	strs := strings.Split(string(s), ":")
	if len(strs) != 2 {
		log.KV("gateSession", s).ErrorStack(3, "Split error")
		panic(nil)
	}
	actorId = strs[0]
	sint, e := cast.ToUint32E(strs[1])
	if e != nil {
		log.KV("gateSession", string(s)).ErrorStack(3, "Split error")
		panic(nil)
	}
	sessionId = sint
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

func GateSession(actorID ActorId, sessionId uint32) GSession {
	return GSession(fmt.Sprintf("%v:%v", actorID, sessionId))
}
