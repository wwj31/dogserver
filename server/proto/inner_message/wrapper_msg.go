package inner_message

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/log"
	"github.com/wwj31/dogactor/tools"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"server/common"
	"server/proto/inner_message/inner"
)

func NewGateWrapperByPb(pb proto.Message, gateSession common.GSession) *inner.GateMsgWrapper {
	data, err := proto.Marshal(pb)
	if err != nil {
		log.KV("err", err).ErrorStack(3, "marshal pb failed")
		return nil
	}
	return &inner.GateMsgWrapper{GateSession: gateSession.String(), MsgName: tools.MsgName(pb), Data: data}
}

// 网关封装的消息信息
func NewGateWrapperByBytes(data []byte, msgName string, gateSession common.GSession) *inner.GateMsgWrapper {
	return &inner.GateMsgWrapper{GateSession: gateSession.String(), MsgName: msgName, Data: data}
}

func UnwrapperGateMsg(msg interface{}) (interface{}, common.GSession, error) {
	wrapper, is := msg.(*inner.GateMsgWrapper)
	if !is {
		return msg, "", nil
	}

	tp, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(wrapper.MsgName))
	if err != nil {
		return nil, common.GSession(wrapper.GateSession), err
	}

	actMsg := tp.New().Interface().(proto.Message)
	err = proto.Unmarshal(wrapper.Data, actMsg.(proto.Message))
	if err != nil {
		return nil, common.GSession(wrapper.GateSession), err
	}
	return actMsg, common.GSession(wrapper.GateSession), nil
}
