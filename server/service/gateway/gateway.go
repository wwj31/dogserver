package gateway

import (
	"time"

	"server/common"
	"server/common/log"
	"server/common/toml"
	"server/proto/inner_message/inner"

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
	msgParser *tools.ProtoParser
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

	s.msgParser = tools.NewProtoParser().Init("message", "MSG")

	_ = s.System().RegistEvent(s.ID(), (*actor.EvDelactor)(nil))

	s.AddTimer(tools.UUID(), tools.NowTime()+int64(time.Hour), s.checkDeadSession, -1)

	if err := s.listener.Start(); err != nil {
		log.Errorw("gateway listener start failed", "err", err, "addr", toml.Get("gateaddr"))
		return
	}
	log.Debugf("gateway OnInit")
}
func (s *GateWay) OnStop() bool {
	s.System().CancelAll(s.ID())
	return true
}

// 定期检查并清理死链接
func (s *GateWay) checkDeadSession(dt int64) {
	for id, session := range s.sessions {
		if time.Now().UnixMilli()-session.LeaseTime > int64(time.Hour) {
			session.Stop()
			delete(s.sessions, id)
			log.Warnw(" find dead session", "sesion", id)
		}
	}
}

// 处理actorcore抛来的事件
func (s *GateWay) OnHandleEvent(event interface{}) {
	switch event.(type) {
	case *actor.EvDelactor:
		evData := event.(*actor.EvDelactor)
		log.Warnw("remote actor is inexistent", "remote actor", evData.ActorId)
	}
}

// 所有消息，直接转发给用户
func (s *GateWay) OnHandleMessage(sourceId, targetId string, v interface{}) {
	switch msg := v.(type) {
	case *inner.GateMsgWrapper:
		// 用户消息，直接转发给用户
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
		msgId, _ := s.msgParser.MsgNameToId(msg.GetMsgName())
		_ = userSessionHandler.SendMsg(network.CombineMsgWithId(msgId, msg.Data))

	default:
		s.InnerHandler(sourceId, v) // 内部消息，单独处理
	}
}
