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

	// TODO ...
	return &inner.JoinRoomRsp{RoomInfo: r.Info()}
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
