package role

import (
	gogo "github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/tools"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player/models"
	"server/service/game/logic/player/models/role/typ"
)

type Role struct {
	models.Model
	role *inner.RoleInfo
}

func New(base models.Model) *Role {
	mod := &Role{Model: base}
	return mod
}

func (s *Role) Data() gogo.Message {
	return s.role
}

func (s *Role) OnLogin(first bool) {
	if first {
		//first
		s.SetAttribute(typ.Level, 1)
		s.SetAttribute(typ.Exp, 0)
		s.SetAttribute(typ.Glod, 0)
	}

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
