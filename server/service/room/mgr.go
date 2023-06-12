package room

import (
	"reflect"

	"github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/actor"

	"server/common/log"
	"server/common/router"
)

func NewMgr() *Mgr {
	return &Mgr{}
}

type Mgr struct {
	actor.Base
	currentMsg actor.Message
	incId      int32
}

func (m *Mgr) OnInit() {
	router.Result(m, m.responseHandle)
	log.Debugf("RoomMgr OnInit")
}

func (m *Mgr) responseHandle(resultMsg any) {
	msg, ok := resultMsg.(proto.Message)
	if !ok {
		return
	}

	var err error
	if m.currentMsg.GetRequestId() != "" {
		err = m.Response(m.currentMsg.GetRequestId(), msg)
	} else {
		err = m.Send(m.currentMsg.GetSourceId(), msg)
	}

	if err != nil {
		log.Warnw("response to actor failed",
			"source", m.currentMsg.GetSourceId(),
			"msg name", m.currentMsg.GetMsgName())
	}
}

func (m *Mgr) OnHandle(msg actor.Message) {
	pt, ok := msg.Payload().(proto.Message)
	if !ok {
		log.Warnw("roomMgr mgr handler msg is not proto",
			"msg", reflect.TypeOf(msg.Payload()).String())
		return
	}

	m.currentMsg = msg
	router.Dispatch(m, pt)
}

func (m *Mgr) RoomId() int32 {
	m.incId++
	return m.incId
}
