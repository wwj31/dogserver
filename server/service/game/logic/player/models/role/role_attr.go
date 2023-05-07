package role

import (
	"github.com/wwj31/dogactor/tools"
	"time"
)

func (s *Role) RoleId() string      { return s.data.RID }
func (s *Role) SetRoleId(v string)  { s.data.RID = v }
func (s *Role) UId() string         { return s.data.UID }
func (s *Role) SetUId(v string)     { s.data.UID = v }
func (s *Role) Name() string        { return s.data.Name }
func (s *Role) Icon() string        { return s.data.Icon }
func (s *Role) CreateAt() time.Time { return tools.TimeParse(s.data.CreateAt) }
func (s *Role) LoginAt() time.Time  { return tools.TimeParse(s.data.LoginAt) }
func (s *Role) LogoutAt() time.Time { return tools.TimeParse(s.data.LogoutAt) }
