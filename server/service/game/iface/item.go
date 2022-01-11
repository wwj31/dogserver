package iface

import "server/proto/message"

type Item interface {
	Modeler

	Enough(items map[int64]int64) bool
	Use(items map[int64]int64) message.ERROR
	Add(items map[int64]int64, push ...bool)
}
