package account

import (
	"fmt"
	"server/common"
	"server/common/actortype"
	"server/common/log"
	table2 "server/db/dbmysql/table"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/iface"
	"time"

	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/expect"
)

type Mgr struct {
	accountsByUId        map[uint64]*Account
	accountsByPlatformId map[string]*Account
	accountsByName       map[string]*Account

	uuidGen common.UID
}

func NewAccountMgr() Mgr {
	return Mgr{
		accountsByUId:        make(map[uint64]*Account),
		accountsByPlatformId: make(map[string]*Account),
		accountsByName:       make(map[string]*Account),
		uuidGen:              common.NewUID(1),
	}
}

func (s *Mgr) LoadAllAccount(loader iface.Loader) {
	var all []table2.Account
	tb := &table2.Account{}
	if tb.SplitNum() > 1 {
		for i := 1; i <= tb.SplitNum(); i++ {
			var ret []table2.Account
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
		//lastLoginRole := account.table.Roles[account.table.LastRoleId]
		//if lastLoginRole != nil {
		//	s.accountsByName[lastLoginRole.Name] = account
		//	account.serverId = actortype.GameName(int32(lastLoginRole.SId))
		//}
	}
}

func (s *Mgr) Login(msg *outer.LoginReq, store iface.Storer) (acc *Account, new bool) {
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
		//UID:      newAcc.table.UUId,
		//RID:      s.uuidGen.GenUuid(),
		SId:      1,
		Name:     fmt.Sprintf("player_%v", newAcc.table.UUId),
		Icon:     "Avatar",
		Country:  "中国",
		CreateAt: time.Now().Unix(),
		LoginAt:  0,
		LogoutAt: 0,
	}
	//newAcc.table.Roles = table2.RoleMap{newRole.RoleId: newRole}
	//newAcc.table.LastRoleId = newRole.RoleId
	newAcc.serverId = actortype.GameName(int32(newRole.SId))

	// 回存db
	if err := store.Store(true, &newAcc.table); err != nil {
		log.Errorw("save failed", "err", err)
		return nil, false
	}

	s.accountsByPlatformId[platformId] = newAcc
	s.accountsByUId[newAcc.table.UUId] = newAcc

	return newAcc, true
}
