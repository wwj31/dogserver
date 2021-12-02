package iface

import "server/db/table"

type Role interface {
	Modeler
	Table() table.Role
}
