package room

type GamblingType int32

const (
	Mahjong GamblingType = 0
	DZZ     GamblingType = 1
)

func (g GamblingType) Int32() int32 {
	return int32(g)
}

type Gambling interface {
	PlayerEnter(player *Player)
	PlayerLeave(player *Player)
	PlayerReady(player *Player)
}
