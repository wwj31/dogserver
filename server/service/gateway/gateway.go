package gateway

import (
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/iniconfig"
	"github.com/wwj31/dogactor/log"
	"github.com/wwj31/dogactor/network"
	"github.com/wwj31/dogactor/tools"
	"time"

	"server/common"
	"server/proto/inner_message/inner"
)

type GateWay struct {
	actor.Base
	config iniconfig.Config

	// 管理所有对外的玩家tcp连接
	listener network.INetListener
	sessions map[uint32]*UserSession

	// 消息映射表
	msgParser *tools.ProtoParser
}

func New(conf iniconfig.Config) *GateWay {
	return &GateWay{
		config: conf,
	}
}

func (s *GateWay) OnInit() {
	s.sessions = make(map[uint32]*UserSession)

	s.listener = network.StartTcpListen(s.config.String("gate_addr"),
		func() network.ICodec { return &network.StreamCodec{} },
		func() network.INetHandler { return &UserSession{gateway: s} },
	)

	s.msgParser = tools.NewProtoParser().Init("message", "MSG")

	_ = s.System().RegistEvent(s.ID(), (*actor.EvDelactor)(nil))

	s.AddTimer(tools.UUID(), time.Hour, s.checkDeadSession, -1)

	if err := s.listener.Start(); err != nil {
		log.KV("err", err).KV("addr", s.config.String("gate_addr")).Error("gateway listener start err")
		return
	}
	log.Debug("gateway OnInit")
}

// 定期检查并清理死链接
func (s *GateWay) checkDeadSession(dt int64) {
	for id, session := range s.sessions {
		if time.Now().UnixMilli()-session.LeaseTime > int64(time.Hour) {
			session.Stop()
			delete(s.sessions, id)
			log.KV("sesion", id).Warn(" find dead session")
		}
	}
}

// 处理actorcore抛来的事件
func (s *GateWay) OnHandleEvent(event interface{}) {
	switch event.(type) {
	case *actor.EvDelactor:
		evData := event.(*actor.EvDelactor)
		log.KV("remote actor", evData.ActorId).Warn("remote actor is inexistent")
	}
}

// 所有消息，直接转发给用户
func (s *GateWay) OnHandleMessage(sourceId, targetId string, v interface{}) {
	switch msg := v.(type) {
	case *inner.GateMsgWrapper:
		// 用户消息，直接转发给用户
		actorId, sessionId := common.GSession(msg.GateSession).SplitGateSession()
		logInfo := log.Fields{"own": s.ID(), "gateSession": msg.GateSession, "sourceId": sourceId, "msgName": msg.MsgName}
		expect.True(s.ID() == actorId, logInfo)
		userSessionHandler := s.sessions[sessionId]
		if userSessionHandler == nil {
			log.KVs(logInfo).Warn("cannot find sessionId")
			return
		}

		log.KVs(logInfo).Info("server message to user")
		msgId, _ := s.msgParser.MsgNameToId(msg.GetMsgName())
		_ = userSessionHandler.SendMsg(network.CombineMsgWithId(msgId, msg.Data))

	default:
		s.InnerHandler(sourceId, v) // 内部消息，单独处理
	}
}
