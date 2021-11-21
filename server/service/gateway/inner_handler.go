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
		s.L2GTSessionAssignGame(msg)
	case *inner.L2GTSessionDisabled: // login通知gate 用户旧session失效
		s.L2GTSessionDisabled(sourceId, msg)
	default:
		return false
	}
	return true
}

// login通知gate登录成功，绑定game server
func (s *GateWay) L2GTSessionAssignGame(msg interface{}) {
	recvData := msg.(*inner.L2GTSessionAssignGame)
	gate, sessionId := common.SplitGateSession(recvData.GateSession)
	fields := log.Fields{"actor": s.GetID(), "gateSession": recvData.GateSession, "game": recvData.GameServerId}
	expect.True(s.GetID() == gate, fields)

	session := s.sessions[sessionId]
	if session == nil {
		log.KVs(fields).Warn("session was closed")
		return
	}
	session.GameId = recvData.GameServerId
	log.KVs(fields).Info("L2GTSessionAssignGame")
}

func (s *GateWay) L2GTSessionDisabled(sourceId string, msg interface{}) {
	recvData := msg.(*inner.L2GTSessionDisabled)
	gate, sessionId := common.SplitGateSession(recvData.GateSession)
	fields := log.Fields{"actor": s.GetID(), "gateSession": recvData.GateSession, "source": sourceId}
	expect.True(s.GetID() == gate, fields)
	session := s.sessions[sessionId]
	if session == nil {
		return
	}

	session.Stop()
	delete(s.sessions, sessionId)
	log.KVs(fields).Info("L2GTUserSessionDisabled")
}
