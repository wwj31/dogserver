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
	return &Alliance{allianceId: id}
}

type (
	Alliance struct {
		actor.Base
		allianceId int32
		members    map[string]*Member
		sessions   map[common.GSession]*Member // 关联登录成功的在线玩家
		masterRID  string

		currentMsg      actor.Message
		currentGSession common.GSession
	}
)

func (s *Alliance) OnInit() {
	mongoDBName := fmt.Sprintf("alliance_%v", s.allianceId)
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
		s.members[member.RID] = member
		log.Debugf("load member %+v", *member)
	}

	log.Debugf("Alliance OnInit %v members:%v", s.ID(), len(s.members))
}

// 所有消息,处理完统一返回流程
func (s *Alliance) responseHandle(resultMsg any) {
	msg, ok := resultMsg.(proto.Message)
	if !ok {
		return
	}

	// 网关消息，直接将消息转发给session, 其他服务消息，走内部通讯接口
	if actortype.IsActorOf(s.currentMsg.GetSourceId(), actortype.GatewayActor) {
		s.Send2Client(s.currentGSession, msg)
	} else {
		var err error
		if s.currentMsg.GetRequestId() != "" {
			err = s.Response(s.currentMsg.GetRequestId(), msg)
		} else {
			err = s.Send(s.currentMsg.GetSourceId(), msg)
		}

		if err != nil {
			log.Warnw("response to actor failed",
				"source", s.currentMsg.GetSourceId(), "msg name", s.currentMsg.GetMsgName())
		}
	}
}
func (s *Alliance) Send2Client(gSession common.GSession, msg proto.Message) {
	member, ok := s.sessions[gSession]
	if !ok {
		return
	}

	log.Infow("output", "alliance", s.ID(), "msg", reflect.TypeOf(msg), "data", msg.String())
	member.GSession.SendToClient(s, msg)
}

func (s *Alliance) OnStop() bool {
	log.Infof("stop Alliance %v", s.ID())
	return true
}

func (s *Alliance) OnHandle(msg actor.Message) {
	s.currentMsg = msg

	message, _, gSession, err := common.UnwrappedGateMsg(msg.Payload())
	expect.Nil(err)
	s.currentGSession = gSession

	pt, ok := message.(proto.Message)
	if !ok {
		log.Warnw("unknown msg", "msg", reflect.TypeOf(message).String())
		return
	}

	log.Infow("input", "alliance", s.ID(), "msg", reflect.TypeOf(pt), "data", pt.String())
	router.Dispatch(s, pt)
}

func (s *Alliance) PlayerOnline(gSession common.GSession, rid string) {
	member, ok := s.members[rid]
	if !ok {
		log.Warnw("can not find member ", "rid", rid)
		return
	}
	member.GSession = gSession
	member.OnlineAt = time.Now()
	s.sessions[gSession] = member

	log.Infow("player online ", "gSession", gSession, "rid", member.RID, "shortId", member.ShortId)
}

func (s *Alliance) PlayerOffline(gSession common.GSession, rid string) {
	member, ok := s.members[rid]
	if !ok {
		log.Warnw("can not find member ", "rid", rid)
		return
	}

	if gSession != member.GSession {
		log.Warnw("session not equal", "rid", rid, "gSession", member.GSession, "offline gSession", gSession)
	}
	delete(s.sessions, member.GSession)

	member.GSession = ""
	member.OfflineAt = time.Now()

	log.Infow("player offline ", "gSession", gSession, "rid", member.RID, "shortId", member.ShortId)
}
