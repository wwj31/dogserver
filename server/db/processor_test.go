package db

import (
	"fmt"
	"server/db/table"
	"testing"
	"time"

	"github.com/spf13/cast"

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

	_ = sys.Add(actor.New(process, &processor{session: db.Session(&gorm.Session{})}, actor.SetMailBoxSize(10000)))

	var total []uint64
	for i := 0; i < 1000; i++ {
		r := tools.Randx_y(1, 99999)
		total = append(total, uint64(r))
	}

	insertMap := make(map[uint64]bool)
	data := "table"
	go func() {
		var operaArr []operator
		for i := 0; i < 100000; i++ {
			id := total[tools.Randx_y(0, len(total))]
			var opera op
			var finish chan struct{}
			if !insertMap[id] {
				opera = _INSERT
				insertMap[id] = true
			} else {
				opera = op(tools.Randx_y(2, 4))
				data = "table:" + cast.ToString(tools.NowTime())
				if opera == _LOAD {
					finish = make(chan struct{}, 1)
				}
			}
			oper := operator{
				state:  opera,
				tab:    &table.Fake{Id: id, Data: data},
				finish: finish,
			}
			if oper.finish != nil {
				go func() {
					<-finish
					fmt.Println("                             load success", oper.tab.Key())
				}()
			}
			operaArr = append(operaArr, oper)
		}

		for _, opera := range operaArr {
			time.Sleep(time.Duration(tools.Randx_y(60000, 100000)))
			_ = sys.Send("", process, "", opera)
		}
	}()
	<-sys.CStop
}
