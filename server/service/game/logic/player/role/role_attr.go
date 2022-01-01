package role

import (
	"server/db/table"
	"server/service/game/logic/player/role/typ"
)

func (s *Role) RoleId() uint64  { return s.tRole.RoleId }
func (s *Role) UUId() uint64    { return s.tRole.UUId }
func (s *Role) SId() uint64     { return s.tRole.SId }
func (s *Role) Name() string    { return s.tRole.Name }
func (s *Role) Icon() string    { return s.tRole.Icon }
func (s *Role) Country() string { return s.tRole.Country }
func (s *Role) IsDelete() bool  { return s.tRole.IsDelete }
func (s *Role) CreateAt() int64 { return s.tRole.CreateAt }
func (s *Role) LoginAt() int64  { return s.tRole.LoginAt }
func (s *Role) IsNewRole() bool { return s.tRole.LoginAt == 0 }
func (s *Role) LogoutAt() int64 { return s.tRole.LogoutAt }

func (s *Role) Attribute(typ typ.Attribute) int64 {
	if s.tRole.Attributes == nil {
		s.tRole.Attributes = table.AttributeMap{}
	}
	return s.tRole.Attributes[int64(typ)]
}

func (s *Role) SetAttribute(typ typ.Attribute, val int64) {
	if s.tRole.Attributes == nil {
		s.tRole.Attributes = table.AttributeMap{}
	}
	s.tRole.Attributes[int64(typ)] = val
}
