package db

import (
	"fmt"
	"server/db/table"
	"testing"
	"time"

	"github.com/spf13/cast"

	"github.com/wwj31/dogactor/expect"

	"github.com/wwj31/dogactor/tools"

	"github.com/wwj31/dogactor/actor"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const process = "process1"

func TestProcessor(t *testing.T) {
	db, _ := gorm.Open(mysql.Open("root:starunion@tcp(127.0.0.1:3306)/game?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		NamingStrategy:         schema.NamingStrategy{SingularTable: true},
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(&table.Fake{})
	sys, _ := actor.NewSystem()

	_ = sys.Add(actor.New(process, &processor{session: db.Session(&gorm.Session{})}, actor.SetMailBoxSize(200)))

	var total []uint64
	for i := 0; i < 10; i++ {
		r := tools.Randx_y(1, 99999)
		total = append(total, uint64(r))
	}

	insertMap := make(map[uint64]bool)
	data := "table"
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Microsecond)
			id := total[tools.Randx_y(0, len(total))]
			var opera op
			if !insertMap[id] {
				opera = _INSERT
				insertMap[id] = true
			} else {
				opera = op(tools.Randx_y(2, 4))
				data = "table:" + cast.ToString(tools.NowTime())
			}

			err := sys.Send("", process, "", operator{
				status: opera,
				tab:    &table.Fake{Id: id, Data: data},
			})
			expect.Nil(err)
			if err == nil {
				fmt.Println("send success", id, opera, data)
			}
		}
	}()
	<-sys.CStop
}
