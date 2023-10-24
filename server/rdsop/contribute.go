package rdsop

import (
	"context"
	"fmt"

	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/tools"

	"server/common/log"
	"server/common/rds"
)

// 玩家游戏提供的业绩（界面上的抽水）

// SetContribute 记录玩家今日、本周、总业绩
func SetContribute(shortId, gold int64) {
	todayKey := ContributeScoreKeyForToday(shortId)
	weekKey := ContributeScoreKeyForWeek(shortId)
	totalKey := ContributeScoreKey(shortId)

	pip := rds.Ins.Pipeline()
	ctx := context.Background()
	pip.IncrBy(ctx, todayKey, gold)
	pip.IncrBy(ctx, weekKey, gold)
	pip.IncrBy(ctx, totalKey, gold)

	pip.Expire(ctx, todayKey, 3*tools.Day)
	pip.Expire(ctx, weekKey, 8*tools.Day)

	if _, er := pip.Exec(ctx); er != nil {
		log.Errorw("SetContribute pip exec failed", "err", er, "short", shortId, "gold", gold)
	}
	log.Infow("set contribute", "short", shortId, "gold", gold)
}

// GetContribute 获得今日、昨日、本周业绩值
func GetContribute(shortId int64) (todayGold, yesterdayGold, weekGold int64) {
	todayKey := ContributeScoreKeyForToday(shortId)
	yesterdayKey := ContributeScoreKeyForYesterday(shortId)
	weekKey := ContributeScoreKeyForWeek(shortId)

	ctx := context.Background()
	todayGold = cast.ToInt64(rds.Ins.Get(ctx, todayKey).Val())
	yesterdayGold = cast.ToInt64(rds.Ins.Get(ctx, yesterdayKey).Val())
	weekGold = cast.ToInt64(rds.Ins.Get(ctx, weekKey).Val())
	return
}

// SetTodayPlaying 设置今日参与游戏的人数
func SetTodayPlaying(shortIds ...int64) {
	key := fmt.Sprintf("playing:%v", tools.Now().Local().Format(tools.StdDateFormat))
	pip := rds.Ins.Pipeline()

	ctx := context.Background()
	pip.SAdd(ctx, key, shortIds)
	pip.Expire(ctx, key, 2*tools.Day)

	pip.Exec(ctx)
	return
}

func GetTodayPlaying() map[int64]struct{} {
	key := fmt.Sprintf("playing:%v", tools.Now().Local().Format(tools.StdDateFormat))
	set := rds.Ins.SMembersMap(context.Background(), key).Val()

	result := make(map[int64]struct{})
	for str, _ := range set {
		result[cast.ToInt64(str)] = struct{}{}
	}
	return result
}
