package rdsop

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"

	"server/common"
	"server/common/rds"
	"server/proto/outermsg/outer"
)

// 统计玩家金币的每一笔变更记录

type GoldUpdateType int32

const (
	UpModifyGold   GoldUpdateType = 1 // 被上级上\下分
	ModifyDownGold GoldUpdateType = 2 // 对下级上\下分
	GameWinOrLose  GoldUpdateType = 3 // 游戏结算输赢
	Rebate         GoldUpdateType = 4 // 领取返利
)

type GoldUpdateReason struct {
	Type        GoldUpdateType // 变更类型
	Gold        int64          // 增加的值(负数为减少)
	AfterGold   int64          // 变动后的值
	UpShortId   int64          // 上级ID 用于"被"上级 上下分
	DownShortId int64          // 下级ID 用于"对"下级 上下分
	GameType    int32          // 游戏类型 0.麻将 1.跑得快 用于游戏中输赢
	OccurAt     time.Time      // 发生时间
}

// SetUpdateGoldRecord 金币变化相关的记录
func SetUpdateGoldRecord(shortId int64, reason GoldUpdateReason, pip ...redis.Pipeliner) {
	key := UpdateGoldRecordKey(shortId)
	str := common.JsonMarshal(reason)
	ctx := context.Background()

	var pipeline redis.Cmdable
	if len(pip) > 0 {
		pipeline = pip[0]
	} else {
		pipeline = rds.Ins
	}

	// 插入新的元素在队列左侧,左边是新的，右边是老的
	pipeline.LPush(ctx, key, str)
	pipeline.LTrim(ctx, key, 0, 200)

	if reason.Type == GameWinOrLose {
		SetPlayerDailyStat(shortId, reason.Gold)
	}
}

// GetUpdateGoldRecord 获取玩家金币记录
func GetUpdateGoldRecord(shortId, start, end int64) (records []GoldUpdateReason, totalLen int64) {
	if start <= end && start >= 0 {
		return
	}

	key := UpdateGoldRecordKey(shortId)
	ctx := context.Background()
	totalLen = rds.Ins.LLen(ctx, key).Val()
	arr := rds.Ins.LRange(ctx, key, start, end).Val()
	for _, str := range arr {
		var reason GoldUpdateReason
		common.JsonUnmarshal(str, &reason)
		records = append(records, reason)
	}
	return
}

func (t GoldUpdateType) ToPB() outer.ReasonType {
	return outer.ReasonType(t)
}
