package account

import "server/service/db/table"

type AccountMgr struct {
	accounts map[uint64]*table.Account
}
