package rdsop

import (
	"context"
	"server/common"

	"server/common/rds"
)

type RebateInfo struct {
	Ratio     int32           `json:"ratio"`     // 自己的利润
	DownRatio map[int64]int32 `json:"downRatio"` // 给下级分的利润
}

// GetRebateInfo 获得玩家返利信息
func GetRebateInfo(shortId int64) RebateInfo {
	str, _ := rds.Ins.Get(context.Background(), AgentRebateKey(shortId)).Result()
	var result RebateInfo
	common.JsonUnmarshal(str, &result)
	return result
}

// SetRebateInfo 设置玩家返利信息
func SetRebateInfo(shortId int64, info RebateInfo) {
	str := common.JsonMarshal(info)
	rds.Ins.Set(context.Background(), AgentRebateKey(shortId), str, -1)
}
