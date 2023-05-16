package account

import (
	"time"

	"server/common/log"
	"server/common/mongodb"
	"server/proto/innermsg/inner"
)

const Collection = "account"

type Account struct {
	UUID          string `bson:"_id"`
	WeiXinOpenID  string `bson:"wei_xin_open_id" index:"true"`
	DeviceID      string `bson:"device_id" index:"true"`
	Phone         string `bson:"phone" index:"true"`
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
func (a *Account) ToPb() *inner.Account {
	return &inner.Account{
		UID:           a.UUID,
		DeviceID:      a.DeviceID,
		Phone:         a.Phone,
		OS:            a.OS,
		ClientVersion: a.ClientVersion,
	}
}

type Role struct {
	RID      string
	ShorID   int64
	CreateAt time.Time
}

func New() *Account {
	return &Account{}
}
