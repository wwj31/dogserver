package iface

import (
	"server/service/game/logic/player/models/role/typ"
)

type Role interface {
	Modeler
	RoleId() uint64
	UUId() uint64
	SId() uint64
	Name() string
	Icon() string
	Country() string
	IsDelete() bool
	CreateAt() int64
	IsNewRole() bool
	LoginAt() int64
	LogoutAt() int64
	Attribute(typ typ.Attribute) int64
	SetAttribute(typ typ.Attribute, val int64)
}
