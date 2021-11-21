package account

import (
	"fmt"
	"server/db/table"
)

type Account struct {
	table.Account
}

func combine(a, b string) string {
	return fmt.Sprintf("%v_%v", a, b)
}
