package room

import (
	"fmt"
	"reflect"

	"github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/actor"

	"server/common/log"
	"server/common/router"
	"server/rdsop"
)

func NewMgr(appId int32) *Mgr {
	return &Mgr{appId: appId}
}

type Mgr struct {
	actor.Base
	currentMsg actor.Message
	incId      int32
	appId      int32

	Rooms map[int32]*Room
}

func (m *Mgr) OnInit() {
	m.Rooms = make(map[int32]*Room, 8)
	router.Result(m, m.responseHandle)
	rdsop.AddRoomMgr(m.appId)
	log.Debugf("RoomMgr OnInit")
}

func (m *Mgr) OnStop() bool {
	rdsop.DelRoomMgr(m.appId)
	log.Debugf("RoomMgr stop %v", m.appId)
	return true
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
		log.Warnw("roomMgr handler msg is not proto",
			"msg", reflect.TypeOf(msg.Payload()).String())
		return
	}

	m.currentMsg = msg
	router.Dispatch(m, pt)
}

func (m *Mgr) GetRoomId() int32 {
	m.incId++
	return m.appId*100000 + m.incId
}

func (m *Mgr) AddRoom(r *Room) error {
	if m.Rooms[r.RoomId] != nil {
		return fmt.Errorf("repeated room id:%v", r.RoomId)
	}
	m.Rooms[r.RoomId] = r
	return nil
}
