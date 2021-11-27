package gateway

import (
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/log"
	"server/common"
	"server/proto/inner_message/inner"
)

// 处理其他服务向gateway发送的消息
func (s *GateWay) InnerHandler(sourceId string, v interface{}) bool {
	switch msg := v.(type) {
	case *inner.L2GTSessionAssignGame: // login分配游戏服，通知gate绑定用户gameActor
		s.L2GTSessionAssignGame(msg)
	case *inner.L2GTSessionDisabled: // login通知gate 用户旧session失效
		s.L2GTSessionDisabled(sourceId, msg)
	default:
		return false
	}
	return true
}

func (s *GateWay) L2GTSessionAssignGame(msg *inner.L2GTSessionAssignGame) {
	gate, sessionId := common.GSession(msg.GateSession).SplitGateSession()
	fields := log.Fields{"actor": s.ID(), "gateSession": msg.GateSession, "game": msg.GameServerId}
	expect.True(s.ID() == gate, fields)

	session := s.sessions[sessionId]
	if session == nil {
		log.KVs(fields).Warn("session was closed")
		return
	}
	session.GameId = msg.GameServerId
	log.KVs(fields).Info("L2GTSessionAssignGame")
}

func (s *GateWay) L2GTSessionDisabled(sourceId string, msg *inner.L2GTSessionDisabled) {
	gate, sessionId := common.GSession(msg.GateSession).SplitGateSession()
	fields := log.Fields{"actor": s.ID(), "gateSession": msg.GateSession, "source": sourceId}
	expect.True(s.ID() == gate, fields)
	session := s.sessions[sessionId]
	if session == nil {
		return
	}

	session.Stop()
	delete(s.sessions, sessionId)
	log.KVs(fields).Info("L2GTUserSessionDisabled")
}
