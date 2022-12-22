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
	data inner.RoleInfo
}

func New(base models.Model) *Role {
	mod := &Role{Model: base}
	mod.data.RID = base.Player.RID()
	return mod
}

func (s *Role) Data() gogo.Message {
	return &s.data
}

func (s *Role) OnLogin(first bool) {
	nowStr := tools.TimeFormat(tools.Now())
	if first {
		//first
		s.SetAttribute(typ.Level, 1)
		s.SetAttribute(typ.Exp, 0)
		s.SetAttribute(typ.Gold, 0)
		s.data.CreateAt = nowStr
	}

	s.data.LoginAt = nowStr
	s.Player.Send2Client(s.roleInfoPush())
}

func (s *Role) OnLogout() {
	s.data.LogoutAt = tools.TimeFormat(tools.Now())
}

func (s *Role) roleInfoPush() *outer.RoleInfoPush {
	return &outer.RoleInfoPush{
		UID:     s.data.UID,
		RID:     s.data.RID,
		SId:     s.data.SId,
		Name:    s.data.Name,
		Icon:    s.data.Icon,
		Country: s.data.Country,
	}
}
