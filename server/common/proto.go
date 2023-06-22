package common

import (
	"reflect"

	"github.com/golang/protobuf/proto"
	"server/common/log"
	"server/proto/outermsg/outer"
)

func ProtoMarshal(msg proto.Message) []byte {
	bytes, err := proto.Marshal(msg)
	if err != nil {
		log.Errorw("proto marshal error", "err", err)
		return nil
	}
	return bytes
}

func ProtoUnmarshal(msgType string, bytes []byte) proto.Message {
	v, ok := outer.Spawner(msgType)
	if !ok {
		log.Errorw("ProtoUnmarshal outer msg no found", "msgType", msgType)
		return nil
	}
	msg := v.(proto.Message)
	err := proto.Unmarshal(bytes, msg)
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
