package role

import (
	"github.com/wwj31/dogactor/tools"
	"time"
)

func (s *Role) RoleId() string      { return s.data.RID }
func (s *Role) SetPhone(v string)   { s.data.Phone = v }
func (s *Role) Phone() string       { return s.data.Phone }
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
