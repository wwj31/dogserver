package role

import (
	"math/rand"
	"time"

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

func randName() string {
	name1 := "name1"
	name2 := "name2"
	r1 := rand.Int31n(int32(conf.LenRandName())) + 1
	r2 := rand.Int31n(int32(conf.LenRandName())) + 1
	name1Conf := conf.GetRandName(int64(r1))
	name2Conf := conf.GetRandName(int64(r2))
	if name1Conf != nil {
		name1 = name1Conf.Name1()
	}
	if name2Conf != nil {
		name2 = name2Conf.Name2()
	}
	return name1 + name2
}

func (s *Role) OnLogin(first bool, enterGameRsp *outer.EnterGameRsp) {
	nowStr := tools.TimeFormat(tools.Now())
	if first {
		s.data.CreateAt = nowStr
		s.data.Phone = s.Player.Role().Phone()
		s.data.Name = randName()
		s.data.Icon = cast.ToString(rand.Int31n(10) + 1)
		s.data.LogoutAt = tools.TimeFormat(tools.Now().Add(-time.Second))
	}

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
