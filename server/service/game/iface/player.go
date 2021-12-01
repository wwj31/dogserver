package iface

import "server/common"

type Player interface {
	Session
	Role
}

type PlayerManager interface {
	PlayerBySession(gateSession common.GSession) (Player, bool)
	PlayerByUID(uid uint64) (Player, bool)
	PlayerByRID(rid uint64) (Player, bool)
}
