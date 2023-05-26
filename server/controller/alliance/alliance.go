package alliance

import (
	"server/common"
	"server/proto/innermsg/inner"
	"server/service/alliance"

	"server/common/router"
)

// 玩家登录，同步并请求数据
var _ = router.Reg(func(alliance *alliance.Alliance, msg *inner.MemberInfoOnLoginReq) any {
	alliance.PlayerOnline(common.GSession(msg.GateSession), msg.RID)
	member := alliance.MemberInfo(msg.RID)
	if member == nil {
		return &inner.MemberInfoOnLoginRsp{}
	}

	return &inner.MemberInfoOnLoginRsp{
		AllianceId: alliance.AllianceId(),
		Position:   member.Position.Int32(),
	}
})

// 玩家下线，通知联盟
var _ = router.Reg(func(alliance *alliance.Alliance, msg *inner.MemberInfoOnLogoutReq) any {
	alliance.PlayerOffline(common.GSession(msg.GateSession), msg.RID)
	return nil
})

var _ = router.Reg(func(alliance *alliance.Alliance, msg *inner.SetMemberReq) any {
	for _, player := range msg.Players {
		alliance.SetMember(player)
	}
	return &inner.SetMemberRsp{}
})
