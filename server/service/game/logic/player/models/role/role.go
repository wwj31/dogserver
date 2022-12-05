package role

import (
	gogo "github.com/gogo/protobuf/proto"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player/models"
	"server/service/game/logic/player/models/role/typ"

	"github.com/wwj31/dogactor/tools"
)

type Role struct {
	models.Model
	role *inner.RoleInfo
}

func New(base models.Model) *Role {
	mod := &Role{Model: base}

	//first
	mod.SetAttribute(typ.Level, 1)
	mod.SetAttribute(typ.Exp, 0)
	mod.SetAttribute(typ.Glod, 0)
	return mod
}

func (s *Role) OnSave() gogo.Message {
	return s.role
}

func (s *Role) OnLogin() {
	s.role.LoginAt = tools.Milliseconds()
	s.Player.Send2Client(s.roleInfoPush())
}

func (s *Role) OnLogout() {
	s.role.LogoutAt = tools.Milliseconds()
}

func (s *Role) roleInfoPush() *outer.RoleInfoPush {
	return &outer.RoleInfoPush{
		UID:     s.role.UID,
		RID:     s.role.RID,
		SId:     s.role.SId,
		Name:    s.role.Name,
		Icon:    s.role.Icon,
		Country: s.role.Country,
	}
}
