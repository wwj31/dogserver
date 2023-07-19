package room

import (
	"server/common/log"
	"server/common/router"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/rdsop"
	"server/service/room"
)

// 解散房间
var _ = router.Reg(func(r *room.Room, msg *inner.DisbandRoomReq) any {
	// 房间还有人，不能解散
	if len(r.Players) > 0 {
		return outer.ERROR_ROOM_HAS_PLAYER_CAN_NOT_DISBAND
	}

	r.Disband()
	r.Exit()
	rdsop.DelRoomInfoFromRedis(r.RoomId)
	rdsop.SubAllianceRoom(r.RoomId, r.AllianceId)
	return &inner.DisbandRoomRsp{}
})

// 获得房间信息
var _ = router.Reg(func(r *room.Room, msg *inner.RoomInfoReq) any {
	return &inner.RoomInfoRsp{RoomInfo: r.Info()}
})

// 加入房间
var _ = router.Reg(func(r *room.Room, msg *inner.JoinRoomReq) any {
	if !r.CanEnterRoom(msg.Player) {
		return outer.ERROR_ROOM_CAN_NOT_ENTER
	}

	// 玩家不在房间所属联盟中
	if r.AllianceId != msg.Player.AllianceId {
		return outer.ERROR_PLAYER_NOT_IN_ALLIANCE
	}

	// 玩家已在房间内
	if r.FindPlayer(msg.Player.ShortId) != nil {
		return outer.ERROR_PLAYER_ALREADY_IN_ROOM
	}

	// 房间满员
	if r.IsFull() {
		return outer.ERROR_ROOM_WAS_FULL_CAN_NOT_ENTER
	}

	if err := r.PlayerEnter(msg.Player); err != nil {
		return err
	}

	return &inner.JoinRoomRsp{
		RoomInfo:     r.Info(),
		GamblingData: r.GamblingData(msg.Player.ShortId),
	}
})

// 离开房间
var _ = router.Reg(func(r *room.Room, msg *inner.LeaveRoomReq) any {
	// 玩家不在房间内
	p := r.FindPlayer(msg.ShortId)
	if p == nil {
		log.Warnw("leave the room cannot find player", "room", r.RoomId, "msg", msg.ShortId)
		return &inner.LeaveRoomRsp{}
	}

	// 房间当前状态不能离开
	if !r.CanLeaveRoom(p.PlayerInfo) {
		return outer.ERROR_ROOM_CAN_NOT_LEAVE
	}

	r.PlayerLeave(msg.ShortId, false)
	return &inner.LeaveRoomRsp{}
})

// 上线通知
var _ = router.Reg(func(r *room.Room, msg *inner.RoomLoginReq) any {
	p := r.FindPlayer(msg.Player.GetShortId())
	var err int32
	if p == nil {
		log.Warnw("room login cannot find player", "msg", msg.String())
		return &inner.RoomLoginRsp{Err: err}
	}

	p.PlayerInfo = msg.Player
	r.Broadcast(&outer.RoomPlayerOnlineNtf{ShortId: p.ShortId, Online: true}, p.ShortId)
	return &inner.RoomLoginRsp{
		RoomInfo:     r.Info(),
		GamblingData: r.GamblingData(msg.Player.ShortId),
	}
})

// 下线通知
var _ = router.Reg(func(r *room.Room, msg *inner.RoomLogoutReq) any {
	p := r.FindPlayer(msg.GetShortId())
	if p == nil {
		log.Warnw("room login cannot find player", "msg", msg.String())
		return nil
	}
	p.GSession = ""

	r.Broadcast(&outer.RoomPlayerOnlineNtf{ShortId: msg.ShortId, Online: false}, msg.GetShortId())
	return nil
})

// gambling 消息
var _ = router.Reg(func(r *room.Room, msg *inner.GamblingMsgToRoomWrapper) any {
	p := r.FindPlayerByRID(msg.RID)
	if p == nil {
		log.Warnw("room GamblingMsgToRoomWrapper cannot find player", "msg", msg.String())
		return outer.ERROR_PLAYER_NOT_IN_ROOM
	}

	msgId, ok := r.System().ProtoIndex().MsgNameToId(msg.MsgType)
	if !ok {
		log.Warnw("MsgGamblingMsgToClientWrapper msg name to id failed", "player", p.ShortId, "room", r.RoomId, "type", r.GameType, "msg", msg.String())
		return outer.ERROR_FAILED
	}

	outerMsg := r.System().ProtoIndex().UnmarshalPbMsg(msgId, msg.Data)
	rsp := r.GamblingHandle(p.ShortId, outerMsg)

	return rsp
})
