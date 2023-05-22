package rdsop

import (
	"context"
	"encoding/json"

	"github.com/spf13/cast"

	"server/common/log"
	"server/common/rds"
	"server/proto/innermsg/inner"
)

func SetPlayerInfo(info *inner.PlayerInfo) {
	b, err := json.Marshal(info)
	if err != nil {
		log.Errorw("SetPlayerInfo json marshal failed", "err", err, "info", info.String())
		return
	}

	rds.Ins.Set(context.Background(), PlayerInfoKey(info.ShortId), string(b), 0)
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

// SetAgentUp 设置上级
func SetAgentUp(shortId, up int64) {
	rds.Ins.Set(context.Background(), AgentUpKey(shortId), cast.ToString(up), 0)
}

// AgentUp 获得上级
func AgentUp(shortId int64) (up int64) {
	str, _ := rds.Ins.Get(context.Background(), AgentUpKey(shortId)).Result()
	return cast.ToInt64(str)
}

// AddAgentDown 添加下级
func AddAgentDown(shortId int64, down ...interface{}) {
	if len(down) == 0 {
		return
	}
	rds.Ins.SAdd(context.Background(), AgentDownKey(shortId), down...)
}

// AgentDownAll 获取全部下级
func AgentDownAll(shortId int64) (up []int64) {
	all, _ := rds.Ins.SMembers(context.Background(), AgentUpKey(shortId)).Result()
	for _, str := range all {
		up = append(up, cast.ToInt64(str))
	}
	return up
}

// ExistAgentDown 判断是否在自己下级中
func ExistAgentDown(shortId, down int64) bool {
	exist, _ := rds.Ins.SIsMember(context.Background(), AgentUpKey(shortId), down).Result()
	return exist
}
