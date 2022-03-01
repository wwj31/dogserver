package role

import (
	"server/service/game/logic/player/models/role/typ"
)

func (s *Role) RoleId() uint64  { return s.role.RoleId }
func (s *Role) UUId() uint64    { return s.role.UUId }
func (s *Role) SId() uint64     { return s.role.SId }
func (s *Role) Name() string    { return s.role.Name }
func (s *Role) Icon() string    { return s.role.Icon }
func (s *Role) Country() string { return s.role.Country }
func (s *Role) CreateAt() int64 { return s.role.CreateAt }
func (s *Role) LoginAt() int64  { return s.role.LoginAt }
func (s *Role) LogoutAt() int64 { return s.role.LogoutAt }

func newAtrributeMap() map[int64]int64 {
	return make(map[int64]int64, typ.AttributeMax)
}

func (s *Role) Attribute(typ typ.Attribute) int64 {
	if s.role.Attributes == nil {
		s.role.Attributes = newAtrributeMap()
	}
	return s.role.Attributes[int64(typ)]
}

func (s *Role) SetAttribute(typ typ.Attribute, val int64) {
	if s.role.Attributes == nil {
		s.role.Attributes = newAtrributeMap()
	}
	s.role.Attributes[int64(typ)] = val
}
