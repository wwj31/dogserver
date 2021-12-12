package role

import (
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/tools"
	"server/db/table"
	"server/proto/message"
	"server/service/game/logic/model"
)

type Role struct {
	model.Model

	tRole table.Role
}

func New(rid uint64, base model.Model) *Role {
	tRole := table.Role{RoleId: rid}
	err := base.Game().Load(&tRole)
	expect.Nil(err)

	role := &Role{
		Model: base,
		tRole: tRole,
	}
	return role
}

func (s *Role) OnLogin() {
	s.tRole.LoginAt = tools.Milliseconds()
	_ = s.Game().Send2Client(s.GateSession(), s.roleInfoPush())
	s.save()
}

func (s *Role) OnLogout() {
	s.tRole.LogoutAt = tools.Milliseconds()
	s.save()
}

func (s *Role) OnStop() {
	s.save()
}

func (s *Role) roleInfoPush() *message.RoleInfoPush {
	return &message.RoleInfoPush{
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
