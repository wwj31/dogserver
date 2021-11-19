package db

import (
	"fmt"
	"github.com/wwj31/dogactor/expect"
	"server/service/db/table"
	"testing"
)

func TestDBLoad(t *testing.T) {
	dbIns := New(
		"root:123456@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"game",
	)
	user := &table.Role{RoleId: 1637335554259221918}
	err := dbIns.Load(user)
	expect.Nil(err)

	fmt.Println(user)
}

func TestDBLoadAll(t *testing.T) {
	dbIns := New(
		"root:123456@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"game",
	)

	all := []*table.Role{}
	dbIns.LoadAll(&table.Role{RoleId: 1}, &all)
	fmt.Println(all)
}
