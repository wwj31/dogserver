package gateway

import (
	"github.com/gogo/protobuf/proto"
	"server/proto/outermsg/outer"
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
	log.Infow("gateway OnInit")
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
		userSession := g.sessions[sessionId]
		if userSession == nil {
			log.Warnw("cannot find sessionId", logInfo...)
			return
		}

		log.Infow("server msg -> user", logInfo...)
		msgId, ok := g.System().ProtoIndex().MsgNameToId(msg.GetMsgName())
		if !ok {
			log.Errorw("can not find msgId by name ", "name", msg.GetMsgName())
			return
		}

		data, err := proto.Marshal(&outer.Base{
			MsgId: msgId,
			Data:  msg.Data,
		})

		if err != nil {
			log.Errorw("marshal base failed ",
				"err", err, "player", userSession.PlayerId)
			return
		}

		_ = userSession.SendMsg(data)

	default:
		resp := g.InnerHandler(m) // 内部消息，单独处理
		if resp != nil && m.GetRequestId() != "" {
			//log.Debugw("resp inner msg", "reqId", m.GetRequestId(), "resp", resp)
			if err := g.Response(m.GetRequestId(), resp); err != nil {
				log.Errorw("respone failed", "err", err)
			}
		}
	}
}
