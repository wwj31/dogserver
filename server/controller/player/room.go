package player

import (
	"server/common"
	"server/common/actortype"
	"server/common/router"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player"
)

func roomInfoInnerToOuter(roomInfo *inner.RoomInfo) *outer.RoomInfo {
	return &outer.RoomInfo{
		Id:       roomInfo.RoomId,
		GameType: roomInfo.GameType,
	}
}

// 创建房间
var _ = router.Reg(func(player *player.Player, msg *outer.CreateRoomReq) any {
	// TODO 其他创建条件

	v, err := player.RequestWait(actortype.RoomMgr(), &inner.CreateRoomReq{CreatorShortId: player.Role().ShortId()})
	if yes, code := common.IsErr(v, err); yes {
		return code
	}
	createRoomRsp := v.(*inner.CreateRoomRsp)

	roomInfo := roomInfoInnerToOuter(createRoomRsp.RoomInfo)
	return &outer.CreateRoomRsp{Room: roomInfo}
})
