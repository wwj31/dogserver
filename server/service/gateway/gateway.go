package gateway

import (
	"time"

	"server/common"
	"server/common/log"
	"server/common/toml"
	"server/proto/innermsg/inner"

	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/network"
	"github.com/wwj31/dogactor/tools"
)

type GateWay struct {
	actor.Base

	// 管理所有对外的玩家tcp连接
	listener network.Listener
	sessions map[uint32]*UserSession

	// 消息映射表
	protoIndex *tools.ProtoIndex
}

func New() *GateWay {
	return &GateWay{}
}

func (s *GateWay) OnInit() {
	s.sessions = make(map[uint32]*UserSession)

	s.listener = network.StartTcpListen(toml.Get("gateaddr"),
		func() network.DecodeEncoder { return &network.StreamCode{MaxDecode: int(10 * common.KB)} },
		func() network.NetSessionHandler { return &UserSession{gateway: s} },
	)

	s.AddTimer(tools.XUID(), tools.Now().Add(time.Hour), s.checkDeadSession, -1)

	if err := s.listener.Start(); err != nil {
		log.Errorw("gateway listener start failed", "err", err, "addr", toml.Get("gateaddr"))
		return
	}
	log.Debugf("gateway OnInit")
}

// 定期检查并清理死链接
func (s *GateWay) checkDeadSession(dt time.Duration) {
	for id, session := range s.sessions {
		if time.Now().UnixMilli()-session.LeaseTime > int64(time.Hour) {
			session.Stop()
			delete(s.sessions, id)
			log.Warnw(" find dead session", "sesion", id)
		}
	}
}

// 所有消息，直接转发给用户
func (s *GateWay) OnHandleMessage(sourceId, targetId string, v interface{}) {
	switch msg := v.(type) {
	case *inner.GateMsgWrapper:
		// 用户消息直接转发前端
		actorId, sessionId := common.GSession(msg.GateSession).Split()
		logInfo := []interface{}{
			"own", s.ID(),
			"gSession", msg.GateSession,
			"sourceId", sourceId,
			"msgName", msg.MsgName}

		if s.ID() != actorId {
			log.Errorw("session disabled gate is not own", logInfo...)
			return
		}
		userSessionHandler := s.sessions[sessionId]
		if userSessionHandler == nil {
			log.Warnw("cannot find sessionId", logInfo...)
			return
		}

		log.Infow("server msg -> user", logInfo...)
		msgId, _ := s.System().ProtoIndex().MsgNameToId(msg.GetMsgName())
		_ = userSessionHandler.SendMsg(network.CombineMsgWithId(msgId, msg.Data))

	default:
		s.InnerHandler(sourceId, v) // 内部消息，单独处理
	}
}
