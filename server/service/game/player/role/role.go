package role

import (
	"github.com/wwj31/dogactor/expect"
	"server/db/table"
	"server/proto/message"
	"server/service/game/player/model"
)

type Role struct {
	model.Model

	table.Role
}

func New(rid uint64, base model.Model) *Role {
	t_role := table.Role{RoleId: rid}
	role := &Role{
		Model: base,
		Role:  t_role,
	}
	err := role.Game().Load(&t_role)
	expect.Nil(err)

	return role
}

func (s *Role) OnLogin() {
	_ = s.Game().Send2Client(s.GateSession(), s.roleInfoPush())
}

func (s *Role) OnLogout() {
	s.Log().Info("Logout Role")
}

func (s *Role) Table() table.Role { return s.Role }

func (s *Role) roleInfoPush() *message.RoleInfoPush {
	return &message.RoleInfoPush{
		UID:     s.UUId,
		RID:     s.RoleId,
		SId:     s.SId,
		Name:    s.Name,
		Icon:    s.Icon,
		Country: s.Country,
	}
}
