package iface

import "server/db/table"

type StoreLoader interface {
	Storer
	Loader
}

type Storer interface {
	Store(insert bool, tablers ...table.Tabler) error
}

type Loader interface {
	Load(data ...table.Tabler) error
	LoadAll(tableName string, arr interface{}) error
}
