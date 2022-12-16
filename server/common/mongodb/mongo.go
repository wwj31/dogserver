package mongodb

import (
	"context"
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
