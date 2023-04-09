package common

import (
	"fmt"

	"github.com/gogo/protobuf/proto"

	"server/common/log"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
)

func NewGateWrapperByPb(pb proto.Message, gateSession GSession) *inner.GateMsgWrapper {
	data, err := proto.Marshal(pb)
	if err != nil {
		log.Errorw("marshal pb failed", "err", err)
		return nil
	}
	return &inner.GateMsgWrapper{GateSession: gateSession.String(), MsgName: ProtoType(pb), Data: data}
}

func NewGateWrapperByBytes(data []byte, msgName string, gateSession GSession) *inner.GateMsgWrapper {
	return &inner.GateMsgWrapper{GateSession: gateSession.String(), MsgName: msgName, Data: data}
}

func UnwrappedGateMsg(msg interface{}) (interface{}, string, GSession, error) {
	wrapper, is := msg.(*inner.GateMsgWrapper)
	if !is {
		return msg, "", "", nil
	}

	//tp, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(wrapper.MsgName))
	//if err != nil {
	//	return nil, GSession(wrapper.GateSession), err
	//}
	//
	//actMsg := tp.New().Interface().(proto.Message)
	v, ok := outer.Spawner(wrapper.MsgName, true)
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
