package models

import (
	gogo "github.com/gogo/protobuf/proto"
	"server/service/game/iface"
)

type Model struct {
	iface.Player
}

func New(player iface.Player) Model {
	model := Model{
		Player: player,
	}
	return model
}

func (s *Model) Data() gogo.Message { return nil }
func (s *Model) OnLoaded()          {}
func (s *Model) OnLogin(first bool) {}
func (s *Model) OnLogout()          {}
