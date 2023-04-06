package gateway

import (
	"time"

	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"

	"github.com/wwj31/dogactor/network"
	"github.com/wwj31/dogactor/tools"
)

type UserSession struct {
	gateway  *GateWay
	PlayerId actortype.ActorId // 关联的player
	KeepLive time.Time
	network.Session
}

func (s *UserSession) OnSessionCreated(sess network.Session) {
	s.Session = sess
	s.KeepLive = time.Now()

	// 这里只做session映射，等待客户端请求登录
	_ = s.gateway.Send(s.gateway.ID(), func() {
		// 黑名单判断
		//ip := s.RemoteIP()
		//ret := s.gateway.CallLua("IPFilter", 1, lua.LString(ip))
		//if len(ret) > 0 && !lua.LVIsFalse(ret[0]) {
		//	s.Stop()
		//	log.KV("ip", s.RemoteIP()).Warn("ip filter")
		//	return
		//}
		s.gateway.sessions[s.Id()] = s
	})
}

func (s *UserSession) OnSessionClosed() {
	if s.PlayerId != "" {
		// 通知player
		_ = s.gateway.Send(s.PlayerId, &inner.GSessionClosed{})
	}

	_ = s.gateway.Send(s.gateway.ID(), func() {
		delete(s.gateway.sessions, s.Id())
	})
}

func (s *UserSession) OnRecv(data []byte) {
	if len(data) < 4 {
		log.Warnw("invalid data len", "len(data)", len(data), "session", s.Id())
	}

	msgId := int32(network.Byte4ToUint32(data[:4]))

	var err error
	defer func() {
		if err != nil {
			log.Errorw("OnRecv error", "err", err, "msgId", msgId)
		}
	}()

	protoIndex := s.gateway.System().ProtoIndex()
	// 心跳
	if msgId == outer.Msg_IdPingReq.Int32() {
		ping := network.NewBytesMessageParse(data, protoIndex).Proto().(*outer.PingReq)
		pong := network.NewPbMessage(&outer.PongRsp{
			ClientTimestamp: ping.ClientTimestamp,
			ServerTimestamp: tools.Now().UnixMilli(),
		}, outer.Msg_IdPongRsp.Int32())
		err = s.SendMsg(pong.Buffer())
		s.KeepLive = time.Now()
		return
	}

	msgName, ok := protoIndex.MsgIdToName(msgId)
	if !ok {
		log.Errorw("proto not find struct", "msgId", msgId)
		return
	}
	gSession := common.GateSession(s.gateway.ID(), s.Id())
	wrapperMsg := common.NewGateWrapperByBytes(data[4:], msgName, gSession)

	switch tag := outer.MsgIDTags[msgId]; tag {
	case actortype.LoginActor:
		err = s.gateway.Send(actortype.LoginActor, wrapperMsg)
	case actortype.PlayerActor:
		err = s.gateway.Send(s.PlayerId, wrapperMsg)
	default:
		log.Errorw("cannot find the message tag; the message has no target for dispatch", "msgId", msgId, "tag", tag)
		return
	}

	log.Infow("user msg -> server",
		"msgId", msgId,
		"msgName", msgName,
		"gSession", gSession,
		"player", s.PlayerId,
	)
}
