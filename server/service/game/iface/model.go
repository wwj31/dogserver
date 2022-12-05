package iface

import (
	"server/db/dbmysql/table"
)

type Modeler interface {
	OnSave(data *table.Player)
	OnLogin()
	OnLogout()
}
