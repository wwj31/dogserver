package table

import (
	"reflect"
)

// 玩家角色属性
type Role struct {
	RoleId     uint64 `gorm:"primary_key"` //角色ID
	UUId       uint64 `gorm:"index"`       //账号ID
	SId        uint64 //分配的区服ID
	Name       string `gorm:"index"`
	Icon       string //头像
	Country    string
	IsDelete   bool         //是否删除
	Attributes AttributeMap `gorm:"type:json"` // 属性集
	CreateAt   int64
	LoginAt    int64
	LogoutAt   int64
}

func init() {
	RegisterTable(&Role{})
}

func (s *Role) TableName() string {
	return reflect.TypeOf(s).Elem().Name()
}

// 分表数量
func (s *Role) Count() int {
	return split
}

func (s *Role) Key() uint64 {
	return s.RoleId
}
