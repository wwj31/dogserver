package rds

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/redis/go-redis/v9"

	"sync"
)

var Ins Client

var (
	rs   *redsync.Redsync
	once sync.Once
)

type Client interface {
	redis.Scripter
	redis.UniversalClient
}
