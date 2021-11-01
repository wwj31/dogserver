package gateway

import (
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/log"
	"server/common"
	"server/proto/inner_message/inner"
)

// 处理其他服务向gateway发送的消息
func (s *GateWay) InnerHandler(sourceId string, msg interface{}) bool {
	switch msg.(type) {
	case *inner.L2GTSessionAssignGame: // login分配游戏服，通知gate绑定用户gameActor
		s.L2GTSessionAssignGame(sourceId, msg)
	case *inner.L2GTUserSessionDisabled: // login通知gate 用户旧session失效
		s.L2GTUserSessionDisabled(sourceId, msg)
	default:
		return false
	}
	return true
}

func (s *GateWay) L2GTSessionAssignGame(sourceId string, msg interface{}) {
	recvData := msg.(*inner.L2GTSessionAssignGame)
	gate, sessionId := common.SplitGateSession(recvData.GateSession)
	fields := log.Fields{"actor": s.GetID(), "gateSession": recvData.GateSession, "source": sourceId}
	expect.True(s.GetID() == gate, fields)

	session := s.sessions[sessionId]
	if session == nil {
		log.KVs(fields).Warn("session was closed")
		return
	}
	session.GameActor = recvData.GameActorId
	log.KVs(log.Fields{"session": recvData.GateSession, "game actorId": recvData.GameActorId}).Info("INNER_MSG_L2G_SESSION_ASSIGN_GAME")
}

func (s *GateWay) L2GTUserSessionDisabled(sourceId string, msg interface{}) {
	recvData := msg.(*inner.L2GTUserSessionDisabled)
	gate, sessionId := common.SplitGateSession(recvData.GetGateSession())
	fields := log.Fields{"actor": s.GetID(), "gateSession": recvData.GateSession, "source": sourceId}
	expect.True(s.GetID() == gate, fields)

	if s.GetID() != gate {
		log.KVs(fields).Error("exception s.GetID() != gate local")
		return
	}
	session := s.sessions[sessionId]
	if session == nil {
		log.KVs(fields).Warn("recvice disabled session proto,but session is nil")
		return
	}
	session.Stop()
	log.KVs(fields).Info("INNER_MSG_L2GT_USER_SESSION_DISABLED ")
}
