package common

import (
	"reflect"
	"server/common/log"
	"server/proto/outermsg/outer"

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

func ProtoUnmarshal(msgType string, bytes []byte) gogo.Message {
	v, ok := outer.Spawner(msgType)
	if !ok {
		log.Errorw("ProtoUnmarshal outer msg no found", "msgType", msgType)
		return nil
	}
	msg := v.(gogo.Message)
	err := gogo.Unmarshal(bytes, msg)
	if err != nil {
		log.Errorw("proto marshal error", "err", err)
		return nil
	}
	return msg
}

func ProtoType(msg interface{}) string {
	typ := reflect.TypeOf(msg)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	return typ.String()
}
