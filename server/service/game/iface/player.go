package iface

import "server/common"

type Player interface {
	GateSession() common.GSession
	SetGateSession(gSession common.GSession)
}

type Manager interface {
	PlayerBySession(gateSession common.GSession) (Player, bool)
	PlayerByUID(uid uint64) (Player, bool)
	PlayerByRID(rid uint64) (Player, bool)
}
