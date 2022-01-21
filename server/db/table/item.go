package table

import "reflect"

// 玩家道具表
type Item struct {
	RoleId    uint64 `gorm:"primary_key"` //角色ID
	Bytes     []byte
	ItemCount int
}

func init() {
	RegisterTable(&Item{})
}

func (s *Item) TableName() string {
	return reflect.TypeOf(s).Elem().Name()
}

// 分表数量
func (s *Item) Count() int {
	return split
}

func (s *Item) Key() uint64 {
	return s.RoleId
}
