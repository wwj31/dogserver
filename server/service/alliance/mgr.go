package alliance

import (
	"context"
	"github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"server/common/actortype"
	"server/common/log"
	"server/common/mongodb"
	"server/common/router"
)

const Coll = "alliancemgr"

func NewMgr() *Mgr {
	return &Mgr{}
}

type Mgr struct {
	actor.Base
	alliances  []int32
	currentMsg actor.Message
}

func (m *Mgr) OnInit() {
	cur, err := mongodb.Ins.Collection(Coll).Find(context.Background(), bson.M{})
	if err != nil {
		log.Errorw("load all alliance member failed", "err", err)
		return
	}

	var allianceIds []int32
	err = cur.All(context.Background(), &allianceIds)
	if err != nil {
		log.Errorw("decode all member failed", "err", err)
		return
	}

	m.alliances = allianceIds

	for _, allianceId := range m.alliances {
		_ = m.System().NewActor(
			actortype.AllianceName(allianceId),
			New(allianceId),
			actor.SetMailBoxSize(1000),
		)
	}

	router.Result(m, m.responseHandle)
	log.Debugf("AllianceMgr OnInit  alliances:%v", allianceIds)
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
		log.Warnw("alliance mgr handler msg is not proto",
			"msg", reflect.TypeOf(msg.Payload()).String())
		return
	}

	router.Dispatch(m, pt)
}
