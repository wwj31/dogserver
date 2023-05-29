package alliance

import (
	"server/common"
	"server/proto/innermsg/inner"
	"server/rdsop"
	"server/service/alliance"

	"server/common/router"
)

// 玩家登录，同步并请求数据
var _ = router.Reg(func(alli *alliance.Alliance, msg *inner.MemberInfoOnLoginReq) any {
	alli.PlayerOnline(common.GSession(msg.GateSession), msg.RID)
	member := alli.MemberInfo(msg.RID)
	if member == nil {
		return &inner.MemberInfoOnLoginRsp{}
	}

	return &inner.MemberInfoOnLoginRsp{
		AllianceId: alli.AllianceId(),
		Position:   member.Position.Int32(),
	}
})

// 玩家下线，通知联盟
var _ = router.Reg(func(alli *alliance.Alliance, msg *inner.MemberInfoOnLogoutReq) any {
	alli.PlayerOffline(common.GSession(msg.GateSession), msg.RID)
	return nil
})

// 设置联盟成员
var _ = router.Reg(func(alli *alliance.Alliance, msg *inner.SetMemberReq) any {
	var position []alliance.Position
	if msg.Position != 0 {
		position = append(position, alliance.Position(msg.Position))
	}

	member := alli.SetMember(msg.Player, msg.Ntf, position...)

	// 获取成员所有的下级，全部加入本联盟
	downPlayers := rdsop.AgentDown(member.ShortId)
	for _, shortId := range downPlayers {
		playerInfo := rdsop.PlayerInfo(shortId)
		alli.SetMember(&playerInfo, true)
	}

	return &inner.SetMemberRsp{Position: member.Position.Int32()}
})
