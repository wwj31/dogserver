package alliance

import (
	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
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
var _ = router.Reg(func(alli *alliance.Alliance, msg *inner.AddMemberReq) any {
	var position []alliance.Position
	if msg.Position != 0 {
		position = append(position, alliance.Position(msg.Position))
	}

	member := alli.AddMember(msg.Player, msg.Ntf, position...)

	// 获取成员所有的下级，全部加入本联盟
	downPlayers := rdsop.AgentDown(member.ShortId)
	for _, shortId := range downPlayers {
		playerInfo := rdsop.PlayerInfo(shortId)
		alli.AddMember(&playerInfo, true)
	}

	return &inner.AddMemberRsp{}
})

// 设置联盟成员职位
var _ = router.Reg(func(alli *alliance.Alliance, msg *inner.SetMemberPositionReq) any {
	member := alli.MemberInfo(msg.Player.RID)
	if member == nil {
		log.Warnw("cannot find member in alliance by setup position", "msg", msg.String())
		return &outer.FailRsp{Error: outer.ERROR_PLAYER_NOT_IN_ALLIANCE}
	}

	member.Position = alliance.Position(msg.Position)
	member.Save()

	// 如果对方在线，就通知更新
	if msg.Player.GSession != "" {
		// 玩家在线，通知Player actor修改联盟id
		err := alli.Send(actortype.PlayerId(member.RID), &inner.AllianceInfoNtf{
			AllianceId: alli.AllianceId(),
			Position:   member.Position.Int32(),
		})

		if err != nil {
			log.Warnw(" send to player actor error by set member position ", "err", err)
		}
	}
	return &inner.SetMemberPositionRsp{}
})

// 解散联盟
var _ = router.Reg(func(alli *alliance.Alliance, msg *inner.DisbandAllianceReq) any {
	if msg.RID != alli.Master().RID {
		return &inner.Error{ErrorInfo: "disband failed rid is not master"}
	}

	rdsop.DeleteAlliance(alli.AllianceId())

	// 获取联盟所有在线成员，并广播解散消息
	for _, member := range alli.Members() {
		if member.GSession.Valid() {
			playerActor := actortype.PlayerId(member.RID)
			_ = alli.Send(playerActor, &inner.AllianceInfoNtf{})
		}

		playerInfo := rdsop.PlayerInfo(member.ShortId)
		playerInfo.AllianceId = 0
		playerInfo.Position = 0
		rdsop.SetPlayerInfo(&playerInfo)
	}
	_, _ = alli.RequestWait(actortype.AllianceMgrName(), msg)

	alli.Disband()
	alli.Exit()

	return &inner.DisbandAllianceRsp{}
})

// 踢出成员
var _ = router.Reg(func(alli *alliance.Alliance, msg *inner.KickOutMembersReq) any {
	for _, shortId := range msg.GetShortIds() {
		member := alli.MemberInfoByShortId(shortId)
		if member == nil {
			log.Warnw("can not find the member by shortId",
				"alliance", alli.AllianceId(), "shortId", shortId)
			continue
		}

		playerInfo := rdsop.PlayerInfo(member.ShortId)
		if playerInfo.AllianceId != alli.AllianceId() {
			log.Warnw("kick out the member is not in alliance",
				"alli", alli.AllianceId(), "member alli", playerInfo.AllianceId)
			continue
		}

		alli.KickOutMember(shortId)

		playerInfo.Position = 0
		playerInfo.AllianceId = 0
		rdsop.SetPlayerInfo(&playerInfo)

		if member.GSession.Valid() {
			_ = alli.Send(actortype.PlayerId(member.RID), &inner.AllianceInfoNtf{
				AllianceId: 0,
				Position:   0,
			})
		}
	}

	return &inner.KickOutMembersRsp{}
})

// 请求联盟基础信息
var _ = router.Reg(func(alli *alliance.Alliance, msg *inner.AllianceInfoReq) any {
	return &inner.AllianceInfoRsp{
		MasterShortId: alli.Master().ShortId,
		MasterRID:     alli.Master().RID,
	}
})
