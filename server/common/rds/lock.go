package rds

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

// Lock For example:
//
//		func(){
//		 	defer Lock("lockXXX")()
//	       do something...
//		}
func Lock(key string) func() {
	lock := Locker(key)
	return func() { _, _ = lock.Unlock() }
}
