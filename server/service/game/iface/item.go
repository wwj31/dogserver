package iface

type Item interface {
	Modeler

	Add(items map[int64]int64, push ...bool)
}
