package room

import (
	"context"

	"github.com/go-redis/redis/v9"

	"server/common"
	"server/common/log"
	"server/common/rds"
	"server/proto/outermsg/outer"
	"server/rdsop"
)

// ProfitRange 根据赢的分和返利参数，或者需要执行的返利区间参数
func (r *Room) ProfitRange(score int64, rebate *outer.RebateParams) *outer.RangeParams {
	if rebate == nil {
		return nil
	}

	for _, param := range []*outer.RangeParams{rebate.RangeL1, rebate.RangeL2, rebate.RangeL3, rebate.RangeL4} {
		if !param.Valid {
			continue
		}

		minLimit := param.Min * common.Gold1000Times
		MaxLimit := param.Max * common.Gold1000Times

		if minLimit < score && score <= MaxLimit {
			return param
		}
	}
	return nil
}

// Rebate 返利计算
func (r *Room) Rebate(record *outer.RebateDetailInfo, totalProfit int64, players []*Player) {
	divProfit := totalProfit / int64(len(players)) // 每个玩家那条上级路线，都均分获得利润
	if divProfit == 0 {
		r.Log().Infow("profit is zero!!", "total profit", totalProfit)
		return
	}

	pip := rds.Ins.Pipeline()
	// 本局参与游戏的玩家挨个地、逐层地、依次向上级返利
	for _, player := range players {
		r.Log().Infow("recur profit start", "start short", player.ShortId, "profit", divProfit)
		r.recurRebate(record, divProfit, player.UpShortId, player.ShortId, 0, pip)

		// 统计玩家贡献（界面中的抽水统计）
		rdsop.SetContribute(player.ShortId, divProfit)
	}

	if _, err := pip.Exec(context.Background()); err != nil {
		log.Errorw("rebate redis failed", "err", err)
	}
}

func (r *Room) recurRebate(record *outer.RebateDetailInfo, profitGold, upShortId, shortId, downShortId int64, addPip redis.Pipeliner) {
	rebateInfo := rdsop.GetRebateInfo(shortId)

	var (
		exactPoint      int32 // 确切的获利点位
		exactProfitGold int64 // 确切的获利
	)

	if rebateInfo.Point > 0 {
		// 自己的获利=自己的分润-下级的分润
		exactPoint = common.Max(0, rebateInfo.Point-rebateInfo.DownPoints[downShortId])
		exactProfitGold = profitGold * int64(exactPoint) / 100 // 实际分润的金币
		record.Gold = exactProfitGold
		record.ShortId = shortId
		rdsop.RecordRebateGold(common.JsonMarshal(record), shortId, exactProfitGold, addPip)
	}

	r.Log().Infow("rebate calculating ...", "room", r.RoomId, "game", r.GameType,
		"short", shortId, "up", upShortId, "down", downShortId, "rebateInfo", rebateInfo,
		"profitGold", profitGold, "point", exactPoint, "gold", exactProfitGold)

	if upShortId == 0 {
		return
	}

	// 向上级递归
	r.recurRebate(record, profitGold, rdsop.AgentUp(upShortId), upShortId, shortId, addPip)
}
