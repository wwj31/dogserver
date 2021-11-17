package account

import (
	"fmt"
	"server/service/db/table"
)

type Account struct {
	table.Account
}

func combine(a, b string) string {
	return fmt.Sprintf("%v_%v", a, b)
}
