package account

import (
	"fmt"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/log"
	"server/common"
	"server/db/table"
	"server/proto/message"
	"server/service/login/iface"
	"time"
)

type AccountMgr struct {
	accountsByUId        map[uint64]*Account
	accountsByPlatformId map[string]*Account
	accountsByName       map[string]*Account

	uuidGen *common.UID
	stored  iface.SaveLoader
}

func NewAccountMgr(stored iface.SaveLoader) *AccountMgr {
	return &AccountMgr{
		accountsByUId:        make(map[uint64]*Account),
		accountsByPlatformId: make(map[string]*Account),
		accountsByName:       make(map[string]*Account),
		uuidGen:              common.NewUID(1),
		stored:               stored,
	}
}

func (s *AccountMgr) LoadAllAccount() {
	var all []table.Account
	err := s.stored.LoadAll((&table.Account{}).TableName(), &all)
	expect.Nil(err)

	for _, data := range all {
		acc := &Account{Account: data}
		s.accountsByPlatformId[acc.PlatformUUId] = acc
		s.accountsByUId[acc.UUId] = acc
		lastLoginRole := acc.Roles[acc.LastRoleId]
		if lastLoginRole != nil {
			s.accountsByName[lastLoginRole.Name] = acc
		}
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
		Icon:     "Avatar",
		Country:  "中国",
		IsDelete: false,
		CreateAt: time.Now().Unix(),
		LoginAt:  time.Now().Unix(),
		LogoutAt: 0,
	}
	newAcc.Roles[newRole.RoleId] = newRole
	newAcc.LastRoleId = newRole.RoleId

	s.accountsByPlatformId[platformId] = newAcc
	s.accountsByUId[newAcc.UUId] = newAcc

	// 回存db
	if err := s.stored.Save(&newAcc.Account); err != nil {
		log.KV("err", err).Error("save err")
		return nil, false
	}
	if err := s.stored.Save(newRole); err != nil {
		log.KV("err", err).Error("save err")
		return nil, false
	}
	return newAcc, true
}
