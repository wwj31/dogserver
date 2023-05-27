package alliance

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"server/common/actortype"
	"server/common/log"
	"server/common/mongodb"
	"server/common/rds"
	"server/common/router"
	"server/proto/innermsg/inner"
	"server/rdsop"
)

const Coll = "alliancemgr"

func NewMgr() *Mgr {
	return &Mgr{}
}

type Mgr struct {
	actor.Base
	alliances  []int32
	currentMsg actor.Message
	incId      int32
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
		if allianceId > m.incId {
			m.incId = allianceId
		}
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

func (m *Mgr) CreateAlliance(masterShortId int64) (int32, error) {
	allianceId := m.getIncId()
	newAlliance := New(allianceId)
	masterInfo := rdsop.PlayerInfo(masterShortId)
	if masterInfo.RID == "" {
		return 0, fmt.Errorf("cannot find player by shortId %v", masterInfo)
	}

	// 玩家已有联盟
	if masterInfo.AllianceId != 0 {
		return 0, fmt.Errorf("the player has join alliance playerInfo:%+v", masterInfo)
	}

	// 盟主不能有上级
	if masterInfo.UpShortId != 0 {
		return 0, fmt.Errorf("the player has upShortId playerInfo:%+v", masterInfo)
	}

	newAlliance.SetMember(&masterInfo, Master)
	err := m.System().NewActor(
		actortype.AllianceName(allianceId), newAlliance,
		actor.SetMailBoxSize(1000),
	)

	if err != nil {
		log.Errorw("create alliance failed", "err", err, "msg")
		return 0, err
	}

	if masterInfo.GSession != "" {
		// 玩家在线，通知Player actor修改联盟id，
		m.Send(actortype.PlayerId(masterInfo.RID), &inner.JoinAllianceNtf{
			AllianceId: allianceId,
			Position:   Master.Int32(),
		})
	} else {
		// 玩家不在线，记录到redis中，等待下次上线更新
		key := rdsop.JoinAllianceKey(masterShortId)
		rds.Ins.Set(context.Background(), key, allianceId, 0)
	}

	// 更新玩家公共数据
	masterInfo.Position = Master.Int32()
	masterInfo.AllianceId = allianceId
	rdsop.SetPlayerInfo(&masterInfo)

	return allianceId, nil
}

func (m *Mgr) getIncId() int32 {
	m.incId++
	return m.incId
}
