package gateway

import (
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/log"
	"github.com/wwj31/dogactor/network"
	"github.com/wwj31/dogactor/tools"
	"server/common"
	"server/proto/inner_message"
	"server/proto/inner_message/inner"
	"server/proto/message"
	"time"
)

type UserSession struct {
	gateway   *GateWay
	GameId    common.ActorId // 处理当前session的game
	LeaseTime int64
	network.INetSession
}

func (s *UserSession) OnSessionCreated(sess network.INetSession) {
	s.INetSession = sess
	s.LeaseTime = time.Now().UnixMilli()

	// 这里只做session映射，等待客户端请求登录
	_ = s.gateway.Send(s.gateway.GetID(), func() {
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
		gateSession := common.GateSession(s.gateway.GetID(), s.Id())
		_ = s.gateway.Send(s.GameId, &inner.GT2GSessionClosed{GateSession: gateSession})
	}

	_ = s.gateway.Send(s.gateway.GetID(), func() {
		delete(s.gateway.sessions, s.Id())
	})
}

func (s *UserSession) OnRecv(data []byte) {
	expect.True(len(data) >= 4, log.Fields{"len(data)": len(data), "session": s.Id()})
	msgId := int32(network.Byte4ToUint32(data[:4]))

	var err error
	defer func() {
		if err != nil {
			log.KVs(log.Fields{"err": err, "msgId": msgId}).Error("OnRecv error")
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
		log.KV("msgId", msgId).Error("proto not find struct")
		return
	}

	log.KV("msgId", msgId).KV("msgName", msgName).Info("UserSession OnRecv Msg")

	gateSession := common.GateSession(s.gateway.GetID(), s.Id())
	wrapperMsg := inner_message.NewGateWrapperByBytes(data[4:], msgName, gateSession)
	if message.MSG_LOGIN_SEGMENT_BEGIN.Int32() <= msgId && msgId <= message.MSG_LOGIN_SEGMENT_END.Int32() {
		err = s.gateway.Send(common.Login_Actor, wrapperMsg)
	} else if message.MSG_GAME_SEGMENT_BEGIN.Int32() <= msgId && msgId <= message.MSG_GAME_SEGMENT_END.Int32() {
		expect.True(s.GameId != "", log.Fields{"session": s.Id(), "msgId": msgId})
		err = s.gateway.Send(s.GameId, wrapperMsg)
	}
}
