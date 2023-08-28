package rdsop

import (
	"context"
	"fmt"

	"github.com/spf13/cast"

	"server/common"
	"server/common/log"
	"server/common/rds"
	"server/proto/outermsg/outer"
)

type RebateInfo struct {
	Point      int32           `json:"point"`      // 自己的返利点
	DownPoints map[int64]int32 `json:"downPoints"` // 给下级分的返利点
}

// GetRebateInfo 获得玩家返利信息
func GetRebateInfo(shortId int64) RebateInfo {
	str, _ := rds.Ins.Get(context.Background(), AgentRebateKey(shortId)).Result()
	var result RebateInfo
	if str == "" {
		return result
	}

	common.JsonUnmarshal(str, &result)
	return result
}

// SetRebateInfo 设置下级玩家返利信息
func SetRebateInfo(shortId, downShortId int64, point int32) (err outer.ERROR) {
	getKey := func(id int64) string { return fmt.Sprintf("agent:rebate:%v", id) }

	rds.LockDo(getKey(shortId), func() {
		rebateInfo := GetRebateInfo(downShortId)
		rebateInfo.DownPoints[downShortId] = point
		var totalPoint int32
		for _, p := range rebateInfo.DownPoints {
			totalPoint += p
		}

		if rebateInfo.Point < totalPoint {
			err = outer.ERROR_AGENT_SET_REBATE_ONLY_OUT_OF_RANGE
			return
		}

		rds.LockDo(getKey(downShortId), func() {
			downRebateInfo := GetRebateInfo(downShortId)
			// 设置的点，必须大于等于该下级当前已有的点位
			if point < downRebateInfo.Point {
				err = outer.ERROR_AGENT_SET_REBATE_ONLY_HIGHER
				return
			}
			downRebateInfo.Point = point

			// 更新下级点位
			str := common.JsonMarshal(downRebateInfo)
			rds.Ins.Set(context.Background(), AgentRebateKey(downShortId), str, -1)

			// 更新自己管理的下级
			str = common.JsonMarshal(rebateInfo)
			rds.Ins.Set(context.Background(), AgentRebateKey(shortId), str, -1)
			log.Infow("set rebate info success", "short", shortId, "downShort", downShortId, "points", rebateInfo.DownPoints)
		})
	})

	return
}

// AddRebateGold 给玩家加返利分数
func AddRebateGold(shortId, score int64) {
	rds.Ins.IncrBy(context.Background(), RebateScoreKey(shortId), score)
}

// GetRebateGold 玩家返利分数
func GetRebateGold(shortId int64) int64 {
	val, err := rds.Ins.Get(context.Background(), RebateScoreKey(shortId)).Result()
	if err == nil {
		return cast.ToInt64(val)
	}
	return 0
}
