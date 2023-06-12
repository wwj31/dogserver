package room

import (
	"server/proto/innermsg/inner"
	"server/service/room"

	"server/common/router"
)

// 创建房间
var _ = router.Reg(func(mgr *room.Mgr, msg *inner.CreateRoomReq) any {
	return &inner.CreateRoomRsp{}
})
