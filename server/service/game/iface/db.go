package iface

import "server/db/table"

type SaveLoader interface {
	Saver
	Loader
}

type Saver interface {
	Save(data ...table.Tabler) error
}

type Loader interface {
	Load(data ...table.Tabler) error
	LoadAll(tableName string, arr interface{}) error
}
