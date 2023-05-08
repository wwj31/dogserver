package gateway

import (
	gogo "github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/actor"

	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
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
		return &outer.Ok{}
	case *inner.KickOutReq:
		gSession := common.GSession(msg.GetGateSession())
		_, sessionId := gSession.Split()
		session, ok := g.sessions[sessionId]
		if !ok {
			log.Warnw("bind session with rid not found session", "gateway", g.ID(), "gSession", gSession.String())
			return &inner.KickOutRsp{}
		}
		if session.PlayerId == msg.RID {
			session.PlayerId = ""
			session.Stop()
			log.Infow("kick player ", "gateway", g.ID(), "session", gSession.String(), "RID", session.PlayerId)
		}
		return &inner.KickOutRsp{}
	}
	return nil
}
