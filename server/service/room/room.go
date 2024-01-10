package room

import (
	"fmt"
	"path"
	"reflect"
	"runtime/debug"
	"time"

	"github.com/wwj31/dogactor/logger"
	"github.com/wwj31/dogactor/tools"

	"server/common/actortype"
	"server/common/log"
	"server/rdsop"

	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"

	"server/common"
	"server/common/router"
	"server/proto/convert"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
)

var gameMaxPlayers = map[GamblingType]int{
	Mahjong:   4,
	RunFaster: 3,
	NiuNiu:    10,
}

func New(info *rdsop.NewRoomInfo) *Room {
	var roomLogPath string
	if log.Path() != "" {
		roomLogPath = path.Join(log.Path(), "room")
	}
	r := &Room{
		RoomId:         info.RoomId,
		GameType:       info.GameType,
		GameParams:     info.Params,
		CreatorShortId: info.CreatorShortId,
		AllianceId:     info.AllianceId,
		ManifestId:     info.ManifestId,

		log: logger.New(logger.Option{
			Level:          logger.DebugLevel,
			LogPath:        roomLogPath,
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
		ManifestId     string            // 归属清单
		MasterShortId  int64             // 盟主短Id

		Players []*Player

		RecordInfo *outer.Recording // 回放数据
		gambling   Gambling
		log        *logger.Logger
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
	r.Request(actortype.AllianceName(r.AllianceId), &inner.AllianceInfoReq{}).Handle(func(resp any, err error) {
		if err != nil {
			r.Log().Errorw("alliance info req failed", "err", err)
			return
		}
		infoResp := resp.(*inner.AllianceInfoRsp)
		r.MasterShortId = infoResp.MasterShortId
	})

	r.Log().Debugf("Room:[%v] OnInit", r.RoomId)
}

func (r *Room) OnStop() bool {
	r.Log().Debugw("room stop", "room", r.RoomId)
	return true
}

func (r *Room) OnHandle(msg actor.Message) {
	defer func() {
		if err := recover(); err != nil {
			r.Log().Errorf(fmt.Sprintf("panic recover:%v msg(%v): %v \n%v", err, reflect.TypeOf(msg), msg.String(), string(debug.Stack())))
		}
	}()

	pt, ok := msg.Payload().(proto.Message)
	if !ok {
		r.Log().Warnw("room handler msg is not proto",
			"msg", reflect.TypeOf(msg.Payload()).String())
		return
	}

	if r.stopping {
		r.Log().Warnw("room is stopping not handle msg", "room", r.RoomId)
		return
	}

	r.CurrentMsg = msg
	if routerErr := router.Dispatch(r, pt); routerErr != nil {
		r.Log().Warnw("room dispatch the message failed", "err", routerErr)
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
		r.Log().Warnw("response to actor failed",
			"source", r.CurrentMsg.GetSourceId(),
			"msg name", r.CurrentMsg.GetMsgName())
	}
}

func (r *Room) IsFull() bool {
	maxNum := gameMaxPlayers[r.GameType]
	if r.GameType == Mahjong {
		switch r.GameParams.Mahjong.GameMode {
		case 0:
			maxNum = 4
		case 1, 2:
			maxNum = 3
		case 3, 4:
			maxNum = 2
		default:
			r.Log().Errorw("unknown majong game mode", "room", r.RoomId, "mode", r.GameParams.Mahjong.GameMode)
			return true
		}
	}
	return len(r.Players) >= maxNum
}

func (r *Room) IsEmpty() bool   { return len(r.Players) == 0 }
func (r *Room) Disband()        { r.stopping = true }
func (r *Room) IsDisband() bool { return r.stopping }

func (r *Room) CanEnterRoom(p *inner.PlayerInfo) bool {
	if r.stopping {
		return false
	}
	return r.gambling.CanEnter(p)
}

func (r *Room) CanLeaveRoom(p *inner.PlayerInfo) bool {
	return r.gambling.CanLeave(p)
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

func (r *Room) PlayerOnline(shortId int64) {
	r.Broadcast(&outer.RoomPlayerOnlineNtf{ShortId: shortId, Online: true}, shortId)
	r.gambling.PlayerOnline(shortId)
}
func (r *Room) PlayerOffline(shortId int64) {
	r.Broadcast(&outer.RoomPlayerOnlineNtf{ShortId: shortId, Online: false}, shortId)
	r.gambling.PlayerOffline(shortId)
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
	msgType := common.ProtoType(msg)
	Data := common.ProtoMarshal(msg)

	wrapper := &inner.GamblingMsgToClientWrapper{
		MsgType: msgType,
		Data:    Data,
	}

	pbIndex := r.System().ProtoIndex()
	msgId, _ := pbIndex.MsgNameToId(msgType)

	if r.gambling.CanRecordingPlayback() {
		r.RecordInfo.Messages = append(r.RecordInfo.Messages, &outer.RecordingMessage{
			SendAt:  tools.Now().UnixMilli(),
			ShortId: shortId,
			MsgId:   outer.Msg(msgId),
			Data:    Data,
		})
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

	r.Log().Infow("gambling data", "shortId", shortId, "data", data.String())
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

// GameRecordingStart 开始记录
func (r *Room) GameRecordingStart() {
	// 在开始记录回放的点，把房间状态先存下来
	r.RecordInfo = &outer.Recording{
		GameStartAt: tools.Now().UnixMilli(),
		Room:        convert.RoomInfoInnerToOuter(r.Info()),
	}
}

// GameRecordingOver 结束记录 传入底分和每个玩家的输赢情况
func (r *Room) GameRecordingOver(baseScore int64, winScore map[int64]int64) {
	r.RecordInfo.GameOverAt = tools.Now().UnixMilli()
	// 结束记录的时候，房间回放信息推上redis
	rdsop.AddRoomRecording(r.RecordInfo)

	// 玩家游戏记录
	msg := &inner.GameHistoryInfoReq{
		Info: &inner.HistoryInfo{
			GameType:    r.GameType,
			RoomId:      r.RoomId,
			GameStartAt: r.RecordInfo.GameStartAt,
			GameOverAt:  r.RecordInfo.GameOverAt,
			BaseScore:   baseScore,
		},
	}
	for shortId, win := range winScore {
		msg.Info.WinGold = win
		player := r.FindPlayer(shortId)
		_ = r.Send(actortype.PlayerId(player.RID), msg)
	}
}
