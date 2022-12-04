package redis

import (
	redisv8 "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"sync"
)

var Ins Client

var (
	rs   *redsync.Redsync
	once sync.Once
)

type Client interface {
	redisv8.Scripter
	redisv8.UniversalClient
}
