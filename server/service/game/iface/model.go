package iface

import "server/db/table"

type Modeler interface {
	OnLogin()
	OnLogout()
	OnStop()
	Table() table.Tabler
	SetTable(t table.Tabler)
}
