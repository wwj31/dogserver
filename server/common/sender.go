package common

import (
	"fmt"
	"server/common/actortype"

	gogo "github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
)

// 封装发送用户和网关的sender

type Sender interface {
	Send2Client(gSession GSession, pb gogo.Message) error
	Send2Gate(id actortype.ActorId, pb gogo.Message) error
}

type SendTools struct {
	sender actor.Sender
}

func NewSendTools(s actor.Sender) SendTools {
	return SendTools{
		sender: s,
	}
}

// 发送至前端
func (s SendTools) Send2Client(gSession GSession, pb gogo.Message) error {
	if gSession.Invalid() {
		return nil
	}

	gateId, _ := gSession.Split()
	wrap := NewGateWrapperByPb(pb, ProtoType(pb), gSession)
	return s.sender.Send(gateId, wrap)
}

// 发送至网关
func (s SendTools) Send2Gate(id actortype.ActorId, pb gogo.Message) error {
	if !actortype.IsActorOf(id, actortype.GateWay_Actor) {
		return fmt.Errorf("send to gate, but type is not gate id:%v", id)
	}

	if err := s.sender.Send(id, pb); err != nil {
		return fmt.Errorf("%w send to gateway error id:%v", err, id)
	}
	return nil
}
