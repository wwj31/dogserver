package iface

import "server/db/table"

type Role interface {
	Table() table.Role
}
