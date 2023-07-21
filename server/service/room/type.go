package room

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/logger"

	"server/proto/innermsg/inner"
)

type GamblingType int32

const (
	Mahjong GamblingType = 0
	DZZ     GamblingType = 1
)

func (g GamblingType) Int32() int32 {
	return int32(g)
}

type Gambling interface {
	Data(shortId int64) proto.Message
	SeatIndex(shortId int64) int
	CanEnter(p *inner.PlayerInfo) bool
	CanLeave(p *inner.PlayerInfo) bool
	CanReady(p *inner.PlayerInfo) bool
	PlayerEnter(player *Player)
	PlayerLeave(player *Player)
	Handle(shortId int64, v any) any
	Log() *logger.Logger
}
