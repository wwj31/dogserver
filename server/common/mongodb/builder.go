package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		Ins.client, err = mongo.Connect(ctx, options.Client().ApplyURI(b.addr))
		if err != nil {
			return
		}

		Ins.collections = map[string]*mongo.Collection{}
		if b.databaseName == "" {
			err = fmt.Errorf("mongo database is nil")
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
