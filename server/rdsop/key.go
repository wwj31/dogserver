package rdsop

import "fmt"

// LockLoginKey 登录流程用的分布式，防止并发登录，保证单个玩家登录都是同步的
func LockLoginKey(key string) string {
	return "lock:login:" + key
}

// SessionKey 仅用于顶号处理，踢掉旧链接
func SessionKey(rid string) string {
	return "session:" + rid
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
