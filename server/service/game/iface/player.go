package iface

import (
	"github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
)

type (
	Player interface {
		actor.Actor
		Session
		RID() string

		Send2Client(pb proto.Message)
		Login(first bool)
		Logout()
		Online() bool

		Gamer() Gamer
		Role() Role
		Item() Item
		Mail() Mailer
		Chat() Chat
	}
)
