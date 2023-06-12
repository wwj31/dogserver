package iface

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"

	"server/common"
	"server/proto/innermsg/inner"
)

type (
	Player interface {
		actor.Actor
		actor.Timer
		actor.Messenger

		Session
		RID() string

		Observer() *common.Observer
		SendToClient(pb proto.Message)
		Online() bool
		PlayerInfo() *inner.PlayerInfo
		UpdateInfoToRedis()

		Gamer() Gamer
		Role() Role
		Mail() Mailer
		Alliance() Alliance
		Agent() Agent
		Room() Room
	}
)
