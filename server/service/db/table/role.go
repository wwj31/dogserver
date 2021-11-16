package table

import (
	"gorm.io/gorm"
	"reflect"
)

// 玩家角色属性
type Role struct {
	gorm.Model
	RoleId   uint64 `gorm:"primary_key"` //角色ID
	UUID     uint64 `gorm:"index"`       //账号ID
	SId      int64  //分配的区服ID
	Name     string `gorm:"index"`
	Icon     string //头像
	Country  string
	IsDelete bool //是否删除
	LoginAt  int64
	LogoutAt int64
}

func init() {
	RegisterTable(&Role{})
}

func (s *Role) TableName() string {
	return reflect.TypeOf(s).Elem().Name()
}

// 分表数量
func (s *Role) Count() int {
	return 0
}
