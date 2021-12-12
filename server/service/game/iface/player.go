package iface

import "server/common"

type Player interface {
	Session
	Game() Gamer

	Login()
	Logout()
	Stop()

	Role() Role
}

type PlayerManager interface {
	SetPlayer(p Player)
	PlayerBySession(gSession common.GSession) (Player, bool)
	PlayerByUID(uid uint64) (Player, bool)
	PlayerByRID(rid uint64) (Player, bool)
	RangeOnline(f func(player Player), except ...uint64)
	Stop()
}
