package role

import "github.com/wwj31/dogactor/tools"

func (s *Role) AddExp() {
	s.tRole.LoginAt = tools.Milliseconds()
	s.Player.Send2Client(s.roleInfoPush())
	s.save()
}
