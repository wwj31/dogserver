package rdsop

import (
	"fmt"
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

// PlayerInfoKey 玩家基础公共信息
func PlayerInfoKey(shortId int64) string {
	return fmt.Sprintf("playerinfo:%v", shortId)
}

// AgentUpKey 玩家的上级代理
func AgentUpKey(shortId int64) string {
	return fmt.Sprintf("agent:%v:up", shortId)
}

// AgentDownKey 玩家的下级代理
func AgentDownKey(shortId int64) string {
	return fmt.Sprintf("agent:%v:down", shortId)
}

// DeleteAlliancesKey 被删除的联盟
func DeleteAlliancesKey() string {
	return fmt.Sprintf("deleted:alliance")
}

// JoinAllianceKey 不在线的玩家，记录进入联盟id
func JoinAllianceKey(shortId int64) string {
	return fmt.Sprintf("joinalliance:%v", shortId)
}

// RoomMgrKey 所有房间管理器节点
func RoomMgrKey() string {
	return "roommgr"
}

// RoomsKey 所有房间
func RoomsKey(allianceId int32) string {
	return fmt.Sprintf("room:%v", allianceId)
}
