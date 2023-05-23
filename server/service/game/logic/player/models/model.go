package models

import (
	gogo "github.com/gogo/protobuf/proto"

	"server/proto/outermsg/outer"
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

func (s *Model) Data() gogo.Message                                   { return nil }
func (s *Model) OnLoaded()                                            {}
func (s *Model) OnLogin(first bool, enterGameRsp *outer.EnterGameRsp) {}
func (s *Model) OnLogout()                                            {}
