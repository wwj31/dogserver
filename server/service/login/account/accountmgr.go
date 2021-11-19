package account

import (
	"fmt"
	"github.com/wwj31/dogactor/log"
	"server/common"
	"server/proto/message"
	"server/service/db/table"
	"server/service/login/iface"
	"time"
)

type AccountMgr struct {
	accountsByUId        map[uint64]*Account
	accountsByPlatformId map[string]*Account

	uuidGen *common.UID
	saver   iface.Saver
}

func NewAccountMgr(saver iface.Saver) *AccountMgr {
	return &AccountMgr{
		accountsByUId:        make(map[uint64]*Account),
		accountsByPlatformId: make(map[string]*Account),
		uuidGen:              common.NewUID(),
		saver:                saver,
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

	// todo ... 回存db
	if err := s.saver.Save(&newAcc.Account); err != nil {
		log.KV("err", err).Error("save err")
		return nil, false
	}
	if err := s.saver.Save(newRole); err != nil {
		log.KV("err", err).Error("save err")
		return nil, false
	}
	return newAcc, true
}
