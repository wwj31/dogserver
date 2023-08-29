package player

import (
	"context"

	"server/common/log"
	"server/common/rds"
	"server/common/router"
	"server/proto/convert"
	"server/proto/outermsg/outer"
	"server/rdsop"
	"server/service/game/logic/player"
)

// // 获取上、下级信息
var _ = router.Reg(func(p *player.Player, msg *outer.AgentMembersReq) any {
	var (
		upMember    *outer.PlayerInfo
		downMembers []*outer.PlayerInfo
	)

	upShortId := rdsop.AgentUp(p.Role().ShortId())
	if upShortId != 0 {
		upInfo := rdsop.PlayerInfo(upShortId)
		upMember = convert.PlayerInnerToOuter(&upInfo)
	}

	downShortIds := rdsop.AgentDown(p.Role().ShortId())
	for _, shortId := range downShortIds {
		if shortId == p.Role().ShortId() {
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
var _ = router.Reg(func(p *player.Player, msg *outer.AgentRebateInfoReq) any {
	rebate := rdsop.GetRebateInfo(p.Role().ShortId())
	return &outer.AgentRebateInfoRsp{
		OwnRebatePoints: rebate.Point,
		DownPoints:      rebate.DownPoints,
	}
})

// 设置下级的返利信息
var _ = router.Reg(func(p *player.Player, msg *outer.SetAgentDownRebateReq) any {
	if msg.Rebate < 0 || msg.Rebate > 100 {
		return outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	downs := rdsop.AgentDown(p.Role().ShortId(), 1)
	var ok bool
	for _, downShortId := range downs {
		if downShortId == msg.ShortId {
			ok = true
			break
		}
	}

	// 不是直属下级，不能设置
	if !ok {
		return outer.ERROR_TARGET_IS_NOT_DOWN
	}

	err := rdsop.SetRebateInfo(p.Role().ShortId(), msg.ShortId, msg.Rebate)
	if err != outer.ERROR_OK {
		return err
	}

	return &outer.SetAgentDownRebateRsp{
		ShortId: msg.ShortId,
		Rebate:  msg.Rebate,
	}
})

// 获取自己可领的返利分数信息
var _ = router.Reg(func(p *player.Player, msg *outer.RebateScoreReq) any {
	return &outer.RebateScoreRsp{
		Gold: rdsop.GetRebateGold(p.Role().ShortId()),
	}
})

// 领取返利分
var _ = router.Reg(func(p *player.Player, msg *outer.ClaimRebateScoreReq) any {
	score := rdsop.GetRebateGold(p.Role().ShortId())
	if score <= 0 {
		return outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	p.Role().AddGold(score)
	pip := rds.Ins.Pipeline()
	rdsop.AddRebateGold(p.Role().ShortId(), -score, pip)
	_, err := pip.Exec(context.Background())
	if err != nil {
		log.Errorw("claim rebate score redis failed", "short", p.Role().ShortId(), "score", score)
		return outer.ERROR_FAILED
	}

	log.Infow("claim rebate gold", "short", p.Role().ShortId(), "score", score)
	return &outer.ClaimRebateScoreRsp{
		Gold: score,
	}
})
