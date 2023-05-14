package account

import (
	"server/common/log"
	"server/common/mongodb"
)

const Collection = "account"

type Account struct {
	UUID          string `bson:"_id"`
	WeiXinOpenID  string `bson:"wei_xin_open_id" index:"true"`
	DeviceID      string `bson:"device_id" index:"true"`
	Phone         string `bson:"phone" index:"true"`
	ShorID        int64  `bson:"shor_id" index:"true"`
	LastLoginRID  string `bson:"last_login_rid" index:"true"`
	OS            string
	ClientVersion string
	Roles         map[string]Role
}

func CreateIndex() {
	err := mongodb.Ins.CreateIndex(Collection, &Account{})
	if err != nil {
		log.Errorf("create mongo index failed", "coll", Collection, "err", err)
	}
}

type Role struct {
	RID      string
	CreateAt string
}

func New() *Account {
	return &Account{}
}
