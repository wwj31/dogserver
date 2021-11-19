package iface

import "server/service/db/table"

type SaveLoader interface {
	Saver
	Loader
}

type Saver interface {
	Save(data table.Tabler) error
}

type Loader interface {
	Load(data table.Tabler) error
}
