package role

import (
	"github.com/wwj31/dogactor/expect"
	"server/db/table"
	"server/service/game/iface"
)

type Role struct {
	table.Role
}

func New(rid uint64, loader iface.Loader) *Role {
	t := table.Role{RoleId: rid}
	err := loader.Load(&t)
	expect.Nil(err)

	return &Role{Role: t}
}

func (s *Role) Table() table.Role { return s.Role }
