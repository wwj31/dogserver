package rdsop

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/tools"

	"server/common"
	"server/common/log"
	"server/common/rds"
	"server/proto/outermsg/outer"
)

// 统计玩家的各种返利信息，以及返利相关的操作

type RebateInfo struct {
	Point      int32           `json:"point"`      // 自己的返利点
	DownPoints map[int64]int32 `json:"downPoints"` // 给下级分的返利点
}

// GetRebateInfo 获得玩家返利信息
func GetRebateInfo(shortId int64) RebateInfo {
	str, _ := rds.Ins.Get(context.Background(), AgentRebateKey(shortId)).Result()
	var result RebateInfo
	if str == "" {
		result.DownPoints = map[int64]int32{}
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
		rebateInfo := GetRebateInfo(shortId)
		if rebateInfo.Point < point {
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
			rebateInfo.DownPoints[downShortId] = point

			pip := rds.Ins.Pipeline()
			ctx := context.Background()
			// 更新下级点位
			pip.Set(ctx, AgentRebateKey(downShortId), common.JsonMarshal(downRebateInfo), -1)
			// 更新自己管理的下级
			pip.Set(ctx, AgentRebateKey(shortId), common.JsonMarshal(rebateInfo), -1)
			_, er := pip.Exec(ctx)
			log.Infow("set rebate info success", "short", shortId, "downShort", downShortId,
				"points", rebateInfo.DownPoints, "err", er)
		})
	})

	return
}

func IncRebateGold(shortId, score int64, pip redis.Pipeliner) {
	pip.IncrBy(context.Background(), RebateGoldKey(shortId), score)
}

// RecordRebateGold 给玩家加返利分数
func RecordRebateGold(info string, shortId, score int64, pip redis.Pipeliner) {
	ctx := context.Background()
	IncRebateGold(shortId, score, pip)

	// 统计今日返利
	statTodayKey := RebateScoreKeyForToday(shortId)
	pip.IncrBy(ctx, statTodayKey, score)
	pip.Expire(ctx, statTodayKey, tools.Day)

	// 统计本周返利
	statWeekKey := RebateScoreKeyForWeek(shortId)
	pip.IncrBy(ctx, statWeekKey, score)
	pip.Expire(ctx, statWeekKey, 7*tools.Day)

	// 统计每笔返利详情
	statDetailKey := RebateScoreKeyForDetail(shortId, tools.Now().Local().Format(tools.StdDateFormat))
	pip.LPush(ctx, statDetailKey, info)
	pip.Expire(ctx, statDetailKey, 11*tools.Day)
}

// GetRebateRecordOf3Day 获得玩家3天的返利记录详情
func GetRebateRecordOf3Day(shortId int64) (records []*outer.RebateDetailInfo) {
	ctx := context.Background()
	pip := rds.Ins.Pipeline()
	var lists []*redis.StringSliceCmd
	for i := 0; i < 10; i++ {
		day := tools.Now().Local().Add(-time.Duration(i) * tools.Day).Format(tools.StdDateFormat)
		statDetailKey := RebateScoreKeyForDetail(shortId, day)
		lists = append(lists, pip.LRange(ctx, statDetailKey, 0, -1))
	}
	_, err := pip.Exec(ctx)
	if err != nil {
		log.Errorw("GetRebateRecordOf3Day redis pip failed", "err", err, "shortId", shortId)
		return
	}

	for _, list := range lists {
		for _, info := range list.Val() {
			v := &outer.RebateDetailInfo{}
			e := json.Unmarshal([]byte(info), v)
			if e != nil {
				log.Warnw("json unmarshal failed", "info", info)
				continue
			}
			records = append(records, v)
		}
	}
	return
}

// GetRebateGold 获得玩家 当前/今日/昨日/本周 的返利总分
func GetRebateGold(shortId int64) (gold, goldOfToday, goldOfYesterday, goldOfWeek int64) {
	pip := rds.Ins.Pipeline()
	ctx := context.Background()

	var (
		cmds                       []*redis.StringCmd
		rebateGoldKey              = RebateGoldKey(shortId)
		rebateScoreKeyForToday     = RebateScoreKeyForToday(shortId)
		rebateScoreKeyForYesterday = RebateScoreKeyForYesterday(shortId)
		rebateScoreKeyForWeek      = RebateScoreKeyForWeek(shortId)
	)

	cmds = append(cmds, pip.Get(ctx, rebateGoldKey))
	cmds = append(cmds, pip.Get(ctx, rebateScoreKeyForToday))
	cmds = append(cmds, pip.Get(ctx, rebateScoreKeyForYesterday))
	cmds = append(cmds, pip.Get(ctx, rebateScoreKeyForWeek))
	pip.Exec(ctx)

	if len(cmds) < 3 {
		log.Errorw("get rebate gold unexpected outcome", "short", shortId, "cmd len", cmds,
			"rebateGoldKey", rebateGoldKey,
			"rebateScoreKeyForToday", rebateScoreKeyForToday,
			"rebateScoreKeyForWeek", rebateScoreKeyForWeek)
		return
	}

	gold = cast.ToInt64(cmds[0].Val())
	goldOfToday = cast.ToInt64(cmds[1].Val())
	goldOfYesterday = cast.ToInt64(cmds[2].Val())
	goldOfWeek = cast.ToInt64(cmds[3].Val())

	log.Infow("get rebate gold ", "short", shortId, "cmd len", cmds,
		"gold", gold,
		"goldOfToday", goldOfToday,
		"goldOfYesterday", goldOfYesterday,
		"goldOfWeek", goldOfWeek,
		"rebateGoldKey", rebateGoldKey,
		"rebateScoreKeyForToday", rebateScoreKeyForToday,
		"rebateScoreKeyForWeek", rebateScoreKeyForWeek)
	return
}
