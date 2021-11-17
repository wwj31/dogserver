package account

import (
	"fmt"
	"server/common"
	"server/proto/message"
	"server/service/db/table"
	"time"
)

type AccountMgr struct {
	accountsByUId        map[uint64]*Account
	accountsByPlatformId map[string]*Account

	uuidGen *common.UID
}

func NewAccountMgr() *AccountMgr {
	return &AccountMgr{
		accountsByUId:        make(map[uint64]*Account),
		accountsByPlatformId: make(map[string]*Account),
		uuidGen:              common.NewUID(),
	}
}

func (s *AccountMgr) Login(msg *message.LoginReq) (acc *Account, new bool) {
	platformId := combine(msg.PlatformName, msg.PlatformUUID)
	if acc = s.accountsByPlatformId[platformId]; acc != nil {
		return acc, false
	}

	// newAccount
	newAcc := &Account{}
	newAcc.UUId = s.uuidGen.Uuid()
	newAcc.PlatformUUId = platformId
	newAcc.OS = msg.OS
	newAcc.ClientVersion = msg.ClientVersion
	newAcc.Roles = make(table.RoleMap)

	// newRole
	newRole := &table.Role{
		UUId:     newAcc.UUId,
		RoleId:   s.uuidGen.Uuid(),
		SId:      "S1",
		Name:     fmt.Sprintf("player_%v", newAcc.UUId),
		Icon:     "",
		Country:  "",
		IsDelete: false,
		LoginAt:  time.Now().Unix(),
		LogoutAt: 0,
	}
	newAcc.Roles[newRole.RoleId] = newRole
	newAcc.LastRoleId = newRole.RoleId

	s.accountsByPlatformId[platformId] = newAcc
	s.accountsByUId[newAcc.UUId] = newAcc

	// todo ... 回存db
	return newAcc, true
}
