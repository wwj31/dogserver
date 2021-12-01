package account

import (
	"fmt"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/log"
	"server/common"
	"server/db/table"
	"server/proto/message"
	"server/service/game/iface"
	"time"
)

type AccountMgr struct {
	accountsByUId        map[uint64]*Account
	accountsByPlatformId map[string]*Account
	accountsByName       map[string]*Account

	uuidGen *common.UID
}

func NewAccountMgr() *AccountMgr {
	return &AccountMgr{
		accountsByUId:        make(map[uint64]*Account),
		accountsByPlatformId: make(map[string]*Account),
		accountsByName:       make(map[string]*Account),
		uuidGen:              common.NewUID(1),
	}
}

func (s *AccountMgr) LoadAllAccount(loader iface.Loader) {
	var all []table.Account
	err := loader.LoadAll((&table.Account{}).TableName(), &all)
	expect.Nil(err)

	for _, data := range all {
		acc := &Account{Account: data}
		s.accountsByPlatformId[acc.PlatformUUId] = acc
		s.accountsByUId[acc.UUId] = acc
		lastLoginRole := acc.Roles[acc.LastRoleId]
		if lastLoginRole != nil {
			s.accountsByName[lastLoginRole.Name] = acc
			acc.serverId = common.GameName(int32(lastLoginRole.SId))
		}
	}
}

func (s *AccountMgr) Login(msg *message.LoginReq, saver iface.Saver) (acc *Account, new bool) {
	platformId := combine(msg.PlatformName, msg.PlatformUUID)
	if acc = s.accountsByPlatformId[platformId]; acc != nil {
		return acc, false
	}

	// newAccount
	newAcc := &Account{}
	newAcc.UUId = s.uuidGen.Uuid()
	newAcc.PlatformUUId = platformId
	newAcc.OS = msg.OS
	newAcc.ClientVersion = msg.ClientVersion // newRole
	newRole := &table.Role{
		UUId:     newAcc.UUId,
		RoleId:   s.uuidGen.Uuid(),
		SId:      1,
		Name:     fmt.Sprintf("player_%v", newAcc.UUId),
		Icon:     "Avatar",
		Country:  "中国",
		IsDelete: false,
		CreateAt: time.Now().Unix(),
		LoginAt:  time.Now().Unix(),
		LogoutAt: 0,
	}
	newAcc.Roles = table.RoleMap{newRole.RoleId: newRole}
	newAcc.LastRoleId = newRole.RoleId
	newAcc.serverId = common.GameName(int32(newRole.SId))

	// 回存db
	if err := saver.Save(&newAcc.Account, newRole); err != nil {
		log.KV("err", err).Error("save err")
		return nil, false
	}

	s.accountsByPlatformId[platformId] = newAcc
	s.accountsByUId[newAcc.UUId] = newAcc

	return newAcc, true
}
