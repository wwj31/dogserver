package account

import (
	"fmt"
	"server/common"
	"server/common/actortype"
	"server/db/table"
)

type Account struct {
	table    table.Account
	serverId actortype.ActorId
	gSession common.GSession
}

func (s *Account) UUId() uint64                   { return s.table.UUId }
func (s *Account) LastRoleId() uint64             { return s.table.LastRoleId }
func (s *Account) ServerId() actortype.ActorId    { return s.serverId }
func (s *Account) GSession() common.GSession      { return s.gSession }
func (s *Account) SetGSession(gs common.GSession) { s.gSession = gs }

func combine(a, b string) string {
	return fmt.Sprintf("%v_%v", a, b)
}
