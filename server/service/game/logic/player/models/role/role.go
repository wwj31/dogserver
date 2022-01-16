package role

import (
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/tools"
	"server/db/table"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player/models"
	"server/service/game/logic/player/models/role/typ"
)

type Role struct {
	models.Model

	tRole table.Role
}

func New(rid uint64, base models.Model) *Role {
	role := &Role{
		Model: base,
		tRole: table.Role{
			RoleId:     rid,
			Attributes: make(table.AttributeMap),
		},
	}
	err := base.Player.Gamer().Load(&role.tRole)
	expect.Nil(err)

	if role.IsNewRole() {
		role.SetAttribute(typ.Level, 1)
		role.SetAttribute(typ.Exp, 0)
		role.SetAttribute(typ.Glod, 0)
	}
	return role
}

func (s *Role) OnLogin() {
	s.tRole.LoginAt = tools.Milliseconds()
	s.Player.Send2Client(s.roleInfoPush())
	s.save()
}

func (s *Role) OnLogout() {
	s.tRole.LogoutAt = tools.Milliseconds()
	s.save()
}

func (s *Role) roleInfoPush() *outer.RoleInfoPush {
	return &outer.RoleInfoPush{
		UID:     s.tRole.UUId,
		RID:     s.tRole.RoleId,
		SId:     s.tRole.SId,
		Name:    s.tRole.Name,
		Icon:    s.tRole.Icon,
		Country: s.tRole.Country,
	}
}

func (s *Role) save() {
	s.SetTable(&s.tRole)
}
