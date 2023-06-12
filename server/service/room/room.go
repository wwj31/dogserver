package room

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"

	"server/common/log"
	"server/common/router"
	"server/proto/innermsg/inner"
)

func New(RoomId int32, creatorShortId int64) *Room {
	return &Room{RoomId: RoomId, CreatorShortId: creatorShortId}
}

type Player struct {
	*inner.PlayerInfo
}
type Room struct {
	actor.Base
	currentMsg actor.Message

	RoomId         int32
	GameType       int32
	CreatorShortId int64

	Players []*Player
}

func (r *Room) OnInit() {
	router.Result(r, r.responseHandle)
	log.Debugf("Room:[%v] OnInit", r.RoomId)
}

func (r *Room) responseHandle(resultMsg any) {
	msg, ok := resultMsg.(proto.Message)
	if !ok {
		return
	}

	var err error
	if r.currentMsg.GetRequestId() != "" {
		err = r.Response(r.currentMsg.GetRequestId(), msg)
	} else {
		err = r.Send(r.currentMsg.GetSourceId(), msg)
	}

	if err != nil {
		log.Warnw("response to actor failed",
			"source", r.currentMsg.GetSourceId(),
			"msg name", r.currentMsg.GetMsgName())
	}
}

func (r *Room) FindPlayer(shortId int64) *Player {
	for _, v := range r.Players {
		if v.ShortId == shortId {
			return v
		}
	}
	return nil
}

func (r *Room) Info() *inner.RoomInfo {
	var players []*inner.PlayerInfo
	for _, v := range r.Players {
		players = append(players, v.PlayerInfo)
	}

	return &inner.RoomInfo{
		RoomId:   r.RoomId,
		GameType: r.GameType,
		Players:  players,
	}
}
