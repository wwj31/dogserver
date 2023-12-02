package rdsop

import (
	"fmt"

	"github.com/wwj31/dogactor/tools"
)

// LockLoginKey 登录流程用的分布式，防止并发登录，保证单个玩家登录都是同步的
func LockLoginKey(key string) string {
	return fmt.Sprintf("lock:login:%v", key)
}

// SessionKey 用于查找在线玩家推送消息、以及顶号相关处理
func SessionKey(rid string) string {
	return fmt.Sprintf("session:%v", rid)
}

// GameNodeKey 记录玩家最近一次进入的game节点
func GameNodeKey(shortId int64) string {
	return fmt.Sprintf("gamenode:%v", shortId)
}

// ShortIDKey 获得并从库里删除一个随机短ID
func ShortIDKey() string {
	return "shortid"
}

// PlayerInfoKey 玩家基础公共信息key
func PlayerInfoKey(shortId int64) string {
	return fmt.Sprintf("playerinfo:%v", shortId)
}

// PlayerDailyStatKey 玩家每日游戏数据信息key
func PlayerDailyStatKey(shortId int64) string {
	return fmt.Sprintf("playerdailystat:%v:%v", shortId, tools.Now().Local().Format(tools.StdDateFormat))
}

// AgentUpKey 玩家的上级代理
func AgentUpKey(shortId int64) string {
	return fmt.Sprintf("agent:%v:up", shortId)
}

// AgentDownKey 玩家的下级代理
func AgentDownKey(shortId int64) string {
	return fmt.Sprintf("agent:%v:down", shortId)
}

// AgentRebateKey 抽水返利点位信息
func AgentRebateKey(shortId int64) string {
	return fmt.Sprintf("agent:%v:rebate", shortId)
}

// RebateGoldKey 返利利润
func RebateGoldKey(shortId int64) string {
	return fmt.Sprintf("rebate:%v:gold", shortId)
}

// RebateScoreKeyForToday 统计今日返利利润
func RebateScoreKeyForToday(shortId int64) string {
	return fmt.Sprintf("rebate:%v:gold_for_day:%v", shortId, tools.Now().Local().Format(tools.StdDateFormat))
}

// RebateScoreKeyForYesterday 统计昨日返利利润
func RebateScoreKeyForYesterday(shortId int64) string {
	return fmt.Sprintf("rebate:%v:gold_for_day:%v", shortId, tools.Now().Add(-tools.Day).Local().Format(tools.StdDateFormat))
}

// RebateScoreKeyForWeek 统计本周返利利润
func RebateScoreKeyForWeek(shortId int64) string {
	weekStart := tools.NewTimeEx(tools.Now()).StartOfWeek()
	return fmt.Sprintf("rebate:%v:gold_for_week:%v", shortId, weekStart.Local().Format(tools.StdDateFormat))
}

// RebateScoreKeyForDetail 每笔返利记录信息
func RebateScoreKeyForDetail(shortId int64, date string) string {
	return fmt.Sprintf("rebate:%v:gold_for_detail:%v", shortId, date)
}

// ContributeScoreKeyForToday 今日业绩
func ContributeScoreKeyForToday(shortId int64) string {
	return fmt.Sprintf("contribute:%v:gold_for_day:%v", shortId, tools.Now().Local().Format(tools.StdDateFormat))
}
func ContributeScoreKeyForYesterday(shortId int64) string {
	return fmt.Sprintf("contribute:%v:gold_for_day:%v", shortId, tools.Now().Local().Add(-tools.Day).Format(tools.StdDateFormat))
}

// ContributeScoreKeyForWeek 本周业绩
func ContributeScoreKeyForWeek(shortId int64) string {
	weekStart := tools.NewTimeEx(tools.Now()).StartOfWeek()
	return fmt.Sprintf("contribute:%v:gold_for_week:%v", shortId, weekStart.Local().Format(tools.StdDateFormat))
}

// ContributeScoreKey 总业绩
func ContributeScoreKey(shortId int64) string {
	return fmt.Sprintf("contribute:%v", shortId)
}

// DeleteAlliancesKey 被删除的联盟bi
func DeleteAlliancesKey() string {
	return fmt.Sprintf("alliance:deleted")
}

// JoinAllianceKey 不在线的玩家，记录进入联盟id
func JoinAllianceKey(shortId int64) string {
	return fmt.Sprintf("alliance:join:%v", shortId)
}

// RoomMgrSetKey 所有房间管理器节点
func RoomMgrSetKey() string {
	return "room:manager:list"
}

// RoomsSetKey 所有房间
func RoomsSetKey(allianceId int32) string {
	return fmt.Sprintf("room:list:%v", allianceId)
}

// RoomsInfoKey 房间信息
func RoomsInfoKey(roomId int64) string {
	return fmt.Sprintf("room:info:%v", roomId)
}

// RoomsRecordingDataKey 房间回放信息
func RoomsRecordingDataKey(roomId int64) string {
	return fmt.Sprintf("room:recording:%v", roomId)
}

// RoomsIncIdKey 房间递增ID
func RoomsIncIdKey() string {
	return fmt.Sprintf("room:inc_id")
}

// GMCmdListKey 离线玩家gm命令队列
func GMCmdListKey(shortId int64) string {
	return fmt.Sprintf("gm:cmd:%v", shortId)
}

// UpdateGoldRecordKey 玩家金币变更记录队列
func UpdateGoldRecordKey(shortId int64) string {
	return fmt.Sprintf("goldrecords:%v", shortId)
}
