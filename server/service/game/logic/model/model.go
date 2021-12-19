package model

import (
	"github.com/wwj31/dogactor/log"
	"server/db/table"
	"server/service/game/iface"
)

type Model struct {
	Player iface.Player
	logger log.Logger
	tab    table.Tabler
}

func New(player iface.Player, loglevel ...int32) Model {
	lv := int32(log.TAG_INFO_I)
	if len(loglevel) == 1 {
		lv = loglevel[0]
	}

	model := Model{
		Player: player,
		logger: log.New(lv),
	}
	return model
}

func (s *Model) OnLogin()            {}
func (s *Model) OnLogout()           {}
func (s *Model) Table() table.Tabler { return s.tab }

func (s *Model) SetTable(t table.Tabler) { s.tab = t }

func (s *Model) Log() *log.Logger {
	return &s.logger
}
