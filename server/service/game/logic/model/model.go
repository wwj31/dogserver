package model

import (
	"server/db/table"
	"server/service/game/iface"
)

type Model struct {
	Player iface.Player
	tab    table.Tabler
}

func New(player iface.Player) Model {

	model := Model{
		Player: player,
	}
	return model
}

func (s *Model) OnLogin()            {}
func (s *Model) OnLogout()           {}
func (s *Model) Table() table.Tabler { return s.tab }

func (s *Model) SetTable(t table.Tabler) { s.tab = t }
