package iface

import (
	"github.com/wwj31/dogactor/actor"
)

type Gamer interface {
	actor.Actor
	SID() int32
}
