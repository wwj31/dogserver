package models

import (
	"server/service/game/iface"
)

type Model struct {
	Player iface.Player
}

func New(player iface.Player) Model {
	model := Model{
		Player: player,
	}
	return model
}

func (s *Model) OnLogin()  {}
func (s *Model) OnLogout() {}
