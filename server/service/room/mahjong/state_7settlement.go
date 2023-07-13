package mahjong

import (
	"server/common/log"
)

// 结算状态

type StateSettlement struct {
	*Mahjong
}

func (s *StateSettlement) State() int {
	return Settlement
}

func (s *StateSettlement) Enter() {
	s.gameCount++
	log.Infow("[Mahjong] enter state settlement",
		"room", s.room.RoomId, "dices", s.room.Dices, "master", s.masterIndex)
}

func (s *StateSettlement) Leave() {
	s.clear()
	log.Infow("[Mahjong] leave state settlement", "room", s.room.RoomId)
}

func (s *StateSettlement) Handle(shortId int64, v any) (result any) {
	return nil
}
