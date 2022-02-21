package db

import (
	"testing"
	"time"

	"github.com/wwj31/dogactor/tools"

	"github.com/wwj31/dogactor/actor"
)

type fakeTable struct {
	key       uint64
	tableName string
}

func (f fakeTable) TableName() string {
	return f.tableName
}

func (f fakeTable) Count() int {
	return 0
}

func (f fakeTable) Key() uint64 {
	return f.key
}

const process = "process1"

func TestProcessor(t *testing.T) {
	sys, _ := actor.NewSystem()
	_ = sys.Add(actor.New("process1", &processor{}))
	_ = sys.Send("", process, "", operator{
		status: 1,
		tab:    fakeTable{key: 1, tableName: "table1"},
	})

	_ = sys.Send("", process, "", operator{
		status: 0,
		tab:    fakeTable{key: 4, tableName: "table1"},
	})

	_ = sys.Send("", process, "", operator{
		status: 0,
		tab:    fakeTable{key: 2, tableName: "table1"},
	})

	_ = sys.Send("", process, "", operator{
		status: 2,
		tab:    fakeTable{key: 1, tableName: "table1"},
	})

	go func() {
		for {
			time.Sleep(time.Millisecond)
			_ = sys.Send("", process, "", operator{
				status: 0,
				tab:    fakeTable{key: uint64(tools.Randx_y(1, 5)), tableName: "table1"},
			})
		}
	}()
	<-sys.CStop
}
