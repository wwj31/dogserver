package room

import (
	"reflect"
	"server/common/actortype"
	"server/rdsop"

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

func New(info *rdsop.NewRoomInfo) *Room {
	r := &Room{
		RoomId:         info.RoomId,
		GameType:       info.GameType,
		GameParams:     info.Params,
		CreatorShortId: info.CreatorShortId,
		AllianceId:     info.AllianceId,
	}
	return r
}

type (
	Player struct {
		*inner.PlayerInfo
		Ready bool
	}

	Room struct {
		actor.Base
		currentMsg actor.Message
		stopping   bool

		RoomId         int64
		GameType       int32             // 游戏类型
		GameParams     *outer.GameParams // 游戏参数
		CreatorShortId int64             // 房间创建者
		AllianceId     int32             // 归属联盟

		Players []*Player

		gambling Gambling
	}
)

func (r *Room) InjectGambling(gambling Gambling) {
	r.gambling = gambling
}

func (r *Room) OnInit() {
	router.Result(r, r.responseHandle)
	log.Debugf("Room:[%v] OnInit", r.RoomId)
}

func (r *Room) OnStop() bool {
	log.Debugw("room stop", "roomId", r.RoomId)
	return true
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
		var errCode outer.ERROR
		errCode, ok = resultMsg.(outer.ERROR)
		if !ok {
			return
		}
		msg = &inner.Error{ErrorCode: int32(errCode)}
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

func (r *Room) IsFull() bool { return len(r.Players) >= gameMaxPlayers[r.GameType] }
func (r *Room) Disband()     { r.stopping = true }

func (r *Room) CanEnter() bool {
	if r.stopping {
		return false
	}

	// TODO 游戏状态中,不能进入
	return true
}

func (r *Room) CanLeave() bool {
	// TODO 游戏状态中,不能离开
	return true
}

func (r *Room) CanReady() bool {
	// TODO 游戏状态中,不能离开
	return true
}

func (r *Room) FindPlayer(shortId int64) *Player {
	for _, v := range r.Players {
		if v.ShortId == shortId {
			return v
		}
	}
	return nil
}

func (r *Room) PlayerEnter(playerInfo *inner.PlayerInfo) *inner.Error {
	if r.FindPlayer(playerInfo.ShortId) != nil {
		return &inner.Error{ErrorCode: int32(outer.ERROR_PLAYER_ALREADY_IN_ROOM)}
	}

	newPlayer := &Player{
		PlayerInfo: playerInfo,
		Ready:      false,
	}
	r.Players = append(r.Players, newPlayer)

	r.Broadcast(&outer.RoomPlayerEnterNtf{Player: &outer.RoomPlayerInfo{
		BaseInfo: convert.PlayerInnerToOuter(playerInfo),
		Ready:    false,
	}})
	r.gambling.PlayerEnter(newPlayer)
	log.Infow("room add player", "roomId", r.RoomId, "player", playerInfo.ShortId)
	return nil
}

func (r *Room) PlayerLeave(shortId int64) {
	var ntf bool
	for i, player := range r.Players {
		if player.ShortId == shortId {
			r.gambling.PlayerLeave(player)
			r.Players = append(r.Players[:i], r.Players[i+1:]...)
			ntf = true
			log.Infow("room del player", "roomId", r.RoomId, "shortId", shortId)
			break
		}
	}

	if ntf {
		r.Broadcast(&outer.RoomPlayerLeaveNtf{ShortId: shortId})
	}
}

func (r *Room) PlayerReady(shortId int64, ready bool) (ok bool, err outer.ERROR) {
	p := r.FindPlayer(shortId)
	if p == nil {
		// 玩家不在房间内
		log.Warnw("leave the room cannot find player", "roomId", r.RoomId, "msg", shortId)
		return false, outer.ERROR_PLAYER_NOT_IN_ROOM
	}
	if p.Ready == ready {
		return true, 0
	}

	p.Ready = ready
	r.gambling.PlayerReady(p)
	r.Broadcast(&outer.RoomPlayerReadyNtf{
		ShortId: p.ShortId,
		Ready:   ready,
	})
	return true, 0
}

func (r *Room) SendToPlayer(shortId int64, msg proto.Message) {
	wrapper := &inner.GamblingMsgToClientWrapper{
		MsgType: common.ProtoType(msg),
		Data:    common.ProtoMarshal(msg),
	}

	player := r.FindPlayer(shortId)
	if player == nil {
		log.Errorw("cannot find player", "roomId", r.RoomId, "shortId", shortId, "msg", common.ProtoType(msg))
		return
	}

	playerActor := actortype.PlayerId(player.RID)
	if err := r.Send(playerActor, wrapper); err != nil {
		log.Errorw("send msg to player failed", "roomId", r.RoomId, "shortId", shortId, "player actor", playerActor)
		return
	}
}

func (r *Room) Broadcast(msg proto.Message, ignores ...int64) {
	ignoreMap := make(map[int64]struct{})
	for _, ig := range ignores {
		ignoreMap[ig] = struct{}{}
	}

	for _, p := range r.Players {
		if _, ignore := ignoreMap[p.ShortId]; ignore {
			continue
		}

		r.SendToPlayer(p.ShortId, msg)
	}
}

func (r *Room) Info() *inner.RoomInfo {
	var players []*inner.RoomPlayerInfo
	for _, player := range r.Players {
		players = append(players, player.InnerPB())
	}

	gameParamsBytes, _ := proto.Marshal(r.GameParams)
	return &inner.RoomInfo{
		RoomId:         r.RoomId,
		GameType:       r.GameType,
		GameParams:     gameParamsBytes,
		CreatorShortId: r.CreatorShortId,
		Players:        players,
	}
}

func (p *Player) InnerPB() *inner.RoomPlayerInfo {
	return &inner.RoomPlayerInfo{
		BaseInfo: p.PlayerInfo,
		Ready:    p.Ready,
	}
}
