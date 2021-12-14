package iface

import (
	"github.com/wwj31/dogactor/actor"
	"server/common"
)

type Gamer interface {
	actor.Actor
	common.Sender
	SaveLoader

	SID() int32
	PlayerMgr() PlayerManager
}
