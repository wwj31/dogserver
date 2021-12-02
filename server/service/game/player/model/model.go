package model

import (
	"github.com/wwj31/dogactor/log"
	"server/service/game/iface"
)

type Model struct {
	iface.Player
	logger log.Logger
}

func New(player iface.Player, loglevel ...int32) Model {
	lv := int32(log.TAG_INFO_I)
	if len(loglevel) == 1 {
		lv = loglevel[0]
	}

	model := Model{
		Player: player,
		logger: *log.New(lv),
	}
	return model
}

func (s *Model) OnLogin()  {}
func (s *Model) OnLogout() {}

func (s *Model) Log() *log.Logger {
	return &s.logger
}
