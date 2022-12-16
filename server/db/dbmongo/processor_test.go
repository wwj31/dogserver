package dbmongo

import (
	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/tools"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"server/common/log"
	"server/common/mongodb"
)

type DocInfo1 struct {
	Id   string `bson:"_id"`
	Val1 string
	Val2 string
}

type DocInfo2 struct {
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

	uids := []string{
		tools.XUID(),
		tools.XUID(),
		tools.XUID(),
		tools.XUID(),
		tools.XUID(),
		tools.XUID(),
		tools.XUID(),
	}

	go func() {
		for {
			uid := uids[rand.Intn(len(uids))]
			docInfo := DocInfo1{
				Id:   uid,
				Val1: "foo" + ":" + cast.ToString(rand.Int()),
				Val2: "bar",
			}
			Store(coll(docInfo), uid, &docInfo)
		}
	}()

	go func() {
		for {
			uid := uids[rand.Intn(len(uids))]
			docInfo := DocInfo2{
				Id:   uid,
				Val1: "foo" + ":" + cast.ToString(rand.Int()),
				Val2: "bar",
			}

			Store(coll(docInfo), uid, &docInfo)
		}
	}()

	time.Sleep(10 * time.Second)
	log.Debugw("stop")
	Stop()
}

func coll(v interface{}) string {
	str := reflect.TypeOf(v).String()
	if str[0] == '*' {
		str = str[1:]
	}
	arr := strings.Split(str, ".")
	return arr[1]
}
