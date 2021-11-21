package account

import (
	"fmt"
	"server/common"
	"server/db/table"
)

type Account struct {
	table.Account
	ServerId common.ActorId
	GSession string
}

func combine(a, b string) string {
	return fmt.Sprintf("%v_%v", a, b)
}
