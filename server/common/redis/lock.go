package redis

import (
	"github.com/go-redsync/redsync/v4"
	"server/common/log"
)

func Locker(key string) *redsync.Mutex {
	return rs.NewMutex(key)
}

func LockDo(key string, fn func()) {
	locker := rs.NewMutex(key)
	if err := locker.Lock(); err != nil {
		log.Errorw("redsync lock failed", "key", key, "err", err)
		return
	}
	defer locker.Unlock()

	fn()
}
