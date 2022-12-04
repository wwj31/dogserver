package mongodb

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestMongo(t *testing.T) {
	if err := Builder().Addr("mongodb://localhost:27017").Database("test").
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
	result, err := Ins.Collection("testColl").InsertOne(context.Background(), &TestColl{
		Id:   primitive.NewObjectID().Hex(),
		Name: "fegh",
		Val:  999,
	})

	if err != nil {
		assert.NoError(t, err)
		return
	}
	fmt.Println(result)
}
