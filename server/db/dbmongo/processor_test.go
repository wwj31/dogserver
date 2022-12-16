package dbmongo

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wwj31/dogactor/tools"
	"server/common/log"
	"server/common/mongodb"
)

type DocInfo1 struct {
	Id   string `bson:"_id"`
	Val1 string
	Val2 string
}

func TestProcessor(t *testing.T) {
	log.Init(-1, "", "", true)
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
