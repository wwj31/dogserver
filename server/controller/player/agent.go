package player

import (
	"server/common/rds"
	"server/common/router"
	"server/proto/convert"
	"server/proto/outermsg/outer"
	"server/rdsop"
	"server/service/game/logic/player"
)

// // 获取上、下级信息
var _ = router.Reg(func(player *player.Player, msg *outer.AgentMembersReq) any {
	var (
		upMember    *outer.PlayerInfo
		downMembers []*outer.PlayerInfo
	)

	upShortId := rdsop.AgentUp(player.Role().ShortId())
	if upShortId != 0 {
		upInfo := rdsop.PlayerInfo(upShortId)
		upMember = convert.PlayerInnerToOuter(&upInfo)
	}

	downShortIds := rdsop.AgentDown(player.Role().ShortId())
	for _, shortId := range downShortIds {
		if shortId == player.Role().ShortId() {
			continue
		}

		downInfo := rdsop.PlayerInfo(shortId)
		downMembers = append(downMembers, convert.PlayerInnerToOuter(&downInfo))
	}

	return &outer.AgentMembersRsp{
		UpMember:    upMember,
		DownMembers: downMembers,
	}
})

// 获取自己的以及下级分配的返利信息
var _ = router.Reg(func(player *player.Player, msg *outer.AgentRebateInfoReq) any {
	rebate := rdsop.GetRebateInfo(player.Role().ShortId())
	return &outer.AgentRebateInfoRsp{
		OwnRebatePoints: rebate.Point,
		DownPoints:      rebate.DownPoints,
	}
})

// 设置下级的返利信息
var _ = router.Reg(func(player *player.Player, msg *outer.SetAgentDownRebateReq) any {
	rds.LockDo()
	rebate := rdsop.GetRebateInfo(player.Role().ShortId())
	return &outer.SetAgentDownRebateRsp{}
})
