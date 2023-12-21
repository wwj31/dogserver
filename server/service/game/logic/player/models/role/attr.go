package role

import (
	"time"

	"server/proto/outermsg/outer"
	"server/rdsop"

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
	s.data.Gold = s.data.Gold + v
	s.Player.SendToClient(&outer.UpdateGoldNtf{Gold: s.data.Gold})
	s.Player.UpdateInfoToRedis()
}

func (s *Role) GoldLine() int64       { return s.data.GoldLine }
func (s *Role) SetGoldLine(v int64)   { s.data.GoldLine = v }
func (s *Role) ForbidLogin() bool     { return s.data.ForbidLogin }
func (s *Role) SetForbidLogin(b bool) { s.data.ForbidLogin = b }

func (s *Role) Name() string        { return s.data.Name }
func (s *Role) Icon() string        { return s.data.Icon }
func (s *Role) Gender() int32       { return s.data.Gender }
func (s *Role) UpShortId() int64    { return rdsop.AgentUp(s.data.ShortId) }
func (s *Role) CreateAt() time.Time { return tools.TimeParse(s.data.CreateAt) }
func (s *Role) LoginAt() time.Time  { return tools.TimeParse(s.data.LoginAt) }
func (s *Role) LogoutAt() time.Time { return tools.TimeParse(s.data.LogoutAt) }
func (s *Role) SetBaseInfo(icon, name string, gender int32) {
	s.data.Icon = icon
	s.data.Name = name
	s.data.Gender = gender
}
