package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"server/common/log"
)

type builder struct {
	addr         string
	databaseName string
	sharding     bool
}

func Builder() *builder {
	return &builder{}
}

func (b *builder) Connect() (err error) {
	once.Do(func() {
		Ins.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(b.addr))
		if err != nil {
			return
		}

		if b.databaseName == "" {
			err = fmt.Errorf("mongo database is nil")
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err = Ins.client.Ping(ctx, readpref.Primary()); err != nil {
			err = fmt.Errorf("mongo connect failed err:%v", err)
			return
		}

		Ins.databaseName = b.databaseName
		Ins.database = Ins.client.Database(b.databaseName)

		Ins.sharding = b.sharding
		if b.sharding {
			result := Ins.client.Database("admin").RunCommand(context.Background(), bson.M{"enablesharding": b.databaseName})
			if result.Err() != nil {
				log.Warnw("enable sharding failed", "err", result.Err())
			}
		}
	})
	return
}

func (b *builder) EnableSharding() *builder {
	b.sharding = true
	return b
}

func (b *builder) Addr(addr string) *builder {
	b.addr = addr
	return b
}

func (b *builder) Database(name string) *builder {
	b.databaseName = name
	return b
}
