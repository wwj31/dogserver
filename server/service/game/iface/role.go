package iface

import (
	"time"
)

type Role interface {
	Modeler
	RoleId() string
	SetUID(v string)
	UID() string

	Phone() string
	SetPhone(v string)

	SetShortId(v int64)
	ShortId() int64

	Gold() int64
	AddGold(v int64)

	Name() string
	Icon() string
	Gender() int32
	UpShortId() int64
	SetBaseInfo(icon, name string, gender int32)
	CreateAt() time.Time
	LoginAt() time.Time
	LogoutAt() time.Time
}
