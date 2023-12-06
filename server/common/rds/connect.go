package rds

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"

	"github.com/redis/go-redis/v9"
)

func Connect(uri string, cluster bool) (err error) {
	once.Do(func() {
		var client Client
		if cluster {
			opt, _ := redis.ParseClusterURL(uri)
			client = redis.NewClusterClient(opt)
		} else {
			opt, _ := redis.ParseURL(uri)
			client = redis.NewClient(opt)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		ping := client.Ping(ctx)
		if ping.Err() != nil {
			err = fmt.Errorf("redis connect ping err:%v uri:%v", ping.Err(), uri)
		}

		Ins = client
		rs = redsync.New(goredis.NewPool(Ins))
	})
	return
}
