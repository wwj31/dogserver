package rdsop

import (
	"context"
	"time"

	"github.com/wwj31/dogactor/tools"

	"server/common/log"
	"server/common/rds"
)

/*
	用户留存相关的统计接口
*/

// AddDailyRegistry 增加今日新用户
func AddDailyRegistry(shortId int64) {
	pip := rds.Ins.Pipeline()
	ctx := context.Background()
	key := DailyRegistrySetKey(tools.Now().Local())
	pip.SAdd(context.Background(), key, shortId)
	pip.Expire(ctx, key, 90*tools.Day)

	if _, er := pip.Exec(ctx); er != nil {
		log.Errorw("AddDailyRegistry pip exec failed", "err", er, "short", shortId)
	}
}

// AddDailyLogin 增加今日登录用户
func AddDailyLogin(shortId int64) {
	pip := rds.Ins.Pipeline()
	ctx := context.Background()
	key := DailyLoginSetKey(tools.Now().Local())
	pip.SAdd(context.Background(), key, shortId)
	pip.Expire(ctx, key, 90*tools.Day)

	if _, er := pip.Exec(ctx); er != nil {
		log.Errorw("AddDailyLogin pip exec failed", "err", er, "short", shortId)
	}
}

// RetentionOf 计算留存 regAt 某日新增留存 dayNum 1.次日留存，3.三日留存，7.七日留存，30.三十日留存
// 返回当日的注册人数，和留存
func RetentionOf(regAt time.Time, dayNum int) (regNum int64, rate float64) {
	// 当日的注册用户统计
	dailyRegistrySetKey := DailyRegistrySetKey(regAt)
	ctx := context.Background()
	num := rds.Ins.SCard(ctx, dailyRegistrySetKey).Val()
	if num == 0 {
		// 当日没有新增玩家
		return 0, 0
	}

	// 第n日的登录用户统计
	loginDateKey := DailyLoginSetKey(regAt.Add(time.Duration(dayNum) * tools.Day))

	// 取两个集合交集，获得regAt注册的用户，dayNum当天还继续登录的人数
	activateNum := rds.Ins.SInterCard(ctx, 0, dailyRegistrySetKey, loginDateKey).Val()
	return num, float64(activateNum) / float64(num)

}
