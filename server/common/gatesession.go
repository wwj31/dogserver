package common

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/tools"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"server/common/log"
	"server/proto/inner_message/inner"
	"strings"
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

func NewGateWrapperByPb(pb proto.Message, gateSession GSession) *inner.GateMsgWrapper {
	data, err := proto.Marshal(pb)
	if err != nil {
		log.Errorw("marshal pb failed", "err", err)
		return nil
	}
	return &inner.GateMsgWrapper{GateSession: gateSession.String(), MsgName: tools.MsgName(pb), Data: data}
}

// 网关封装的消息信息(避免一次序列化操作)
func NewGateWrapperByBytes(data []byte, msgName string, gateSession GSession) *inner.GateMsgWrapper {
	return &inner.GateMsgWrapper{GateSession: gateSession.String(), MsgName: msgName, Data: data}
}

func UnwrapperGateMsg(msg interface{}) (interface{}, GSession, error) {
	wrapper, is := msg.(*inner.GateMsgWrapper)
	if !is {
		return msg, "", nil
	}

	tp, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(wrapper.MsgName))
	if err != nil {
		return nil, GSession(wrapper.GateSession), err
	}

	actMsg := tp.New().Interface().(proto.Message)
	err = proto.Unmarshal(wrapper.Data, actMsg.(proto.Message))
	if err != nil {
		return nil, GSession(wrapper.GateSession), err
	}
	return actMsg, GSession(wrapper.GateSession), nil
}
