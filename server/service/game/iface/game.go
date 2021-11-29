package iface

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"server/common"
)

type Gamer interface {
	actor.Actor
	common.Sender

	SID() int32
	RegistMsg(msg proto.Message, handle common.Handle)
}
