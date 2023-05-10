package rds

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
)

type Builder struct {
	clusterMode    bool
	addr           []string
	userName       string
	password       string
	maxRetries     int
	dialTimeout    time.Duration
	readTimeout    time.Duration
	writeTimeout   time.Duration
	onConnectHande func()
}

func NewBuilder() *Builder {
	return &Builder{
		clusterMode:  false,
		addr:         []string{"localhost:6379"},
		userName:     "",
		password:     "",
		maxRetries:   3,
		dialTimeout:  3 * time.Second,
		readTimeout:  5 * time.Second,
		writeTimeout: 5 * time.Second,
	}
}

func (b *Builder) Connect() (err error) {
	once.Do(func() {
		var client Client
		if b.clusterMode {
			opt := b.clusterOptions()
			client = redis.NewClusterClient(opt)
		} else {
			opt := b.options()
			client = redis.NewClient(opt)
		}

		if b.onConnectHande != nil {
			b.onConnectHande()
		}

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		ping := client.Ping(ctx)
		if ping.Err() != nil {
			err = fmt.Errorf("redis connect ping err:%v", ping.Err())
		}

		Ins = client
		rs = redsync.New(goredis.NewPool(Ins))
	})
	return
}

func (b *Builder) Addr(addr ...string) *Builder {
	b.addr = addr
	return b
}

func (b *Builder) ClusterMode() *Builder {
	b.clusterMode = true
	return b
}

func (b *Builder) UserName(name string) *Builder {
	b.userName = name
	return b
}

func (b *Builder) Password(password string) *Builder {
	b.password = password
	return b
}

func (b *Builder) MaxRetries(maxRetries int) *Builder {
	b.maxRetries = maxRetries
	return b
}

func (b *Builder) DialTimeout(dialTimeout time.Duration) *Builder {
	b.dialTimeout = dialTimeout
	return b
}

func (b *Builder) ReadTimeout(readTimeout time.Duration) *Builder {
	b.readTimeout = readTimeout
	return b
}

func (b *Builder) WriteTimeout(writeTimeout time.Duration) *Builder {
	b.writeTimeout = writeTimeout
	return b
}

func (b *Builder) OnConnect(fn func()) *Builder {
	b.onConnectHande = fn
	return b
}

func (b *Builder) options() *redis.Options {
	return &redis.Options{
		Addr:         b.addr[0],
		Username:     b.userName,
		Password:     b.password,
		MaxRetries:   b.maxRetries,
		DialTimeout:  b.dialTimeout,
		ReadTimeout:  b.readTimeout,
		WriteTimeout: b.writeTimeout,
	}
}

func (b *Builder) clusterOptions() *redis.ClusterOptions {
	return &redis.ClusterOptions{
		Addrs:        b.addr,
		Username:     b.userName,
		Password:     b.password,
		MaxRetries:   b.maxRetries,
		DialTimeout:  b.dialTimeout,
		ReadTimeout:  b.readTimeout,
		WriteTimeout: b.writeTimeout,
	}
}
