package table

import "reflect"

// 玩家角色属性

type Player struct {
	RoleId    uint64 `gorm:"primary_key"` //角色ID
	RoleBytes []byte // 角色属性
	ItemBytes []byte // 道具数据
	MailBytes []byte // 邮件数据
}

func init() {
	RegisterTable(&Player{})
}

func (s *Player) TableName() string {
	return reflect.TypeOf(s).Elem().Name()
}

// 分表数量
func (s *Player) Count() int {
	return split
}

func (s *Player) Key() uint64 {
	return s.RoleId
}
