package iface

type Item interface {
	Modeler

	Enough(items map[int64]int64) bool
	Add(items map[int64]int64, push ...bool)
}
