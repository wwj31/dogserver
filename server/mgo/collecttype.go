package mgo

import (
	"reflect"
	"strings"

	gogo "github.com/gogo/protobuf/proto"

	"server/common"
	"server/common/log"
)

func GoGoCollectionType(message gogo.Message) string {
	str := strings.Split(common.ProtoType(message), ".")
	if len(str) < 2 {
		log.Errorw("msg name get failed", "type", reflect.TypeOf(message).String())
		return ""
	}

	if reflect.ValueOf(message).IsNil() {
		log.Errorw("doc is nil interface{}", "type", str[1])
		return ""
	}
	return strings.ToLower(str[1])
}
