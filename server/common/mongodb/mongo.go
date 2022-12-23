package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"server/common/log"
	"sync"
)

var (
	Ins  mongoDB
	once sync.Once
)

type mongoDB struct {
	client       *mongo.Client
	database     *mongo.Database
	addr         string
	databaseName string
	sharding     bool
	collections  map[string]*mongo.Collection
}

func (m *mongoDB) CreateCollection(name string) error {
	err := m.database.CreateCollection(context.Background(), name)
	commandErr, ok := err.(mongo.CommandError)
	if ok && commandErr.HasErrorCode(48) {
		return nil
	}
	if err != nil {
		return err
	}

	if m.sharding {
		result := m.client.Database("admin").RunCommand(context.Background(), bson.D{
			{"shardcollection", m.databaseName + "." + name},
			{"key", bson.M{"_id": "hashed"}},
		})

		if result.Err() != nil {
			log.Errorw("shard collection failed", "err", result.Err())
		}
	}
	return nil
}

func (m *mongoDB) Collection(name string) *mongo.Collection {
	coll := m.collections[name]
	if coll == nil {
		if m.database == nil {
			return nil
		}
		if err := m.CreateCollection(name); err != nil {
			log.Errorw("create collection failed ", "err", err)
			return nil
		}
		coll = m.database.Collection(name)
		m.collections[name] = coll
	}

	return coll
}
