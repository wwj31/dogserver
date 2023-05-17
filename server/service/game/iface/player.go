package iface

import (
	"github.com/gogo/protobuf/proto"
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
		ShortId() int64

		Observer() *common.Observer
		Send2Client(pb proto.Message)
		Online() bool

		Account() *inner.Account
		Gamer() Gamer
		Role() Role
		Mail() Mailer
	}
)
