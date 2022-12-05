package iface

import (
	gogo "github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
)

type Gamer interface {
	actor.Actor

	SID() uint16
	MsgToPlayer(rid string, sid uint16, msg gogo.Message)
}
