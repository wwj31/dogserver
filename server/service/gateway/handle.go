package gateway

import (
	gogo "github.com/gogo/protobuf/proto"
	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
)

// InnerHandler 处理其他服务向gateway发送的消息
func (s *GateWay) InnerHandler(sourceId string, v interface{}) gogo.Message {
	switch msg := v.(type) {
	case *inner.BindSessionWithRID:
		gSession := common.GSession(msg.GetGateSession())
		_, sessionId := gSession.Split()
		session, ok := s.sessions[sessionId]
		if !ok {
			log.Warnw("bind session with rid not found session", "gateway", s.ID(), "gSession", gSession.String())
			return nil
		}
		session.PlayerId = actortype.PlayerId(msg.RID)
		log.Infow("bing session with player", "session", sessionId, "player", session.PlayerId)
		return &outer.Ok{}
	default:
	}
	return nil
}
