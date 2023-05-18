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
	SetIcon(icon string)
	CreateAt() time.Time
	LoginAt() time.Time
	LogoutAt() time.Time
}
