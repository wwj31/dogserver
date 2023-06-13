package player

import (
	"server/common"
	"server/common/actortype"
	"server/common/router"
	"server/proto/convert"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player"
)

// 创建房间
var _ = router.Reg(func(player *player.Player, msg *outer.CreateRoomReq) any {
	// TODO 其他创建条件

	v, err := player.RequestWait(actortype.RoomMgr(), &inner.CreateRoomReq{
		Creator: player.PlayerInfo(),
	})
	if yes, code := common.IsErr(v, err); yes {
		return code
	}
	createRoomRsp := v.(*inner.CreateRoomRsp)

	player.Room().SetRoomInfo(createRoomRsp.RoomInfo)
	roomInfo := convert.RoomInfoInnerToOuter(createRoomRsp.RoomInfo)
	return &outer.CreateRoomRsp{Room: roomInfo}
})

// 加入房间
var _ = router.Reg(func(p *player.Player, msg *outer.JoinRoomReq) any {
	if p.Room().RoomId() != 0 {
		return outer.ERROR_PLAYER_ALREADY_IN_ROOM
	}

	roomActor := actortype.RoomName(p.Room().RoomId())
	v, err := p.RequestWait(roomActor, &inner.JoinRoomReq{Player: p.PlayerInfo()})
	if yes, code := common.IsErr(v, err); yes {
		return code
	}
	joinRoomRsp := v.(*inner.JoinRoomRsp)

	p.Room().SetRoomInfo(joinRoomRsp.RoomInfo)
	roomInfo := convert.RoomInfoInnerToOuter(joinRoomRsp.RoomInfo)
	return &outer.JoinRoomRsp{Room: roomInfo}
})

// 离开房间
var _ = router.Reg(func(player *player.Player, msg *outer.LeaveRoomReq) any {
	// TODO ...
	return &outer.LeaveRoomRsp{}
})
