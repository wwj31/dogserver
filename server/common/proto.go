package common

import (
	"server/common/log"

	"github.com/gogo/protobuf/proto"
)

func ProtoMarshal(msg proto.Message) []byte {
	bytes, err := proto.Marshal(msg)
	if err != nil {
		log.Errorw("proto marshal error", "err", err)
		return nil
	}
	return bytes
}
