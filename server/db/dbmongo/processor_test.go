package dbmongo

import (
	"github.com/stretchr/testify/assert"
	"github.com/wwj31/dogactor/tools"
	"reflect"
	"server/common/mongodb"
	"strings"
	"testing"
)

type DocInfo1 struct {
	Id   string `bson:"_id"`
	Val1 string
	Val2 string
}

func TestProcessor(t *testing.T) {
	if err := mongodb.Builder().Addr("mongodb://localhost:27017").
		Database("test").Connect(); err != nil {
		assert.NoError(t, err)
		return
	}

	uid := tools.XUID()
	docInfo1 := DocInfo1{
		Id:   uid,
		Val1: "foo",
		Val2: "bar",
	}
	str := reflect.TypeOf(docInfo1).String()
	if str[0] == '*' {
		str = str[1:]
	}
	arr := strings.Split(str, ".")

	Store(arr[1], uid, &docInfo1)
	Stop()
}
