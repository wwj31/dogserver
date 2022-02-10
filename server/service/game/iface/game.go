package iface

import (
	"github.com/wwj31/dogactor/actor"
)

type Gamer interface {
	actor.Actor
	UuidGenerator
	StoreLoader

	SID() uint16
}

type UuidGenerator interface {
	GenUuid() uint64
}
