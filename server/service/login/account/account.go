package account

import (
	"fmt"
	"server/common"
	"server/db/table"
)

type Account struct {
	table.Account
	serverId common.ActorId
	gSession common.GSession
}

func (s *Account) ServerId() common.ActorId       { return s.serverId }
func (s *Account) GSession() common.GSession      { return s.gSession }
func (s *Account) SetgSession(gs common.GSession) { s.gSession = gs }

func combine(a, b string) string {
	return fmt.Sprintf("%v_%v", a, b)
}
