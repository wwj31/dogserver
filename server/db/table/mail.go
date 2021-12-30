package table

import (
	"reflect"
)

type Mail struct {
	UUId  uint64 `gorm:"primary_key"` //账号ID
	Bytes []byte
}

func init() {
	RegisterTable(&Mail{})
}

func (s *Mail) TableName() string {
	return reflect.TypeOf(s).Elem().Name()
}

// 分表数量
func (s *Mail) Count() int {
	return split
}

func (s *Mail) Key() uint64 {
	return s.UUId
}