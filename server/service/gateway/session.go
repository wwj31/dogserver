package gateway

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/expect"
	"time"

	"github.com/wwj31/dogactor/actor"

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
		log.Infow("session opened ", "sessionId", u.Id())
	})
}

func (u *UserSession) OnSessionClosed() {
	log.Infow("session closed ", "sessionId", u.Id(), "player", u.PlayerId)
	if u.PlayerId != "" {
		gSession := common.GateSession(u.gateway.ID(), u.Id())
		_ = u.gateway.Send(u.PlayerId, &inner.GSessionClosed{
			GateSession: gSession.String(),
		})
	}

	_ = u.gateway.Send(u.gateway.ID(), func() {
		delete(u.gateway.sessions, u.Id())
	})
}

func (u *UserSession) OnRecv(data []byte) {
	var (
		base = &outer.Base{}
		err  error
	)
	err = proto.Unmarshal(data, base)
	if err != nil {
		log.Warnw("base unmarshal failed", "session", u.Id(), "player", u.PlayerId)
		return
	}

	protoIndex := u.gateway.System().ProtoIndex()
	// 心跳
	if base.MsgId == outer.Msg_IdHeartReq.Int32() {
		heartReq := &outer.HeartReq{}
		_ = proto.Unmarshal(base.Data, heartReq)
		rsp := &outer.HeartRsp{
			ClientTimestamp: heartReq.ClientTimestamp,
			ServerTimestamp: tools.Now().UnixMilli(),
		}

		heartRsp, _ := proto.Marshal(rsp)
		pong, bErr := proto.Marshal(&outer.Base{
			MsgId: outer.Msg_IdHeartRsp.Int32(),
			Data:  heartRsp,
		})
		expect.Nil(bErr)

		err = u.SendMsg(pong)
		u.KeepLive = time.Now()
		log.Infow("heart ", "req", heartReq.String(), "rsp", rsp.String())
		return
	}

	msgName, ok := protoIndex.MsgIdToName(base.MsgId)
	if !ok {
		log.Errorw("proto not find struct", "msgId", base.MsgId)
		return
	}

	gSession := common.GateSession(u.gateway.ID(), u.Id())
	wrapperMsg := common.NewGateWrapperByBytes(base.Data, msgName, gSession)

	log.Infow("user msg -> server",
		"msgId", base.MsgId,
		"msgName", msgName,
		"gSession", gSession,
		"player", u.PlayerId,
	)

	var targetId actor.Id
	switch tag := outer.MsgIDTags[base.MsgId]; tag {
	case actortype.LoginActor:
		targetId = actortype.LoginActor
	case actortype.PlayerActor:
		targetId = u.PlayerId
	default:
		log.Errorw(" the message has no target for dispatch",
			"msgId", base.MsgId, "tag", tag)
		return
	}

	err = u.gateway.Send(targetId, wrapperMsg)

}
