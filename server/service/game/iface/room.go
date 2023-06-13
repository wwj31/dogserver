package iface

import (
	"server/proto/innermsg/inner"
)

type Room interface {
	Modeler
	RoomId() int32
	SetRoomInfo(info *inner.RoomInfo)
}
