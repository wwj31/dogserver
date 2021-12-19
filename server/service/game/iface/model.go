package iface

import "server/db/table"

type Modeler interface {
	OnLogin()
	OnLogout()
	Table() table.Tabler
	SetTable(t table.Tabler)
}
