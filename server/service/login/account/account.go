package account

import (
	"fmt"
	"server/service/db/table"
)

type account struct {
	table.Account
}

func NewAccount(platformName, platformUUID string) *account {
	acc := &account{}
	acc.PlatformUUId = combine(platformName, platformUUID)
	return acc
}

func combine(a, b string) string {
	return fmt.Sprintf("%v_%v", a, b)
}
