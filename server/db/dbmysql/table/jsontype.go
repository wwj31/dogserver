package table

import (
	"database/sql/driver"
	"encoding/json"
	"server/proto/innermsg/inner"
)

/*
	mysql5.7 版本支持json格式
	为json定义新的类型并实现 Value() (driver.Value, error) 和 Scan(input interface{}) error
	注:
		1.Value方法接收者必须是值类型！！    (否则增删查改会报错)
		2.Scan方法接收者必须是指针类型！！   (否则增删查改会报错)
    example:
		gorm的 mssql.JSON 定义
*/
type (
	RoleMap map[uint64]*inner.RoleInfo
)

// RoleMap
func (s RoleMap) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *RoleMap) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), s)
}
