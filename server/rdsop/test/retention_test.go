package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/wwj31/dogactor/tools"

	"server/common/rds"
	"server/rdsop"
)

func init() {
	if err := rds.Connect("redis://:123456@localhost:6379/3", false); err != nil {
		log.Fatalf("redids connect failed:%v", err)
	}
}

func TestRetention(t *testing.T) {
	rdsop.AddDailyRegistry(1)
	rdsop.AddDailyRegistry(2)
	rdsop.AddDailyRegistry(3)
	rdsop.AddDailyRegistry(4)
	rdsop.AddDailyRegistry(5)
	rdsop.AddDailyRegistry(6)

	rdsop.AddDailyLogin(2)
	rdsop.AddDailyLogin(4)
	total, v := rdsop.RetentionOf(tools.Now().Local(), 0)
	fmt.Println(total, v)
}
