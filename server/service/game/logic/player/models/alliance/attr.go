package alliance

func (s *Alliance) AllianceId() int32 {
	return s.data.AllianceId
}
func (s *Alliance) SetAllianceId(id int32) {
	s.data.AllianceId = id
	s.Player.UpdateInfoToRedis()
}

func (s *Alliance) Position() int32 {
	return s.data.Position
}
func (s *Alliance) SetPosition(p int32) {
	s.data.Position = p
	s.Player.UpdateInfoToRedis()
}
