package role

import (
	"math/rand"
	"server/rdsop"

	"github.com/spf13/cast"

	gogo "github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/tools"

	"server/config/conf"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player/models"
)

type Role struct {
	models.Model
	data      inner.RoleInfo
	upShortId int64
}

func New(base models.Model) *Role {
	mod := &Role{Model: base}
	mod.data.RID = base.Player.RID()
	return mod
}

func (s *Role) Data() gogo.Message {
	return &s.data
}

func randName() string {
	r1 := rand.Int31n(int32(conf.LenRandName()))
	r2 := rand.Int31n(int32(conf.LenRandName()))
	name1 := conf.GetRandName(int64(r1)).Name1()
	name2 := conf.GetRandName(int64(r2)).Name2()
	return name1 + name2
}

func (s *Role) OnLogin(first bool, enterGameRsp *outer.EnterGameRsp) {
	nowStr := tools.TimeFormat(tools.Now())
	if first {
		//first
		s.data.CreateAt = nowStr
		s.data.Phone = s.Player.Role().Phone()
		s.data.Name = randName()
		s.data.Icon = cast.ToString(rand.Int31n(10) + 1)
		s.data.LogoutAt = nowStr
	}

	s.upShortId = rdsop.AgentUp(s.Player.Role().ShortId())

	s.data.LoginAt = nowStr
	enterGameRsp.RoleInfo = s.roleInfo()
}

func (s *Role) OnLogout() {
	s.data.LogoutAt = tools.TimeFormat(tools.Now())
}

func (s *Role) roleInfo() *outer.RoleInfo {
	return &outer.RoleInfo{
		RID:     s.data.RID,
		ShortId: s.data.ShortId,
		Phone:   s.data.Phone,
		Name:    s.data.Name,
		Icon:    s.data.Icon,
		Gold:    s.data.Gold,
	}
}
