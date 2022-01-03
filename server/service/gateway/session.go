package gateway

import (
	"time"

	"server/common"
	"server/common/log"
	"server/proto/inner_message/inner"
	"server/proto/message"

	"github.com/wwj31/dogactor/network"
	"github.com/wwj31/dogactor/tools"
)

type UserSession struct {
	gateway   *GateWay
	GameId    common.ActorId // 处理当前session的game
	LeaseTime int64
	network.NetSession
}

func (s *UserSession) OnSessionCreated(sess network.NetSession) {
	s.NetSession = sess
	s.LeaseTime = time.Now().UnixMilli()

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
	if s.GameId != "" {
		// 连接断开，通知game
		gSession := common.GateSession(s.gateway.ID(), s.Id())
		_ = s.gateway.Send(s.GameId, &inner.GT2GSessionClosed{GateSession: gSession.String()})
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

	// 心跳
	if msgId == message.MSG_PING.Int32() {
		ping := network.NewBytesMessageParse(data, s.gateway.msgParser).Proto().(*message.Ping)
		pong := network.NewPbMessage(&message.Pong{
			ClientTimestamp: ping.ClientTimestamp,
			ServerTimestamp: tools.Milliseconds(),
		}, message.MSG_PONG.Int32())
		err = s.SendMsg(pong.Buffer())
		s.LeaseTime = time.Now().UnixMilli()
		return
	}

	msgName, ok := s.gateway.msgParser.MsgIdToName(msgId)
	if !ok {
		log.Errorw("proto not find struct", "msgId", msgId)
		return
	}

	gSession := common.GateSession(s.gateway.ID(), s.Id())
	wrapperMsg := common.NewGateWrapperByBytes(data[4:], msgName, gSession)

	if message.MSG_LOGIN_SEGMENT_BEGIN.Int32() <= msgId && msgId <= message.MSG_LOGIN_SEGMENT_END.Int32() {
		err = s.gateway.Send(common.Login_Actor, wrapperMsg)
	} else if message.MSG_GAME_SEGMENT_BEGIN.Int32() <= msgId && msgId <= message.MSG_GAME_SEGMENT_END.Int32() {
		if s.GameId == "" {
			return
		}
		err = s.gateway.Send(s.GameId, wrapperMsg)
	}
	log.Infow("user msg -> server",
		"msgId", msgId,
		"msgName", msgName,
		"gSession", gSession,
	)
}
