package account

import (
	"fmt"
	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/db/table"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/iface"
	"time"

	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/expect"
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
			var ret []table.Account
			err := loader.LoadAll((tb).ModelName()+cast.ToString(i), &ret)
			expect.Nil(err)
			all = append(all, ret...)
		}
	} else {
		err := loader.LoadAll((tb).ModelName(), &all)
		expect.Nil(err)
	}

	for _, data := range all {
		account := &Account{table: data}
		s.accountsByPlatformId[account.table.PlatformUUId] = account
		s.accountsByUId[account.table.UUId] = account
		lastLoginRole := account.table.Roles[account.table.LastRoleId]
		if lastLoginRole != nil {
			s.accountsByName[lastLoginRole.Name] = account
			account.serverId = actortype.GameName(int32(lastLoginRole.SId))
		}
	}
}

func (s *AccountMgr) Login(msg *outer.LoginReq, storer iface.Storer) (acc *Account, new bool) {
	platformId := combine(msg.PlatformName, msg.PlatformUUID)
	if acc = s.accountsByPlatformId[platformId]; acc != nil {
		return acc, false
	}

	newAcc := &Account{}
	newAcc.table.UUId = s.uuidGen.GenUuid()
	newAcc.table.PlatformUUId = platformId
	newAcc.table.OS = msg.OS
	newAcc.table.ClientVersion = msg.ClientVersion

	newRole := &inner.RoleInfo{
		UUId:     newAcc.table.UUId,
		RoleId:   s.uuidGen.GenUuid(),
		SId:      1,
		Name:     fmt.Sprintf("player_%v", newAcc.table.UUId),
		Icon:     "Avatar",
		Country:  "中国",
		CreateAt: time.Now().Unix(),
		LoginAt:  0,
		LogoutAt: 0,
	}
	newAcc.table.Roles = table.RoleMap{newRole.RoleId: newRole}
	newAcc.table.LastRoleId = newRole.RoleId
	newAcc.serverId = actortype.GameName(int32(newRole.SId))

	// 回存db
	if err := storer.Store(true, &newAcc.table); err != nil {
		log.Errorw("save failed", "err", err)
		return nil, false
	}

	s.accountsByPlatformId[platformId] = newAcc
	s.accountsByUId[newAcc.table.UUId] = newAcc

	return newAcc, true
}
