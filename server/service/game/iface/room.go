package iface

import (
	"server/proto/innermsg/inner"
)

type Room interface {
	Modeler
	RoomId() int64
	SetRoomInfo(info *inner.RoomInfo)
}
