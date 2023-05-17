package gateway

import (
	gogo "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/network"
	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"time"
)

// InnerHandler 处理其他服务向gateway发送的消息
func (g *GateWay) InnerHandler(m actor.Message) gogo.Message {
	payload := m.Payload()
	switch msg := payload.(type) {
	case *inner.BindSessionWithRID:
		gSession := common.GSession(msg.GetGateSession())
		_, sessionId := gSession.Split()
		session, ok := g.sessions[sessionId]
		if !ok {
			log.Warnw("bind session with rid not found session", "gateway", g.ID(), "gSession", gSession.String())
			return nil
		}
		session.PlayerId = actortype.PlayerId(msg.RID)
		log.Infow("bing session with player", "session", sessionId, "player", session.PlayerId)
		return &inner.Ok{}
	case *inner.KickOutReq:
		gSession := common.GSession(msg.GetGateSession())
		_, sessionId := gSession.Split()
		session, ok := g.sessions[sessionId]
		if !ok {
			return &inner.KickOutRsp{}
		}
		if actortype.RID(session.PlayerId) == msg.RID {
			session.PlayerId = ""
			_ = session.SendMsg(g.protoData(&outer.FailRsp{Error: outer.ERROR_REPEAT_LOGIN}))
			g.AddTimer("", time.Now().Add(time.Second), func(dt time.Duration) {
				session.Stop()
			})
			log.Infow("kick player ", "gateway", g.ID(), "session", gSession.String(), "RID", session.PlayerId)
		}
		return &inner.KickOutRsp{}
	}
	return nil
}

func (g *GateWay) protoData(message proto.Message) []byte {
	name := g.System().ProtoIndex().MsgName(message)
	msgId, _ := g.System().ProtoIndex().MsgNameToId(name)
	b, _ := proto.Marshal(message)
	return network.CombineMsgWithId(msgId, b)
}
