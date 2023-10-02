package room

import (
	"fmt"
	"reflect"
	"time"

	"github.com/wwj31/dogactor/logger"

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
		log: logger.New(logger.Option{
			Level:          logger.DebugLevel,
			LogPath:        "./log/room",
			FileName:       fmt.Sprintf("room%v", info.RoomId),
			FileMaxAge:     30,
			FileMaxSize:    1024,
			FileMaxBackups: 10,
			DisplayConsole: true,
			Skip:           3,
		}),
	}
	return r
}

type (
	Player struct {
		*inner.PlayerInfo
		EnterAt time.Time
	}

	Room struct {
		actor.Base
		stopping       bool
		CurrentMsg     actor.Message
		RoomId         int64
		GameType       GamblingType      // 游戏类型
		GameParams     *outer.GameParams // 游戏参数
		CreatorShortId int64             // 房间创建者
		AllianceId     int32             // 归属联盟

		Players []*Player

		gambling Gambling
		log      *logger.Logger
	}
)

func (r *Room) InjectGambling(gambling Gambling) {
	r.gambling = gambling
}

func (r *Room) GamblingHandle(shortId int64, v any) (result any) {
	var rsp any

	rsp = r.gambling.Handle(shortId, v)
	if rsp == nil {
		return nil
	}

	_, ok := rsp.(outer.ERROR)
	if ok {
		return rsp
	}

	return &inner.GamblingMsgToClientWrapper{
		MsgType: common.ProtoType(rsp.(proto.Message)),
		Data:    common.ProtoMarshal(rsp.(proto.Message)),
	}
}

func (r *Room) OnInit() {
	router.Result(r, r.responseHandle)
	log.Debugf("Room:[%v] OnInit", r.RoomId)
}

func (r *Room) OnStop() bool {
	log.Debugw("room stop", "room", r.RoomId)
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
		log.Warnw("room is stopping not handle msg", "room", r.RoomId)
		return
	}

	r.CurrentMsg = msg
	if routerErr := router.Dispatch(r, pt); routerErr != nil {
		log.Warnw("room dispatch the message failed", "err", routerErr)
	}
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
	if r.CurrentMsg.GetRequestId() != "" {
		err = r.Response(r.CurrentMsg.GetRequestId(), msg)
	} else {
		err = r.Send(r.CurrentMsg.GetSourceId(), msg)
	}

	if err != nil {
		log.Warnw("response to actor failed",
			"source", r.CurrentMsg.GetSourceId(),
			"msg name", r.CurrentMsg.GetMsgName())
	}
}

func (r *Room) IsFull() bool { return len(r.Players) >= gameMaxPlayers[r.GameType] }
func (r *Room) Disband()     { r.stopping = true }

func (r *Room) CanEnterRoom(p *inner.PlayerInfo) bool {
	if r.stopping {
		return false
	}
	return r.gambling.CanEnter(p)
}

func (r *Room) CanLeaveRoom(p *inner.PlayerInfo) bool {
	return r.gambling.CanLeave(p)
}

func (r *Room) CanReadyInRoom(p *inner.PlayerInfo) bool {
	return r.gambling.CanReady(p)
}

func (r *Room) CanSetGold(p *inner.PlayerInfo) bool {
	return r.gambling.CanSetGold(p)
}

func (r *Room) FindPlayer(shortId int64) *Player {
	for _, v := range r.Players {
		if v.ShortId == shortId {
			return v
		}
	}
	return nil
}
func (r *Room) FindPlayerByRID(rid string) *Player {
	for _, v := range r.Players {
		if v.RID == rid {
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
		EnterAt:    time.Now(),
	}
	r.Players = append(r.Players, newPlayer)
	r.gambling.PlayerEnter(newPlayer)

	seatIndex := r.gambling.SeatIndex(playerInfo.ShortId)
	r.Broadcast(&outer.RoomPlayerEnterNtf{Player: newPlayer.OuterPB(int32(seatIndex))})
	r.Log().Infow("room add player", "room", r.RoomId, "gameType", r.GameType,
		"player", playerInfo.ShortId, "seat", seatIndex, "gold", playerInfo.Gold)
	return nil
}

func (r *Room) PlayerLeave(shortId int64, kickOut bool) {
	var (
		rid    string
		delIdx int
	)

	for seatIdx, player := range r.Players {
		if player.ShortId == shortId {
			rid = player.RID
			r.gambling.PlayerLeave(player)
			delIdx = seatIdx
			r.Log().Infow("room del player", "room", r.RoomId, "shortId", shortId, "gold", player.Gold)
			break
		}
	}

	if rid == "" {
		return
	}

	r.Broadcast(&outer.RoomPlayerLeaveNtf{ShortId: shortId})
	r.Players = append(r.Players[:delIdx], r.Players[delIdx+1:]...)

	if kickOut {
		// 通知game，玩家被房间踢出
		playerActor := actortype.PlayerId(rid)
		_ = r.Send(playerActor, &inner.RoomKickOutNtf{RoomId: r.RoomId})
	}
}

func (r *Room) PlayerReady(shortId int64, ready bool) (ok bool, err outer.ERROR) {
	p := r.FindPlayer(shortId)
	if p == nil {
		// 玩家不在房间内
		r.Log().Warnw("leave the room cannot find player", "room", r.RoomId, "msg", shortId)
		return false, outer.ERROR_PLAYER_NOT_IN_ROOM
	}

	return true, 0
}

func (r *Room) SendToPlayer(shortId int64, msg proto.Message) {
	wrapper := &inner.GamblingMsgToClientWrapper{
		MsgType: common.ProtoType(msg),
		Data:    common.ProtoMarshal(msg),
	}

	player := r.FindPlayer(shortId)
	if player == nil {
		r.Log().Errorw("cannot find player", "room", r.RoomId, "shortId", shortId, "msg", common.ProtoType(msg))
		return
	}

	playerActor := actortype.PlayerId(player.RID)
	if err := r.Send(playerActor, wrapper); err != nil {
		r.Log().Errorw("send msg to player failed", "room", r.RoomId, "shortId", shortId, "player actor", playerActor)
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
		seadIndex := r.gambling.SeatIndex(player.ShortId)
		players = append(players, player.InnerPB(int32(seadIndex)))
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

// GamblingData 游戏数据
func (r *Room) GamblingData(shortId int64) []byte {
	data := r.gambling.Data(shortId)
	payload, err := proto.Marshal(data)
	if err != nil {
		r.Log().Errorw("gambling data marshal failed", "room", r.RoomId, "param", r.GameParams.String())
	}

	r.Log().Infow("gambling data", "room id ", r.RoomId, "data", data.String())
	return payload
}

func (p *Player) InnerPB(seatIdx int32) *inner.RoomPlayerInfo {
	return &inner.RoomPlayerInfo{
		SeatIndex: seatIdx,
		BaseInfo:  p.PlayerInfo,
		EnterAt:   p.EnterAt.UnixMilli(),
	}
}
func (p *Player) OuterPB(seatIdx int32) *outer.RoomPlayerInfo {
	return &outer.RoomPlayerInfo{
		SeatIndex: seatIdx,
		BaseInfo:  convert.PlayerInnerToOuter(p.PlayerInfo),
		EnterAt:   p.EnterAt.UnixMilli(),
	}
}

func (r *Room) Log() *logger.Logger {
	return r.log
}
