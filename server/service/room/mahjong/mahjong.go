package mahjong

import "server/service/room"

func New(r *room.Room) *mahjong {
	return &mahjong{Room: r}
}

type mahjong struct {
	*room.Room
}

func (m mahjong) PlayerEnter(player *room.Player) {
	//TODO implement me
	panic("implement me")
}

func (m mahjong) PlayerLeave(player *room.Player) {
	//TODO implement me
	panic("implement me")
}

func (m mahjong) PlayerReady(player *room.Player) {
	//TODO implement me
	panic("implement me")
}
