package rdsop

import (
	"fmt"
	"github.com/spf13/cast"
)

// LockLoginKey 登录流程用的分布式，防止并发登录，保证单个玩家登录都是同步的
func LockLoginKey(key string) string {
	return "lock:login:" + key
}

// SessionKey 用于查找在线玩家推送消息、以及顶号相关处理
func SessionKey(rid string) string {
	return "session:" + rid
}

// GameNodeKey 记录玩家最近一次进入的game节点
func GameNodeKey(shortId int64) string {
	return "gameNode:" + cast.ToString(shortId)
}

// ShortIDKey 获得并从库里删除一个随机短ID
func ShortIDKey() string {
	return "shortId"
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
	return "joinalliance:" + cast.ToString(shortId)
}
