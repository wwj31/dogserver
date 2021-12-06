package iface

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
}
