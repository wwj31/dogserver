package role

import (
	"server/common"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player/models"
	"server/service/game/logic/player/models/role/typ"

	"github.com/gogo/protobuf/proto"

	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/tools"
)

type Role struct {
	models.Model
	role inner.RoleInfo
}

func New(base models.Model) *Role {
	mod := &Role{Model: base}

	if !base.Player.IsNewRole() {
		err := proto.Unmarshal(base.Player.PlayerData().RoleBytes, &mod.role)
		expect.Nil(err)
	} else {
		mod.SetAttribute(typ.Level, 1)
		mod.SetAttribute(typ.Exp, 0)
		mod.SetAttribute(typ.Glod, 0)
	}
	return mod
}

func (s *Role) OnSave() {
	s.Player.PlayerData().RoleBytes = common.ProtoMarshal(&s.role)
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
		UID:     s.role.UUId,
		RID:     s.role.RoleId,
		SId:     s.role.SId,
		Name:    s.role.Name,
		Icon:    s.role.Icon,
		Country: s.role.Country,
	}
}
