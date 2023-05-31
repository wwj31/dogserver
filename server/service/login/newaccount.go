package login

import (
	"context"
	"fmt"
	"time"

	"server/rdsop"

	"github.com/go-redis/redis/v9"
	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/tools"

	"server/common/log"
	"server/common/mongodb"
	"server/common/rds"
	"server/service/login/account"
)

func (s *Login) initAccount(acc *account.Account, os, ver string, upShortId int64) (err error) {
	var shortIdVal interface{}
	shortIdVal, err = rds.Ins.EvalSha(context.Background(), s.sha1, []string{rdsop.ShortIDKey()}).Result()
	if err == redis.Nil {
		log.Errorw("lua script exec failed", "err", err.Error())
		return
	}

	if err != nil {
		err = fmt.Errorf("shor id get failed :%v", err)
		log.Errorw(err.Error())
		return
	}

	arr, ok := shortIdVal.([]interface{})
	if !ok || len(arr) != 1 {
		err = fmt.Errorf("shortIdVal  failed:%v len:%v", shortIdVal, len(arr))
		log.Errorw(err.Error())
		return
	}

	acc.UUID = tools.XUID()
	acc.OS = os
	acc.ClientVersion = ver
	rid := acc.UUID
	acc.Roles = make(map[string]account.Role)
	newShortID := cast.ToInt64(arr[0])
	acc.Roles[rid] = account.Role{RID: rid, ShorID: newShortID, CreateAt: time.Now()}
	acc.LastShortID = acc.Roles[rid].ShorID
	acc.LastLoginRID = rid
	if _, err = mongodb.Ins.Collection(account.Collection).InsertOne(context.Background(), acc); err != nil {
		log.Errorw("login insert new account failed ", "UUID", acc.UUID, "err", err)
		return
	}

	// 绑定上级，建立层级关系
	if upShortId > 0 {
		rdsop.BindAgent(upShortId, newShortID)
	}
	return nil
}
