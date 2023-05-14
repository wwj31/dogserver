package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
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

func (m *mongoDB) CreateIndex(name string, tagStruct interface{}) error {
	coll := m.Collection(name)
	var models []mongo.IndexModel

	t := reflect.TypeOf(tagStruct)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return fmt.Errorf("tagStruct is not a struct kind:%v", t.Kind())
	}

	numFields := t.NumField()

	for i := 0; i < numFields; i++ {
		field := t.Field(i)
		tag := field.Tag.Get("index")

		if tag == "true" {
			bsonName := field.Tag.Get("bson")
			if bsonName == "" {
				bsonName = field.Name
			}

			keys := bson.D{{Key: bsonName, Value: 1}}
			indexModel := mongo.IndexModel{Keys: keys}
			models = append(models, indexModel)
		}
	}

	_, err := coll.Indexes().CreateMany(context.Background(), models)
	if err != nil {
		return err
	}

	return nil
}
