package iface

import (
	"github.com/golang/protobuf/proto"
	"server/common"
)

type Player interface {
	Session
	Game() Gamer

	Login()
	Logout()
	Stop()

	IsNewRole() bool
	Role() Role
	Item() Item
}

type PlayerManager interface {
	SetPlayer(p Player)
	PlayerBySession(gSession common.GSession) (Player, bool)
	PlayerByUID(uid uint64) (Player, bool)
	PlayerByRID(rid uint64) (Player, bool)
	OfflinePlayer(gSession common.GSession)
	RangeOnline(f func(player Player), except ...uint64)

	Broadcast(msg proto.Message, except ...uint64)

	Stop()
}
