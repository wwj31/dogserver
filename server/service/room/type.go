package room

import "server/proto/innermsg/inner"

type GamblingType int32

const (
	Mahjong GamblingType = 0
	DZZ     GamblingType = 1
)

func (g GamblingType) Int32() int32 {
	return int32(g)
}

type Gambling interface {
	SeatIndex(shortId int64) int
	CanEnter(p *inner.PlayerInfo) bool
	CanLeave(p *inner.PlayerInfo) bool
	CanReady(p *inner.PlayerInfo) bool
	PlayerEnter(player *Player)
	PlayerLeave(player *Player)
	PlayerReady(player *Player)
	Handle(v any, shortId int64) any
}
