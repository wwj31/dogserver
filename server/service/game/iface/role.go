package iface

import (
	"server/service/game/logic/player/models/role/typ"
)

type Role interface {
	Modeler
	RoleId() string
	SetRoleId(v string)
	UId() string
	SetUId(v string)

	SId() uint64
	Name() string
	Icon() string
	Country() string
	CreateAt() int64
	LoginAt() int64
	LogoutAt() int64
	Attribute(typ typ.Attribute) int64
	SetAttribute(typ typ.Attribute, val int64)
}
