package mongodb

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestMongo(t *testing.T) {
	if err := Builder().Addr("mongodb://localhost:27017").
		Database("test").
		Connect(); err != nil {
		assert.NoError(t, err)
		return
	}

	if err := Ins.CreateCollection("testcoll"); err != nil {
		assert.NoError(t, err)
		return
	}

	type TestColl struct {
		Id   string `bson:"_id"`
		Name string
		Val  int
	}
	id := primitive.NewObjectID().Hex()
	result, err := Ins.Collection("testColl").InsertOne(context.Background(), &TestColl{
		Id:   id,
		Name: "fegh",
		Val:  999,
	})
	if err != nil {
		assert.NoError(t, err)
		return
	}

	test := &TestColl{}
	err = Ins.Collection("testColl").FindOne(context.Background(), bson.M{"_id": id}).Decode(test)
	if err != nil {
		assert.NoError(t, err)
		return
	}

	fmt.Println(result)
	fmt.Println("find", test)
	_, err = Ins.Collection("testColl").DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		assert.NoError(t, err)
		return
	}
}
