package role

import (
	"time"

	"server/common"
	"server/proto/outermsg/outer"

	"github.com/wwj31/dogactor/tools"
)

func (s *Role) RoleId() string { return s.data.RID }

func (s *Role) SetUID(v string) { s.data.UID = v }
func (s *Role) UID() string     { return s.data.UID }

func (s *Role) SetPhone(v string) { s.data.Phone = v }
func (s *Role) Phone() string     { return s.data.Phone }

func (s *Role) SetShortId(v int64) { s.data.ShortId = v }
func (s *Role) ShortId() int64     { return s.data.ShortId }

func (s *Role) Gold() int64 { return s.data.Gold }
func (s *Role) AddGold(v int64) {
	s.data.Gold = common.Max(s.data.Gold+v, 0)
	s.Player.SendToClient(&outer.UpdateGoldNtf{Gold: s.data.Gold})
}

func (s *Role) Name() string        { return s.data.Name }
func (s *Role) Icon() string        { return s.data.Icon }
func (s *Role) Gender() int32       { return s.data.Gender }
func (s *Role) UpShortId() int64    { return s.upShortId }
func (s *Role) CreateAt() time.Time { return tools.TimeParse(s.data.CreateAt) }
func (s *Role) LoginAt() time.Time  { return tools.TimeParse(s.data.LoginAt) }
func (s *Role) LogoutAt() time.Time { return tools.TimeParse(s.data.LogoutAt) }
func (s *Role) SetBaseInfo(icon, name string, gender int32) {
	s.data.Icon = icon
	s.data.Name = name
	s.data.Gender = gender
}
