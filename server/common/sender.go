package common

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
)

type Sender interface {
	Send2Client(gSession GSession, pb proto.Message) error
	Send2Gate(id ActorId, pb proto.Message) error
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
func (s *SendTools) Send2Client(gSession GSession, pb proto.Message) error {
	if gSession.Invalid() {
		return fmt.Errorf("gSession is invalid %v", gSession)
	}

	gateId, _ := gSession.Split()
	wrap := NewGateWrapperByPb(pb, gSession)
	return s.sender.Send(gateId, wrap)
}

// 发送至网关
func (s *SendTools) Send2Gate(id ActorId, pb proto.Message) error {
	if !IsActorOf(id, GateWay_Actor) {
		return fmt.Errorf("send to gate, but type is not gate id:%v", id)
	}

	if err := s.sender.Send(id, pb); err != nil {
		return fmt.Errorf("%w send to gateway error id:%v", err, id)
	}
	return nil
}