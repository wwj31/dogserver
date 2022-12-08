package gateway

import (
	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/proto/innermsg/inner"
)

// InnerHandler 处理其他服务向gateway发送的消息
func (s *GateWay) InnerHandler(sourceId string, v interface{}) {
	switch msg := v.(type) {
	case *inner.BindSessionWithRID:
		gSession := common.GSession(msg.GetGateSession())
		_, sessionId := gSession.Split()
		session, ok := s.sessions[sessionId]
		if !ok {
			log.Warnw("bind session with rid not found session", "gateway", s.ID(), "gSession", gSession.String())
			return
		}
		session.PlayerId = actortype.PlayerId(msg.RID)
	default:
	}
}
