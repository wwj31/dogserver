package player

import (
	"context"
	"reflect"
	"time"

	"server/service/alliance"

	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/tools"

	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/common/rds"
	"server/common/router"
	"server/proto/convert"
	"server/proto/innermsg/inner"
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

// // 获取下级每日游戏统计信息
var _ = router.Reg(func(p *player.Player, msg *outer.AgentDownDailyStatReq) any {
	stats := rdsop.PlayerDailyStat(msg.ShortIds...)
	return &outer.AgentDownDailyStatRsp{
		DownDailyStats: stats,
	}
})

// 获取自己的以及下级分配的返利信息
var _ = router.Reg(func(p *player.Player, msg *outer.AgentRebateInfoReq) any {
	rebate := rdsop.GetRebateInfo(p.Role().ShortId())

	allPlaying := rdsop.GetTodayPlaying()
	// 下级信息
	var infos []*outer.AgentDownRebateInfo
	downs := rdsop.AgentDown(p.Role().ShortId(), 1) // 获取直属下级
	for _, downShortId := range downs {
		if downShortId == p.Role().ShortId() {
			continue
		}

		t, y, w := rdsop.GetContribute(downShortId)
		downAll := rdsop.AgentDown(downShortId)
		playingCount := 0
		for _, id := range downAll {
			if _, exist := allPlaying[id]; exist {
				playingCount++
			}
		}

		infos = append(infos, &outer.AgentDownRebateInfo{
			ShortId:        downShortId,
			TodayGold:      t,
			YesterdayGold:  y,
			WeekGold:       w,
			TotalPlayers:   int32(len(downAll)),
			TodayPlayCount: int32(playingCount),
			DownPoint:      rebate.DownPoints[downShortId],
		})
	}

	return &outer.AgentRebateInfoRsp{
		OwnRebatePoints: rebate.Point,
		RebateInfos:     infos,
	}
})

// 设置下级的返利信息
var _ = router.Reg(func(p *player.Player, msg *outer.SetAgentDownRebateReq) any {
	if msg.Rebate < 0 || msg.Rebate > 100 {
		return outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	downShortIds := rdsop.AgentDown(p.Role().ShortId(), 1)
	exist := exist(msg.ShortId, downShortIds)

	// 不是直属下级，不能设置
	if !exist {
		return outer.ERROR_TARGET_IS_NOT_DOWN
	}

	err := rdsop.SetRebateInfo(p.Role().ShortId(), msg.ShortId, msg.Rebate)
	if err != outer.ERROR_OK {
		return err
	}

	downInfo := rdsop.PlayerInfo(msg.ShortId)
	//playerInfo := rdsop.PlayerInfo(p.Role().ShortId())
	if downInfo.Position <= alliance.Normal.Int32() {
		allianceActor := actortype.AllianceName(p.Alliance().AllianceId())
		rsp, err := p.RequestWait(allianceActor, &inner.SetMemberPositionReq{
			Player:   &downInfo,
			Position: int32(alliance.Captain),
		})
		if yes, err := common.IsErr(rsp, err); yes {
			return err
		}
	}

	return &outer.SetAgentDownRebateRsp{
		ShortId: msg.ShortId,
		Rebate:  msg.Rebate,
	}
})

// 团队总抽水
var _ = router.Reg(func(p *player.Player, msg *outer.TotalContributeReq) any {
	allDowns := rdsop.AgentDown(p.Role().ShortId())

	var total int64
	for _, shortId := range allDowns {
		key := rdsop.ContributeScoreKey(shortId)
		total += cast.ToInt64(rds.Ins.Get(context.Background(), key).Val())
	}
	return &outer.TotalContributeRsp{TotalGold: total}
})

// 获取返利详情信息
var _ = router.Reg(func(p *player.Player, msg *outer.AgentRebateDetailInfoReq) any {
	// 最近三天的每一笔返利信息
	records := rdsop.GetRebateRecordOf3Day(p.Role().ShortId())
	return &outer.AgentRebateDetailInfoRsp{Infos: records}
})

// 获取自己可领的返利分数信息
var _ = router.Reg(func(p *player.Player, msg *outer.RebateScoreReq) any {
	gold, goldOfToday, goldOfYesterday, goldOfWeek := rdsop.GetRebateGold(p.Role().ShortId())
	return &outer.RebateScoreRsp{
		Gold:            gold,
		GoldOfToday:     goldOfToday,
		GoldOfYesterday: goldOfYesterday,
		GoldOfWeek:      goldOfWeek,
	}
})

// 领取返利分
var _ = router.Reg(func(p *player.Player, msg *outer.ClaimRebateScoreReq) any {
	gold, _, _, _ := rdsop.GetRebateGold(p.Role().ShortId())
	if gold <= 0 {
		return outer.ERROR_MSG_REQ_PARAM_INVALID
	}

	p.Role().AddGold(gold)
	pip := rds.Ins.Pipeline()
	rdsop.IncRebateGold(p.Role().ShortId(), -gold, pip)

	reason := rdsop.GoldUpdateReason{
		Type:      rdsop.Rebate,
		Gold:      gold,
		AfterGold: p.Role().Gold(),
		OccurAt:   tools.Now(),
	}
	rdsop.SetUpdateGoldRecord(
		p.Role().ShortId(),
		reason,
		pip,
	)
	_, err := pip.Exec(context.Background())
	if err != nil {
		log.Errorw("claim rebate gold redis failed", "short", p.Role().ShortId(), "gold", gold)
		return outer.ERROR_FAILED
	}

	log.Infow("claim rebate gold", "short", p.Role().ShortId(), "gold", gold)
	return &outer.ClaimRebateScoreRsp{
		Gold: gold,
	}
})

// 对下级 上\下分操作
var _ = router.Reg(func(p *player.Player, msg *outer.SetScoreForDownReq) any {
	// 如果是上分，先判断自己钱够不够
	if msg.Gold > 0 && p.Role().Gold() < msg.Gold {
		return outer.ERROR_AGENT_UP_GOLD_FAILED
	}

	// 拿到直属下级,只能对直属下级操作
	downShortIds := rdsop.AgentDown(p.Role().ShortId(), 1)
	exist := exist(msg.ShortId, downShortIds)
	if !exist {
		return outer.ERROR_TARGET_IS_NOT_DOWN
	}

	downPlayerInfo := rdsop.PlayerInfo(msg.ShortId)
	if downPlayerInfo.RID == "" {
		log.Warnw("SetScoreForDownReq down rid is nil", "short", msg.ShortId)
		return outer.ERROR_CAN_NOT_FIND_PLAYER_INFO
	}

	// 下级还在房间，不能上下分操作
	if downPlayerInfo.RoomId != 0 {
		return outer.ERROR_AGENT_DOWN_GOLD_PLAYER_IN_ROOM
	}

	// 确保下级玩家处于激活状态
	// 找下级玩家最近登录过的game节点，如果没找到就用自己所在的game节点
	var dispatchGameId string
	gameNodeId, _ := rds.Ins.Get(context.Background(), rdsop.GameNodeKey(msg.ShortId)).Result()
	if gameNodeId != "" {
		dispatchGameId = gameNodeId
	} else {
		dispatchGameId = p.Gamer().ID()
	}

	v, pullErr := p.RequestWait(dispatchGameId, &inner.PullPlayer{RID: downPlayerInfo.RID})
	if yes, errCode := common.IsErr(v, pullErr); yes {
		log.Warnw("SetScoreForDownReq pull player failed", "msg", msg.String(), "err code", errCode)
		return outer.ERROR_FAILED
	}

	// 修改下级玩家的分数
	modifyMsg := &inner.ModifyGoldReq{
		Gold:      msg.Gold,
		Set:       false, // 增加分数
		SmallZero: false, // 上下分操作，不允许扣成负数
	}
	result, err := p.RequestWait(actortype.PlayerId(downPlayerInfo.RID), modifyMsg, time.Second)
	if err != nil {
		log.Warnw("SetScoreForDownReq request wait failed", "msg", msg.String(), "down player", downPlayerInfo)
		return outer.ERROR_FAILED
	}

	rsp, ok := result.(*inner.ModifyGoldRsp)
	if !ok {
		log.Warnw("SetScoreForDownReq ModifyGoldRsp assert failed",
			"msg", msg.String(), "down player", downPlayerInfo, "type", reflect.TypeOf(result).String())
		return outer.ERROR_FAILED
	}

	// 下级修改成功了，才修改自己的分
	if rsp.Success {
		p.Role().AddGold(-msg.Gold)

		//  都修改成功，就添加双方的分数变化记录
		pip := rds.Ins.Pipeline()
		// 记录自己的金币变化
		rdsop.SetUpdateGoldRecord(p.Role().ShortId(), rdsop.GoldUpdateReason{
			Type:        rdsop.ModifyDownGold, // 对下级上下分
			Gold:        -msg.Gold,
			AfterGold:   p.Role().Gold(),
			DownShortId: msg.ShortId,
			OccurAt:     tools.Now(),
		}, pip)

		// 记录下级的金币变化
		rdsop.SetUpdateGoldRecord(downPlayerInfo.ShortId, rdsop.GoldUpdateReason{
			Type:      rdsop.UpModifyGold, // 被上级上\下分
			Gold:      msg.Gold,
			AfterGold: rsp.Info.Gold,
			UpShortId: p.Role().ShortId(),
			OccurAt:   tools.Now(),
		}, pip)

		if _, err := pip.Exec(context.Background()); err != nil {
			log.Warnw("SetUpdateGoldRecord exec failed", "err", err)
		}
	}

	return &outer.SetScoreForDownRsp{
		ShortId: msg.ShortId,
		Gold:    rsp.Info.Gold,
	}
})

// 判断目标是否在指定的下级中
func exist(dstShortId int64, downShortIds []int64) bool {
	for _, id := range downShortIds {
		if id == dstShortId {
			return true
		}
	}
	return false
}
