package role

import "server/db/table"

func (s *Role) RoleId() uint64  { return s.tRole.RoleId }
func (s *Role) UUId() uint64    { return s.tRole.UUId }
func (s *Role) SId() uint64     { return s.tRole.SId }
func (s *Role) Name() string    { return s.tRole.Name }
func (s *Role) Icon() string    { return s.tRole.Icon }
func (s *Role) Country() string { return s.tRole.Country }
func (s *Role) IsDelete() bool  { return s.tRole.IsDelete }
func (s *Role) CreateAt() int64 { return s.tRole.CreateAt }
func (s *Role) LoginAt() int64  { return s.tRole.LoginAt }
func (s *Role) LogoutAt() int64 { return s.tRole.LogoutAt }

func (s *Role) GetAttr(typ int64) int64 {
	if s.tRole.Attributes == nil {
		s.tRole.Attributes = table.AttributeMap{}
	}
	return s.tRole.Attributes[typ]
}
