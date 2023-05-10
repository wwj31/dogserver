package gateway

import (
	"github.com/wwj31/dogactor/actor"
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

func (u *UserSession) OnSessionCreated(s network.Session) {
	u.Session = s
	u.KeepLive = time.Now()

	_ = u.gateway.Send(u.gateway.ID(), func() {
		u.gateway.sessions[u.Id()] = u
	})
}

func (u *UserSession) OnSessionClosed() {
	if u.PlayerId != "" {
		// 通知player
		_ = u.gateway.Send(u.PlayerId, &inner.GSessionClosed{})
	}

	_ = u.gateway.Send(u.gateway.ID(), func() {
		delete(u.gateway.sessions, u.Id())
	})
}

func (u *UserSession) OnRecv(data []byte) {
	if len(data) < 4 {
		log.Warnw("invalid data len", "len(data)", len(data), "session", u.Id())
	}

	msgId := int32(network.Byte4ToUint32(data[:4]))

	var err error
	defer func() {
		if err != nil {
			log.Errorw("OnRecv error", "err", err, "msgId", msgId)
		}
	}()

	protoIndex := u.gateway.System().ProtoIndex()
	// 心跳
	if msgId == outer.Msg_IdHeartReq.Int32() {
		ping := network.NewBytesMessageParse(data, protoIndex).Proto().(*outer.HeartReq)
		pong := network.NewPbMessage(&outer.HeartRsp{
			ClientTimestamp: ping.ClientTimestamp,
			ServerTimestamp: tools.Now().UnixMilli(),
		}, outer.Msg_IdHeartRsp.Int32())
		err = u.SendMsg(pong.Buffer())
		u.KeepLive = time.Now()
		return
	}

	msgName, ok := protoIndex.MsgIdToName(msgId)
	if !ok {
		log.Errorw("proto not find struct", "msgId", msgId)
		return
	}
	gSession := common.GateSession(u.gateway.ID(), u.Id())
	wrapperMsg := common.NewGateWrapperByBytes(data[4:], msgName, gSession)

	var targetId actor.Id
	switch tag := outer.MsgIDTags[msgId]; tag {
	case actortype.LoginActor:
		targetId = actortype.LoginActor
	case actortype.PlayerActor:
		targetId = u.PlayerId
	default:
		log.Errorw("cannot find the message tag; the message has no target for dispatch", "msgId", msgId, "tag", tag)
		return
	}

	err = u.gateway.Send(targetId, wrapperMsg)

	log.Infow("user msg -> server",
		"msgId", msgId,
		"msgName", msgName,
		"gSession", gSession,
		"player", u.PlayerId,
	)
}
