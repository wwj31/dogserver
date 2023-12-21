package rdsop

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/wwj31/dogactor/tools"

	"server/proto/outermsg/outer"

	"server/common/log"
	"server/common/rds"
	"server/proto/innermsg/inner"
)

func GetAndSetPlayerInfo(shortId int64, fn func(info *inner.PlayerInfo)) {
	key := fmt.Sprintf("lock:playerinfo:%v", shortId)
	rds.LockDo(key, func() {
		info := PlayerInfo(shortId)
		if info.ShortId == 0 {
			info.ShortId = shortId
		}

		fn(&info)

		b, err := json.Marshal(info)
		if err != nil {
			log.Errorw("SetPlayerInfo json marshal failed", "err", err, "info", info.String())
			return
		}
		rds.Ins.Set(context.Background(), PlayerInfoKey(info.ShortId), string(b), 0)
	})
}

func PlayerInfo(shortId int64) (info inner.PlayerInfo) {
	str, err := rds.Ins.Get(context.Background(), PlayerInfoKey(shortId)).Result()
	if err != nil {
		log.Errorw("PlayerInfo redis get failed", "err", err, "key", PlayerInfoKey(shortId))
		return
	}

	err = json.Unmarshal([]byte(str), &info)
	if err != nil {
		log.Errorw("PlayerInfo json unmarshal failed", "err", err, "key", PlayerInfoKey(shortId))
		return
	}
	return info
}

// SetPlayerDailyStat 统计玩家每日的输赢分数和场次
func SetPlayerDailyStat(shortId, score int64) {
	stat := &outer.PlayerDailyStat{}
	key := PlayerDailyStatKey(shortId)
	str, err := rds.Ins.Get(context.Background(), key).Result()
	if err == nil {
		jsonErr := json.Unmarshal([]byte(str), stat)
		if jsonErr != nil {
			log.Errorw("PlayerInfo json unmarshal failed", "err", jsonErr, "key", key, "str", str)
			return
		}
	} else if err != redis.Nil {
		log.Errorw("PlayerInfo redis get key failed", "err", err, "key", key)
		return
	}

	stat.ShortId = shortId
	stat.GameCount++
	stat.Gold += score

	b, _ := json.Marshal(stat)

	now := tools.Now()
	rds.Ins.Set(context.Background(), key, string(b), now.Add(2*24*time.Hour).Sub(now))
}

// PlayerDailyStat 获得玩家的今日统计信息
func PlayerDailyStat(shortIds ...int64) map[int64]*outer.PlayerDailyStat {
	pip := rds.Ins.Pipeline()
	ctx := context.Background()
	for _, id := range shortIds {
		key := PlayerDailyStatKey(id)
		pip.Get(ctx, key)
	}

	cmders, err := pip.Exec(ctx)
	if err != nil {
		log.Errorw("PlayerDailyStat json unmarshal failed", "err", err)
		return nil
	}

	result := map[int64]*outer.PlayerDailyStat{}
	for _, cmder := range cmders {
		str := cmder.(*redis.StringCmd).Val()
		stat := &outer.PlayerDailyStat{}
		if e := json.Unmarshal([]byte(str), stat); e != nil {
			log.Errorw("PlayerDailyStat json unmarshal failed", "err", err)
			continue
		}
		if stat.ShortId != 0 {
			result[stat.ShortId] = stat
		}
	}

	return result
}
