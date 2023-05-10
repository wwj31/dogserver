package role

import (
	gogo "github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/tools"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player/models"
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

func (s *Role) OnLogin(first bool, enterGameRsp *outer.EnterGameRsp) {
	nowStr := tools.TimeFormat(tools.Now())
	if first {
		//first
		s.data.CreateAt = nowStr
	}

	s.data.LoginAt = nowStr
	enterGameRsp.RoleInfo = s.roleInfoPush()
}

func (s *Role) OnLogout() {
	s.data.LogoutAt = tools.TimeFormat(tools.Now())
}

func (s *Role) roleInfoPush() *outer.RoleInfo {
	return &outer.RoleInfo{
		RID:     s.data.RID,
		ShortID: s.data.ShortID,
		Phone:   s.data.Phone,
		Name:    s.data.Name,
		Icon:    s.data.Icon,
		Gold:    s.data.Gold,
	}
}
