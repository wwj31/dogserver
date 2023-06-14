package rdsop

import (
	"context"
	"math/rand"

	"github.com/spf13/cast"

	"server/common/log"
	"server/common/rds"
)

func AddRoomMgr(mgrId int32) {
	rds.Ins.SAdd(context.Background(), RoomMgrKey(), mgrId, 0)
}

func DelRoomMgr(mgrId int32) {
	rds.Ins.SRem(context.Background(), RoomMgrKey(), mgrId, 0)
}

var cursor uint64

func init() {
	num, err := rds.Ins.SCard(context.Background(), RoomMgrKey()).Result()
	if num == 0 {
		log.Warnw("room mgr num == 0", "err", err)
		return
	}
	cursor = uint64(rand.Intn(int(num + 1)))
}

func GetRoomMgrId() int32 {
	result, cur, err := rds.Ins.SScan(context.Background(), RoomMgrKey(), uint64(cursor), "", 1).Result()
	if err != nil {
		log.Errorw("redis sscan room mgr failed", "cursor", cursor, "err", err)
		return 0
	}
	cursor = cur
	if len(result) > 0 {
		return cast.ToInt32(result[0])
	}
	return 0
}
