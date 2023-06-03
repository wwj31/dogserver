package iface

import (
	"server/common"
)

type Session interface {
	GateSession() common.GSession
	SetGateSession(gSession common.GSession)
}
