package redis

import (
	redisv8 "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"time"
)

type builder struct {
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

func Builder() *builder {
	return &builder{
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

func (b *builder) OK() {
	once.Do(func() {
		var client Client
		if b.clusterMode {
			opt := b.clusterOptions()
			client = redisv8.NewClusterClient(opt)
		} else {
			opt := b.options()
			client = redisv8.NewClient(opt)
		}

		if b.onConnectHande != nil {
			b.onConnectHande()
		}

		Ins = client
		rs = redsync.New(goredis.NewPool(Ins))
	})
}

func (b *builder) Addr(addr ...string) *builder {
	b.addr = addr
	return b
}

func (b *builder) ClusterMode() *builder {
	b.clusterMode = true
	return b
}

func (b *builder) UserName(name string) *builder {
	b.userName = name
	return b
}

func (b *builder) Password(password string) *builder {
	b.password = password
	return b
}

func (b *builder) MaxRetries(maxRetries int) *builder {
	b.maxRetries = maxRetries
	return b
}

func (b *builder) DialTimeout(dialTimeout time.Duration) *builder {
	b.dialTimeout = dialTimeout
	return b
}

func (b *builder) ReadTimeout(readTimeout time.Duration) *builder {
	b.readTimeout = readTimeout
	return b
}

func (b *builder) WriteTimeout(writeTimeout time.Duration) *builder {
	b.writeTimeout = writeTimeout
	return b
}

func (b *builder) OnConnect(fn func()) *builder {
	b.onConnectHande = fn
	return b
}

func (b *builder) options() *redisv8.Options {
	return &redisv8.Options{
		Addr:         b.addr[0],
		Username:     b.userName,
		Password:     b.password,
		MaxRetries:   b.maxRetries,
		DialTimeout:  b.dialTimeout,
		ReadTimeout:  b.readTimeout,
		WriteTimeout: b.writeTimeout,
	}
}

func (b *builder) clusterOptions() *redisv8.ClusterOptions {
	return &redisv8.ClusterOptions{
		Addrs:        b.addr,
		Username:     b.userName,
		Password:     b.password,
		MaxRetries:   b.maxRetries,
		DialTimeout:  b.dialTimeout,
		ReadTimeout:  b.readTimeout,
		WriteTimeout: b.writeTimeout,
	}
}
