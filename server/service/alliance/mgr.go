package alliance

import (
	"context"
	"fmt"
	"reflect"

	"github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"server/common/actortype"
	"server/common/log"
	"server/common/mongodb"
	"server/common/router"
	"server/proto/innermsg/inner"
	"server/rdsop"
)

const Coll = "alliancemgr"

func NewMgr() *Mgr {
	return &Mgr{}
}

type AlliBaseInfo struct {
	AllianceId    int32 `bson:"_id"`
	MasterShortId int64 `bson:"master_short_id"`
	Disband       bool  `bson:"disband"`
}
type Mgr struct {
	actor.Base
	alliances  []*AlliBaseInfo
	currentMsg actor.Message
	incId      int32
}

func (m *Mgr) OnInit() {
	cur, err := mongodb.Ins.Collection(Coll).Find(context.Background(), bson.M{})
	if err != nil {
		log.Errorw("load all alliance member failed", "err", err)
		return
	}
	for cur.Next(context.Background()) {
		alli := &AlliBaseInfo{}
		err = cur.Decode(alli)
		if err != nil {
			log.Warnw("alliance mgr decode failed", "err", err)
			continue
		}

		_ = m.System().NewActor(
			actortype.AllianceName(alli.AllianceId),
			New(alli.AllianceId),
			actor.SetMailBoxSize(1000),
		)

		if alli.AllianceId > m.incId {
			m.incId = alli.AllianceId
		}

		m.alliances = append(m.alliances, alli)
	}

	router.Result(m, m.responseHandle)
	log.Debugf("AllianceMgr OnInit  alliances:%v", m.alliances)
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

	m.currentMsg = msg
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

	allianceActor := actortype.AllianceName(allianceId)
	err := m.System().NewActor(allianceActor, newAlliance, actor.SetMailBoxSize(1000))
	if err != nil {
		log.Errorw("create alliance failed", "err", err, "msg")
		return 0, err
	}

	// 请求加入盟主
	result, joinErr := m.RequestWait(allianceActor, &inner.AddMemberReq{
		Player:   &masterInfo,
		Position: Master.Int32(),
		Ntf:      true,
	})

	if joinErr != nil {
		log.Errorw("create alliance success,but master set failed", "err", joinErr, "masterInfo", masterInfo.String())
		return 0, joinErr
	}

	if _, ok := result.(*inner.AddMemberRsp); !ok {
		err = fmt.Errorf("create alliance success,but master set failed by assert")
		log.Errorw("create alliance success,but master set failed by assert",
			"type", reflect.TypeOf(result).String(), "masterInfo", masterInfo.String())
		return 0, err
	}

	// 维护联盟列表
	newAlli := &AlliBaseInfo{
		AllianceId:    allianceId,
		MasterShortId: masterShortId,
		Disband:       false,
	}
	m.alliances = append(m.alliances, newAlli)

	_, err = mongodb.Ins.Collection(Coll).UpdateByID(context.Background(),
		allianceId,
		bson.M{"$set": newAlli},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Warnw("update new alliance to mongo failed", "err", err)
	}
	return allianceId, nil
}

func (m *Mgr) Disband(allianceId int32) {
	for k, alli := range m.alliances {
		if alli.AllianceId == allianceId {
			alli.Disband = true
			_, err := mongodb.Ins.Collection(Coll).UpdateByID(context.Background(), allianceId, bson.M{"$set": alli})

			if err != nil {
				log.Warnw("disband alliance failed", "err", err, "id", allianceId)
			}

			m.alliances = append(m.alliances[:k], m.alliances[k+1:]...)
			break
		}
	}
}

func (m *Mgr) getIncId() int32 {
	m.incId++
	return m.incId
}
