package room

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"reflect"
	"server/common/rds"

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
	appId      int32

	Rooms map[int64]*Room
}

func (m *Mgr) OnInit() {
	m.Rooms = make(map[int64]*Room, 8)
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
	if routerErr := router.Dispatch(m, pt); routerErr != nil {
		log.Warnw("roomMgr dispatch the message failed", "err", routerErr)
	}
}

func (m *Mgr) GetRoomId() (int64, error) {
	i64, err := rds.Ins.Incr(context.Background(), rdsop.RoomsIncIdKey()).Result()
	if err != nil {
		log.Errorw("room incr id failed", "err", err)
		return 0, err
	}
	return i64, nil
}

func (m *Mgr) AddRoom(r *Room) error {
	if m.Rooms[r.RoomId] != nil {
		return fmt.Errorf("repeated room id:%v", r.RoomId)
	}
	m.Rooms[r.RoomId] = r
	return nil
}
