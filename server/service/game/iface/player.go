package iface

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"server/common"
)

type (
	Player interface {
		actor.Actor
		Session

		Send2Client(pb proto.Message)
		Login()
		Logout()
		IsNewRole() bool

		Gamer() Gamer
		Role() Role
		Item() Item
		Mail() Mailer
	}
)

type PlayerManager interface {
	SetPlayer(gSession common.GSession, id common.ActorId)
	PlayerBySession(gSession common.GSession) (common.ActorId, bool)
	GSessionByPlayer(id common.ActorId) (common.GSession, bool)
	DelGSession(gateSession common.GSession)
	RangeOnline(f func(gs common.GSession, player common.ActorId))

	Broadcast(msg proto.Message)
}
