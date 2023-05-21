package iface

import (
	"time"
)

type Role interface {
	Modeler
	RoleId() string
	Phone() string
	SetPhone(v string)

	Name() string
	Icon() string
	Gender() int32
	SetBaseInfo(icon, name string, gender int32)
	CreateAt() time.Time
	LoginAt() time.Time
	LogoutAt() time.Time
}
