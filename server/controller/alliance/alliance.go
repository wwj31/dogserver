package alliance

import (
	"server/common"
	"server/proto/innermsg/inner"
	"server/service/alliance"

	"server/common/router"
)

// 保证玩家登录之前，联盟中玩家的GSession被正确关联
var _ = router.Reg(func(alliance *alliance.Alliance, msg *inner.OnlineNtf) any {
	alliance.PlayerOnline(common.GSession(msg.GateSession), msg.RID)
	return &inner.Ok{}
})

var _ = router.Reg(func(alliance *alliance.Alliance, msg *inner.OfflineNtf) any {
	alliance.PlayerOffline(common.GSession(msg.GateSession), msg.RID)
	return nil
})

var _ = router.Reg(func(alliance *alliance.Alliance, msg *inner.SetMemberReq) any {
	for _, player := range msg.Players {
		alliance.SetMember(player)
	}
	return &inner.SetMemberRsp{}
})
