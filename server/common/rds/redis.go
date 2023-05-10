package rds

import (
	redisv9 "github.com/go-redis/redis/v9"
	"github.com/go-redsync/redsync/v4"
	"sync"
)

var Ins Client

var (
	rs   *redsync.Redsync
	once sync.Once
)

type Client interface {
	redisv9.Scripter
	redisv9.UniversalClient
}
