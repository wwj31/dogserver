package table

import (
	"fmt"
	"reflect"
)

type Account struct {
	UUId         uint64 `gorm:"primary_key"` //账号ID
	PlatformUUId string `gorm:"index"`       //平台Id

	//最近一次登录角色Id
	LastRoleId uint64

	// map[Role.RoleId]*Role 一个玩家可以拥有多个角色
	Roles RoleMap `gorm:"type:json"`

	// map[SDK.SDKName+SDK.SDKId]*SDK 三方账号信息
	SDKInfo SDKMap `gorm:"type:json"`

	// 客户端版本号
	ClientVersion string
}

func init() {
	RegisterTable(&Account{})
}

func (s *Account) TableName() string {
	return reflect.TypeOf(s).Elem().Name()
}

// 分表数量
func (s *Account) Count() int {
	return 0
}

//三方账号信息
type SDK struct {
	SDKName string //平台昵称
	SDKId   string //三方账号ID
	UUId    int64  //账号ID
}

func SDKUUID(sdkId, sdkName string) string {
	return fmt.Sprintf("%s_%s", sdkName, sdkId)
}
