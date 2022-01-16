package iface

import (
	"server/proto/outermsg/outer"
)

type Item interface {
	Modeler

	Enough(items map[int64]int64) bool
	Use(items map[int64]int64) outer.ERROR
	Add(items map[int64]int64, push ...bool)
}
