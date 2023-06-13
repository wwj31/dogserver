package room

import (
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"

	"server/common"
	"server/common/log"
	"server/common/router"
	"server/proto/convert"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
)

var gameMaxPlayers = map[int32]int{
	0: 4,
	1: 3,
}

func New(roomId, gameType int32, creator *inner.PlayerInfo) *Room {
	r := &Room{
		RoomId:         roomId,
		GameType:       gameType,
		CreatorShortId: creator.ShortId,
	}
	//r.AddPlayer(creator)
	return r
}

type (
	Player struct {
		*inner.PlayerInfo
	}

	Room struct {
		actor.Base
		currentMsg actor.Message
		stopping   bool

		RoomId         int32
		GameType       int32
		CreatorShortId int64

		Players []*Player
	}
)

func (r *Room) OnInit() {
	router.Result(r, r.responseHandle)
	log.Debugf("Room:[%v] OnInit", r.RoomId)
}

func (r *Room) OnHandle(msg actor.Message) {
	pt, ok := msg.Payload().(proto.Message)
	if !ok {
		log.Warnw("room handler msg is not proto",
			"msg", reflect.TypeOf(msg.Payload()).String())
		return
	}

	if r.stopping {
		log.Warnw("room is stopping not handle msg", "roomId", r.RoomId)
		return
	}

	r.currentMsg = msg
	router.Dispatch(r, pt)
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

func (r *Room) IsFull() bool { return len(r.Players) >= gameMaxPlayers[r.GameType] }

func (r *Room) AddPlayer(playerInfo *inner.PlayerInfo) *inner.Error {
	if r.FindPlayer(playerInfo.ShortId) != nil {
		return &inner.Error{ErrorCode: int32(outer.ERROR_PLAYER_ALREADY_IN_ROOM)}
	}

	for _, p := range r.Players {
		gSession := common.GSession(p.GSession)
		if gSession.Invalid() {
			continue
		}
		gSession.SendToClient(r, &outer.RoomPlayerEnterNtf{Player: convert.PlayerInnerToOuter(playerInfo)})
	}
	r.Players = append(r.Players, &Player{PlayerInfo: playerInfo})
	log.Infow("room add player", "roomId", r.RoomId, "player", playerInfo.ShortId)
	return nil
}

func (r *Room) Stop() {
	r.stopping = true
}

func (r *Room) DelPlayer(shortId int64) {
	for i, player := range r.Players {
		if player.ShortId == shortId {
			r.Players = append(r.Players[:i], r.Players[i+1:]...)
			return
		}
	}
}

func (r *Room) Info() *inner.RoomInfo {
	var players []*inner.PlayerInfo
	for _, v := range r.Players {
		players = append(players, v.PlayerInfo)
	}

	return &inner.RoomInfo{
		RoomId:         r.RoomId,
		GameType:       r.GameType,
		CreatorShortId: r.CreatorShortId,
		Players:        players,
	}
}
