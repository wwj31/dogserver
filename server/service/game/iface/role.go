package iface

import "server/service/game/logic/player/role/typ"

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
	LoginAt() int64
	LogoutAt() int64
	GetAttribute(typ typ.Attribute) int64
}
