package alliance

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/common/mongodb"
	"server/common/router"
	"time"
)

func New(id int32) *Alliance {
	return &Alliance{allianceId: id, members: make(map[string]*Member)}
}

type (
	Alliance struct {
		actor.Base
		allianceId int32
		members    map[string]*Member
		sessions   map[common.GSession]*Member // 关联登录成功的在线玩家
		masterRID  *Member

		currentMsg      actor.Message
		currentGSession common.GSession
	}
)

func (a *Alliance) OnInit() {
	mongoDBName := fmt.Sprintf("alliance_%v", a.allianceId)
	cur, err := mongodb.Ins.Collection(mongoDBName).Find(context.Background(), bson.M{})
	if err != nil {
		log.Errorw("load all alliance member failed", "err", err)
		return
	}

	var members []*Member
	err = cur.All(context.Background(), &members)
	if err != nil {
		log.Errorw("decode all member failed", "err", err)
		return
	}

	if len(members) == 0 {
		log.Warnw("alliance has no member")
		return
	}

	for _, member := range members {
		a.members[member.RID] = member
		if member.Position == Master {
			a.masterRID = member
		}
		log.Debugf("load member %+v", *member)
	}

	log.Debugf("Alliance OnInit %v members:%v", a.ID(), len(a.members))
}

// 所有消息,处理完统一返回流程
func (a *Alliance) responseHandle(resultMsg any) {
	msg, ok := resultMsg.(proto.Message)
	if !ok {
		return
	}

	// 网关消息，直接将消息转发给session, 其他服务消息，走内部通讯接口
	if actortype.IsActorOf(a.currentMsg.GetSourceId(), actortype.GatewayActor) {
		a.Send2Client(a.currentGSession, msg)
	} else {
		var err error
		if a.currentMsg.GetRequestId() != "" {
			err = a.Response(a.currentMsg.GetRequestId(), msg)
		} else {
			err = a.Send(a.currentMsg.GetSourceId(), msg)
		}

		if err != nil {
			log.Warnw("response to actor failed",
				"source", a.currentMsg.GetSourceId(), "msg name", a.currentMsg.GetMsgName())
		}
	}
}
func (a *Alliance) Send2Client(gSession common.GSession, msg proto.Message) {
	member, ok := a.sessions[gSession]
	if !ok {
		return
	}

	log.Infow("output", "alliance", a.ID(), "msg", reflect.TypeOf(msg), "data", msg.String())
	member.GSession.SendToClient(a, msg)
}

func (a *Alliance) OnStop() bool {
	log.Infof("stop Alliance %v", a.ID())
	return true
}

func (a *Alliance) OnHandle(msg actor.Message) {
	a.currentMsg = msg

	message, _, gSession, err := common.UnwrappedGateMsg(msg.Payload())
	expect.Nil(err)
	a.currentGSession = gSession

	pt, ok := message.(proto.Message)
	if !ok {
		log.Warnw("unknown msg", "msg", reflect.TypeOf(message).String())
		return
	}

	log.Infow("input", "alliance", a.ID(), "msg", reflect.TypeOf(pt), "data", pt.String())
	router.Dispatch(a, pt)
}

func (a *Alliance) PlayerOnline(gSession common.GSession, rid string) {
	member, ok := a.members[rid]
	if !ok {
		log.Warnw("can not find member ", "rid", rid)
		return
	}
	member.GSession = gSession
	member.OnlineAt = time.Now()
	a.sessions[gSession] = member

	log.Infow("player online ", "gSession", gSession, "rid", member.RID, "shortId", member.ShortId)
}

func (a *Alliance) PlayerOffline(gSession common.GSession, rid string) {
	member, ok := a.members[rid]
	if !ok {
		log.Warnw("can not find member ", "rid", rid)
		return
	}

	if gSession != member.GSession {
		log.Warnw("session not equal", "rid", rid, "gSession", member.GSession, "offline gSession", gSession)
	}
	delete(a.sessions, member.GSession)

	member.GSession = ""
	member.OfflineAt = time.Now()

	log.Infow("player offline ", "gSession", gSession, "rid", member.RID, "shortId", member.ShortId)
}
