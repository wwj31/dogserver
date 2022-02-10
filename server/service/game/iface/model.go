package iface

import "server/db/table"

type Modeler interface {
	OnSave(data *table.Player)
	OnLogin()
	OnLogout()
}
