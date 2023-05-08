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
	sessions map[uint64]*UserSession

	// 消息映射表
	protoIndex *tools.ProtoIndex
}

func New() *GateWay {
	return &GateWay{}
}

func (g *GateWay) OnInit() {
	g.sessions = make(map[uint64]*UserSession)

	addr := toml.Get("gate_addr")
	g.listener = network.StartWSListen(addr,
		func() network.DecodeEncoder { return &network.WSCode{MaxDecode: 100 * tools.KB} },
		func() network.SessionHandler { return &UserSession{gateway: g} },
	)

	g.AddTimer(tools.XUID(), tools.Now().Add(time.Hour), g.checkDeadSession, -1)

	if err := g.listener.Start(); err != nil {
		log.Errorw("gateway listener start failed", "err", err, "addr", addr)
		return
	}
	log.Infow("gateway OnInit ", "addr", addr)
}

// 定期检查并清理死链接
func (g *GateWay) checkDeadSession(dt time.Duration) {
	for id, session := range g.sessions {
		if time.Now().Sub(session.KeepLive) > time.Hour {
			session.Stop()
			delete(g.sessions, id)
			log.Warnw(" find dead session", "sesion", id)
		}
	}
}

// OnHandle 主要转发消息至玩家client，少量内部消息处理
func (g *GateWay) OnHandle(m actor.Message) {
	payload := m.Payload()
	switch msg := payload.(type) {
	case *inner.GateMsgWrapper:
		// 用户消息直接转发前端
		actorId, sessionId := common.GSession(msg.GateSession).Split()
		logInfo := []interface{}{
			"gSession", msg.GateSession,
			"sourceId", m.GetSourceId(),
			"msgName", msg.MsgName,
		}

		if g.ID() != actorId {
			log.Errorw("session disabled gate is not own", logInfo...)
			return
		}
		userSessionHandler := g.sessions[sessionId]
		if userSessionHandler == nil {
			log.Warnw("cannot find sessionId", logInfo...)
			return
		}

		log.Infow("server msg -> user", logInfo...)
		msgId, _ := g.System().ProtoIndex().MsgNameToId(msg.GetMsgName())
		_ = userSessionHandler.SendMsg(network.CombineMsgWithId(msgId, msg.Data))

	default:
		resp := g.InnerHandler(m) // 内部消息，单独处理
		if resp != nil && m.GetRequestId() != "" {
			log.Debugw("resp ", "reqId", m.GetRequestId())
			if err := g.Response(m.GetRequestId(), resp); err != nil {
				log.Errorw("respone failed", "err", err)
			}
		} //wait_cebjpknm1tui4lpi2eh0@1670855890094264048@gateway_1_Actor#:8888
	}
}
