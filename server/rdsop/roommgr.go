package rdsop

import (
	"context"
	"math/rand"
	"sync"

	"github.com/spf13/cast"

	"server/common/log"
	"server/common/rds"
)

func AddRoomMgr(mgrId int32) {
	rds.Ins.SAdd(context.Background(), RoomMgrSetKey(), mgrId)
}

func DelRoomMgr(mgrId int32) {
	rds.Ins.SRem(context.Background(), RoomMgrSetKey(), mgrId)
}

var (
	cursor     int64 = -1
	cursorLock sync.Mutex
)

func GetRoomMgrId() int32 {
	cursorLock.Lock()
	defer cursorLock.Unlock()

	if cursor == -1 {
		num, err := rds.Ins.SCard(context.Background(), RoomMgrSetKey()).Result()
		if num == 0 {
			log.Warnw("room mgr num == 0", "err", err)
			return -1
		}
		cursor = int64(rand.Intn(int(num)))
	}

	result, cur, err := rds.Ins.SScan(context.Background(), RoomMgrSetKey(), uint64(cursor), "", 1).Result()
	if err != nil {
		log.Errorw("redis sscan room mgr failed", "cursor", cursor, "err", err)
		return -1
	}

	cursor = int64(cur)
	if len(result) > 0 {
		return cast.ToInt32(result[0])
	}

	log.Errorw("can not find room mgr", "cursor", cursor)
	return -1
}
