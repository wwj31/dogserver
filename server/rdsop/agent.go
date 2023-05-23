package rdsop

import (
	"context"

	"github.com/spf13/cast"

	"server/common/rds"
)

// SetAgentUp 设置上级
func SetAgentUp(shortId, up int64) {
	rds.Ins.Set(context.Background(), AgentUpKey(shortId), cast.ToString(up), 0)
}

// AgentUp 获得上级
func AgentUp(shortId int64) (up int64) {
	str, _ := rds.Ins.Get(context.Background(), AgentUpKey(shortId)).Result()
	return cast.ToInt64(str)
}

// AgentUpAll 获取所有上级,结果的头部是上一级，尾部是顶级
func AgentUpAll(shortId int64) (upAll []int64) {
	for shortId != 0 {
		upId := AgentUp(shortId)
		upAll = append(upAll, upId)
		shortId = upId
	}

	return
}

// AddAgentDown 添加下级
func AddAgentDown(shortId int64, down ...interface{}) {
	if len(down) == 0 {
		return
	}
	rds.Ins.SAdd(context.Background(), AgentDownKey(shortId), down...)
}

// AgentDown 获得下级 downNum 获取至第几层级，不填表示全部获取
func AgentDown(shortId int64, downNum ...int) (down []int64) {
	var (
		downLv int
		next   int
		ids    []int64
	)

	if len(downNum) > 0 {
		downLv = downNum[0]
	}

	ids = append(ids, shortId)
	for downLv == 0 || next >= downLv {
		var tmpIds []int64 // 当前层级的所有下级
		for _, id := range ids {
			all, _ := rds.Ins.SMembers(context.Background(), AgentUpKey(id)).Result()
			for _, str := range all {
				downId := cast.ToInt64(str)
				tmpIds = append(tmpIds, downId)
				down = append(down, downId)
			}
		}
		ids = tmpIds

		next++
	}
	return down
}

// ExistAgentDown 判断是否在自己下级中
func ExistAgentDown(shortId, down int64) bool {
	exist, _ := rds.Ins.SIsMember(context.Background(), AgentUpKey(shortId), down).Result()
	return exist
}