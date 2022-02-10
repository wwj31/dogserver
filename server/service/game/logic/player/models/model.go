package models

import (
	"server/db/table"
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

func (s *Model) OnSave(data *table.Player) {}
func (s *Model) OnLogin()                  {}
func (s *Model) OnLogout()                 {}
