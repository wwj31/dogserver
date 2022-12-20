package role

import (
	"github.com/wwj31/dogactor/tools"
	"server/common"
	"server/service/game/logic/player/event"
	"server/service/game/logic/player/models/role/typ"
	"time"
)

func (s *Role) RoleId() string      { return s.data.RID }
func (s *Role) SetRoleId(v string)  { s.data.RID = v }
func (s *Role) UId() string         { return s.data.UID }
func (s *Role) SetUId(v string)     { s.data.UID = v }
func (s *Role) SId() uint64         { return s.data.SId }
func (s *Role) Name() string        { return s.data.Name }
func (s *Role) Icon() string        { return s.data.Icon }
func (s *Role) Country() string     { return s.data.Country }
func (s *Role) CreateAt() time.Time { return tools.TimeParse(s.data.CreateAt) }
func (s *Role) LoginAt() time.Time  { return tools.TimeParse(s.data.LoginAt) }
func (s *Role) LogoutAt() time.Time { return tools.TimeParse(s.data.LogoutAt) }

func newAtrributeMap() map[int64]int64 {
	return make(map[int64]int64, typ.AttributeMax)
}

func (s *Role) Attribute(typ typ.Attribute) int64 {
	if s.data.Attributes == nil {
		s.data.Attributes = newAtrributeMap()
	}
	return s.data.Attributes[int64(typ)]
}

func (s *Role) SetAttribute(t typ.Attribute, val int64) {
	if s.data.Attributes == nil {
		s.data.Attributes = newAtrributeMap()
	}
	old := s.data.Attributes[t.Int64()]
	s.data.Attributes[t.Int64()] = val

	common.EmitEvent(s.Player.Observer(), event.ChangeAttribute{
		Type:   t,
		OldVal: old,
		NewVal: val,
	})
}
