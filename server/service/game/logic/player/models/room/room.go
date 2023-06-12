package room

import (
	"server/proto/outermsg/outer"
	"server/service/game/logic/player/models"
)

type Room struct {
	models.Model
}

func New(base models.Model) *Room {
	mod := &Room{Model: base}
	return mod
}

func (s *Room) OnLogin(first bool, enterGameRsp *outer.EnterGameRsp) {
	if first {
	}
}

func (s *Room) OnLogout() {

}
