package account

import (
	"fmt"
	"github.com/spf13/cast"
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

	uuidGen common.UID
}

func NewAccountMgr() AccountMgr {
	return AccountMgr{
		accountsByUId:        make(map[uint64]*Account),
		accountsByPlatformId: make(map[string]*Account),
		accountsByName:       make(map[string]*Account),
		uuidGen:              common.NewUID(1),
	}
}

func (s *AccountMgr) LoadAllAccount(loader iface.Loader) {
	var all []table.Account
	tb := &table.Account{}
	if tb.Count() > 1 {
		for i := 1; i <= tb.Count(); i++ {
			ret := []table.Account{}
			err := loader.LoadAll((tb).TableName()+cast.ToString(i), &ret)
			expect.Nil(err)
			all = append(all, ret...)
		}
	} else {
		err := loader.LoadAll((tb).TableName(), &all)
		expect.Nil(err)
	}

	for _, data := range all {
		account := &Account{table: data}
		s.accountsByPlatformId[account.table.PlatformUUId] = account
		s.accountsByUId[account.table.UUId] = account
		lastLoginRole := account.table.Roles[account.table.LastRoleId]
		if lastLoginRole != nil {
			s.accountsByName[lastLoginRole.Name] = account
			account.serverId = common.GameName(int32(lastLoginRole.SId))
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
	newAcc.table.UUId = s.uuidGen.Uuid()
	newAcc.table.PlatformUUId = platformId
	newAcc.table.OS = msg.OS
	newAcc.table.ClientVersion = msg.ClientVersion // newRole
	newRole := &table.Role{
		UUId:     newAcc.table.UUId,
		RoleId:   s.uuidGen.Uuid(),
		SId:      1,
		Name:     fmt.Sprintf("player_%v", newAcc.table.UUId),
		Icon:     "Avatar",
		Country:  "中国",
		IsDelete: false,
		CreateAt: time.Now().Unix(),
		LoginAt:  0,
		LogoutAt: 0,
	}
	newAcc.table.Roles = table.RoleMap{newRole.RoleId: newRole}
	newAcc.table.LastRoleId = newRole.RoleId
	newAcc.serverId = common.GameName(int32(newRole.SId))

	// 回存db
	if err := saver.Save(&newAcc.table, newRole); err != nil {
		log.KV("err", err).Error("save err")
		return nil, false
	}

	s.accountsByPlatformId[platformId] = newAcc
	s.accountsByUId[newAcc.table.UUId] = newAcc

	return newAcc, true
}
