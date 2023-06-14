package room

import (
	"server/common/log"
	"server/common/router"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/room"
)

// 获得房间信息
var _ = router.Reg(func(r *room.Room, msg *inner.RoomInfoReq) any {
	return &inner.RoomInfoRsp{RoomInfo: r.Info()}
})

// 加入房间
var _ = router.Reg(func(r *room.Room, msg *inner.JoinRoomReq) any {
	// 玩家已在房间内
	if r.FindPlayer(msg.Player.ShortId) != nil {
		return &inner.Error{ErrorCode: int32(outer.ERROR_PLAYER_ALREADY_IN_ROOM)}
	}

	if r.IsFull() {
		return &inner.Error{ErrorCode: int32(outer.ERROR_ROOM_WAS_FULL_CAN_NOT_ENTER)}
	}

	// TODO 其他进入房间的条件

	r.AddPlayer(msg.Player)
	return &inner.JoinRoomRsp{RoomInfo: r.Info()}
})

// 解散房间
var _ = router.Reg(func(r *room.Room, msg *inner.DisbandRoomReq) any {
	// 房间还有人，不能解散
	if len(r.Players) > 0 {
		return &inner.Error{ErrorCode: int32(outer.ERROR_ROOM_HAS_PLAYER_CAN_NOT_DISBAND)}
	}

	r.Disband()
	r.Exit()
	return &inner.DisbandRoomRsp{}
})

// 离开房间
var _ = router.Reg(func(r *room.Room, msg *inner.LeaveRoomReq) any {
	// 玩家不在房间内
	if r.FindPlayer(msg.ShortId) == nil {
		log.Warnw("leave the room cannot find player", "roomId", r.RoomId, "msg", msg.ShortId)
		return &inner.LeaveRoomRsp{}
	}

	// TODO ...
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
	return &inner.RoomLoginRsp{RoomInfo: r.Info()}
})

// 下线通知
var _ = router.Reg(func(r *room.Room, msg *inner.RoomLogoutReq) any {
	p := r.FindPlayer(msg.GetShortId())
	if p == nil {
		log.Warnw("room login cannot find player", "msg", msg.String())
		return &inner.RoomLogoutRsp{}
	}
	p.GSession = ""

	return &inner.RoomLogoutRsp{}
})
