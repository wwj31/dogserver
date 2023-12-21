package rdsop

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/wwj31/dogactor/tools"

	"server/common/rds"
)

/*
	用户统计
*/

/////////////////////////////////////////////////////	留存相关的统计接口

// AddDailyRegistry 增加今日新用户
func AddDailyRegistry(shortId int64, pip redis.Pipeliner) {
	ctx := context.Background()
	key := DailyRegistrySetKey(tools.Now().Local())
	pip.SAdd(context.Background(), key, shortId)
	pip.Expire(ctx, key, 90*tools.Day)
}

// AddDailyLogin 增加今日登录用户
func AddDailyLogin(shortId int64, pip redis.Pipeliner) {
	ctx := context.Background()
	key := DailyLoginSetKey(tools.Now().Local())
	pip.SAdd(context.Background(), key, shortId)
	pip.Expire(ctx, key, 90*tools.Day)
}

// Retention 计算留存 regAt 某日新增留存 dayNum 1.次日留存，3.三日留存，7.七日留存，30.三十日留存
// 返回当日的注册人数，和留存
func Retention(regAt time.Time, dayNum int) (regNum int64, rate float64) {
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

///////////////////////////////////////////////////// 实时在线人数统计

// SetRealTimeUser 设置用户在线
func SetRealTimeUser(shortId int64, pip redis.Pipeliner) {
	ctx := context.Background()
	key := RealTimeUserKey(tools.Now().Local())
	pip.SAdd(context.Background(), key, shortId)
	pip.Expire(ctx, key, tools.Day)
}

// UnsetRealTimeUser 取消用户在线
func UnsetRealTimeUser(shortId int64, pip redis.Pipeliner) {
	ctx := context.Background()
	key := RealTimeUserKey(tools.Now().Local())
	pip.SRem(context.Background(), key, shortId)
	pip.Expire(ctx, key, tools.Day)
}

// RealTimeUserCount 活动实时在线人数
func RealTimeUserCount() int64 {
	ctx := context.Background()
	key := RealTimeUserKey(tools.Now().Local())
	return rds.Ins.SCard(ctx, key).Val()
}
