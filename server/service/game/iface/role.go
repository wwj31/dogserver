package iface

import (
	"time"
)

type Role interface {
	Modeler
	RoleId() string
	SetRoleId(v string)
	UId() string
	SetUId(v string)

	Name() string
	Icon() string
	CreateAt() time.Time
	LoginAt() time.Time
	LogoutAt() time.Time
}
