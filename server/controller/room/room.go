package room

import (
	"server/common/router"
	"server/proto/innermsg/inner"
	"server/service/room"
)

// 获得房间信息
var _ = router.Reg(func(r *room.Room, msg *inner.RoomInfoReq) any {
	return &inner.RoomInfoRsp{RoomInfo: r.Info()}
})
