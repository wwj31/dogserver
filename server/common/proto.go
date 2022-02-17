package common

import (
	"reflect"
	"server/common/log"

	gogo "github.com/gogo/protobuf/proto"
)

func ProtoMarshal(msg gogo.Message) []byte {
	bytes, err := gogo.Marshal(msg)
	if err != nil {
		log.Errorw("proto marshal error", "err", err)
		return nil
	}
	return bytes
}

func ProtoType(msg interface{}) string {
	str := reflect.TypeOf(msg).String()
	if str[0] == '*' {
		str = str[1:]
	}
	return str
}
