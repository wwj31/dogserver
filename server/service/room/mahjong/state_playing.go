package mahjong

import (
	"server/common/log"
)

// 游戏状态

type StatePlaying struct {
	*Mahjong
}

func (s *StatePlaying) State() int {
	return Playing
}

func (s *StatePlaying) Enter() {
	log.Infow("[Mahjong] enter state playing", "room", s.room.RoomId)
}

func (s *StatePlaying) Leave() {
	log.Infow("[Mahjong] leave state playing", "room", s.room.RoomId)
}

func (s *StatePlaying) Handle(shortId int64, v any) (result any) {
	return nil
}
