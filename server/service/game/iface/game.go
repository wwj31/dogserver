package iface

import (
	"github.com/wwj31/dogactor/actor"
)

type Gamer interface {
	actor.Actor
	UuidGenerator
	StoreLoader

	SID() uint16
	MsgToPlayer(rid uint64, sid uint16, msg interface{})
}

type UuidGenerator interface {
	GenUuid() uint64
}
