package iface

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"server/common"
)

type Gamer interface {
	actor.Actor
	common.Sender
	SaveLoader

	SID() int32
	PlayerMgr() PlayerManager

	RegistMsg(msg proto.Message, handle common.Handle)
}
