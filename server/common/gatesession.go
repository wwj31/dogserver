package common

import (
	"fmt"
	"server/common/log"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/cast"
)

type GSession string

func (s GSession) Split() (gateId ActorId, sessionId uint32) {
	strs := strings.Split(string(s), ":")
	if len(strs) != 2 {
		log.Errorw("split failed", "gateSession", s)
		panic(nil)
	}
	gateId = strs[0]
	sint, e := cast.ToUint32E(strs[1])
	if e != nil {
		log.Errorw("split failed", "gateSession", s)
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

func GateSession(gateId ActorId, sessionId uint32) GSession {
	return GSession(fmt.Sprintf("%v:%v", gateId, sessionId))
}

func NewGateWrapperByPb(pb proto.Message, msgName string, gateSession GSession) *inner.GateMsgWrapper {
	data, err := proto.Marshal(pb)
	if err != nil {
		log.Errorw("marshal pb failed", "err", err)
		return nil
	}
	//return &inner.GateMsgWrapper{GateSession: gateSession.String(), MsgName: tools.MsgName(pb), Data: data}
	return &inner.GateMsgWrapper{GateSession: gateSession.String(), MsgName: msgName, Data: data}
}

// 网关封装的消息信息(避免一次序列化操作)
func NewGateWrapperByBytes(data []byte, msgName string, gateSession GSession) *inner.GateMsgWrapper {
	return &inner.GateMsgWrapper{GateSession: gateSession.String(), MsgName: msgName, Data: data}
}

func UnwrapperGateMsg(msg interface{}) (interface{}, string, GSession, error) {
	wrapper, is := msg.(*inner.GateMsgWrapper)
	if !is {
		return msg, "", "", nil
	}

	v, ok := outer.Spawner(wrapper.MsgName)
	if !ok {
		return nil, "", GSession(wrapper.GateSession), fmt.Errorf("msg not found in outer Spawner %v", wrapper.MsgName)
	}
	actMsg := v.(proto.Message)

	err := proto.Unmarshal(wrapper.Data, actMsg)
	if err != nil {
		return nil, wrapper.MsgName, GSession(wrapper.GateSession), err
	}
	return actMsg, wrapper.MsgName, GSession(wrapper.GateSession), nil
}
