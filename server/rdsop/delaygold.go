package rdsop

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"

	"server/common/log"
	"server/common/rds"
)

const delayGold = "delayGold"

func delayGoldKey(shortId int64) string {
	return fmt.Sprintf("%v:%v", delayGold, shortId)
}

func AddDelayModifyGold(shortId int64, gold int64) int64 {
	val, err := rds.Ins.IncrBy(context.Background(), delayGoldKey(shortId), gold).Result()
	if err != nil {
		log.Errorw("AddDelayModifyGold failed ", "err", err)
		return 0
	}
	return val
}

func SyncGold(shortId int64, fn func(val int64)) {
	result, err := rds.Ins.Get(context.Background(), delayGoldKey(shortId)).Result()
	if err != nil {
		if err != redis.Nil {
			log.Errorw("AddDelayModifyGold failed ", "err", err)
		}
		return
	}

	val := cast.ToInt64(result)
	for val != 0 {
		fn(val)
		val = AddDelayModifyGold(shortId, -val)
	}
}
