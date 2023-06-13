package room

import (
	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/proto/innermsg/inner"
	"server/service/room"

	"server/common/router"
)

// 创建房间
var _ = router.Reg(func(mgr *room.Mgr, msg *inner.CreateRoomReq) any {
	roomId := mgr.RoomId()
	newRoom := room.New(roomId, msg.GameType, msg.Creator)
	_ = mgr.AddRoom(newRoom)

	roomActor := actortype.RoomName(roomId)

	if err := mgr.System().NewActor(roomActor, newRoom); err != nil {
		log.Errorw("create room failed", "msg", msg, "err", err)
		return &inner.Error{ErrorInfo: err.Error()}
	}

	v, err := mgr.RequestWait(roomActor, &inner.RoomInfoReq{})
	if yes, code := common.IsErr(v, err); yes {
		return code
	}

	roomInfoRsp := v.(*inner.RoomInfoRsp)
	return &inner.CreateRoomRsp{RoomInfo: roomInfoRsp.RoomInfo}
})
