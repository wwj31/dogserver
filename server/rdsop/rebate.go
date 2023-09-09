package rdsop

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
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
	if result.DownPoints == nil {
		result.DownPoints = map[int64]int32{}
	}

	return result
}

// SetRebateInfoByDoor 后台设置玩家返利信息
func SetRebateInfoByDoor(shortId int64, point int32) (info RebateInfo) {
	rds.LockDo(rebateSetSyncKey(shortId), func() {
		bateInfo := GetRebateInfo(shortId)
		bateInfo.Point = point
		info = bateInfo
		str := common.JsonMarshal(bateInfo)
		rds.Ins.Set(context.Background(), AgentRebateKey(shortId), str, -1)
	})
	return
}

func rebateSetSyncKey(id int64) string {
	return fmt.Sprintf("agent:rebate:%v", id)
}

// SetRebateInfo 设置下级玩家返利信息
func SetRebateInfo(shortId, downShortId int64, point int32) (err outer.ERROR) {
	rds.LockDo(rebateSetSyncKey(shortId), func() {
		rebateInfo := GetRebateInfo(downShortId)
		rebateInfo.DownPoints[downShortId] = point

		// 获得点位最高的下级
		var HighestPoint int32
		for _, p := range rebateInfo.DownPoints {
			HighestPoint = common.Max(HighestPoint, p)
		}

		if rebateInfo.Point < HighestPoint {
			err = outer.ERROR_AGENT_SET_REBATE_ONLY_OUT_OF_RANGE
			return
		}

		rds.LockDo(rebateSetSyncKey(downShortId), func() {
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
func AddRebateGold(shortId, score int64, pip redis.Pipeliner) {
	ctx := context.Background()

	// 累计返利
	pip.IncrBy(ctx, RebateGoldKey(shortId), score)

	// 今日统计返利
	statTodayKey := RebateScoreKeyForToday(shortId)
	pip.IncrBy(ctx, statTodayKey, score)
	pip.Expire(ctx, statTodayKey, 24*time.Hour)

	// 本周统计统计返利
	statWeekKey := RebateScoreKeyForWeek(shortId)
	pip.IncrBy(ctx, statWeekKey, score)
	pip.Expire(ctx, statWeekKey, 7*24*time.Hour)
}

// GetRebateGold 玩家返利分数
func GetRebateGold(shortId int64) (gold, goldOfToday, goldOfWeek int64) {
	pip := rds.Ins.Pipeline()
	ctx := context.Background()

	var (
		cmds                   []*redis.StringCmd
		rebateGoldKey          = RebateGoldKey(shortId)
		rebateScoreKeyForToday = RebateScoreKeyForToday(shortId)
		rebateScoreKeyForWeek  = RebateScoreKeyForWeek(shortId)
	)

	cmds = append(cmds, pip.Get(ctx, rebateGoldKey))
	cmds = append(cmds, pip.Get(ctx, rebateScoreKeyForToday))
	cmds = append(cmds, pip.Get(ctx, rebateScoreKeyForWeek))
	_, err := pip.Exec(ctx)
	if err != nil {
		log.Errorw("get rebate gold failed", "err", err)
		return
	}

	if len(cmds) < 3 {
		log.Errorw("get rebate gold unexpected outcome", "short", shortId, "cmd len", cmds,
			"rebateGoldKey", rebateGoldKey,
			"rebateScoreKeyForToday", rebateScoreKeyForToday,
			"rebateScoreKeyForWeek", rebateScoreKeyForWeek)
		return
	}

	gold = cast.ToInt64(cmds[0].Val())
	goldOfToday = cast.ToInt64(cmds[1].Val())
	goldOfWeek = cast.ToInt64(cmds[2].Val())

	log.Errorw("get rebate gold ", "short", shortId, "cmd len", cmds,
		"gold", gold, "goldOfToday", goldOfToday, "goldOfWeek", goldOfWeek,
		"rebateGoldKey", rebateGoldKey,
		"rebateScoreKeyForToday", rebateScoreKeyForToday,
		"rebateScoreKeyForWeek", rebateScoreKeyForWeek)
	return
}