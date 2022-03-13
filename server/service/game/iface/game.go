package iface

import (
	gogo "github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
)

type Gamer interface {
	actor.Actor
	UuidGenerator
	StoreLoader

	SID() uint16
	MsgToPlayer(rid uint64, sid uint16, msg gogo.Message)
}

type UuidGenerator interface {
	GenUuid() uint64
}
