package db

import (
	"fmt"
	"github.com/wwj31/dogactor/expect"
	"server/db/table"
	"testing"
)

func TestDBLoad(t *testing.T) {
	dbIns := New(
		"root:123456@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"game",
	)
	user := &table.Role{RoleId: 1637335554259221918}
	acc := &table.Account{UUId: 123}
	err := dbIns.Load(user, acc)
	expect.Nil(err)

	fmt.Println(user)
	fmt.Println(acc)
}

func TestDBSave(t *testing.T) {
	dbIns := New(
		"root:123456@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"game",
	)
	user := &table.Role{RoleId: 1637335554259221918}
	acc := &table.Account{UUId: 123}
	err := dbIns.Save(user, acc)
	expect.Nil(err)
	fmt.Println(user)
}

func TestDBLoadAll(t *testing.T) {
	dbIns := New(
		"root:123456@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"game",
	)

	var all []table.Account
	tbName := (&table.Account{}).TableName()
	dbIns.LoadAll(tbName, &all)
	for _, v := range all {
		fmt.Println(v)
	}

}
