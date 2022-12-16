package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type builder struct {
	addr         string
	databaseName string
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

		Ins.collections = map[string]*mongo.Collection{}
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

		Ins.database = Ins.client.Database(b.databaseName)
	})
	return
}

func (b *builder) Addr(addr string) *builder {
	b.addr = addr
	return b
}

func (b *builder) Database(name string) *builder {
	b.databaseName = name
	return b
}
