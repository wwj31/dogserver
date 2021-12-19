package mail

import (
	"server/service/game/logic/model"
)

type Mail struct {
	model.Model
}

func (s *Mail) OnLogin() {
	s.Player.Send2Client(nil)
}
