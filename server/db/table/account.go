package table

import (
	"reflect"
)

const split = 2

type Account struct {
	UUId         uint64 `gorm:"primary_key"` //账号ID
	PlatformUUId string `gorm:"index"`       //平台Id

	//最近一次登录角色Id
	LastRoleId uint64

	// 一个玩家可以拥有多个角色
	Roles RoleMap `gorm:"type:json"`

	// 客户端版本号
	ClientVersion string

	OS string
}

func init() {
	RegisterTable(&Account{})
}

func (s *Account) TableName() string {
	return reflect.TypeOf(s).Elem().Name()
}

// 分表数量
func (s *Account) Count() int {
	return split
}

func (s *Account) Key() uint64 {
	return s.UUId
}
