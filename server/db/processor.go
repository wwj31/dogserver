package db

import (
	"fmt"
	"reflect"
	"server/common/log"
	"server/db/table"
	"time"

	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/tools"
	"gorm.io/gorm"
)

const (
	INSERT = iota + 1
	UPDATE
)

type operatorTable struct {
	status int
	t      table.Tabler
}
type processor struct {
	actor.Base
	session      *gorm.DB
	list         []operatorTable
	nextExecTime string
}

func (s *processor) OnInit() {
	s.list = make([]operatorTable, 0, 100)
}

func (s *processor) OnHandleMessage(sourceId, targetId string, msg interface{}) {
	newOpera, ok := msg.(operatorTable)
	if !ok {
		log.Errorw("processor receive invalid data", "type", reflect.TypeOf(msg).String())
		return
	}

	var repeated bool
	for index, opera := range s.list {
		// 相同数据执行替换操作
		if opera.t.TableName() == newOpera.t.TableName() &&
			opera.t.Key() == newOpera.t.Key() &&
			opera.status == newOpera.status {

			s.list[index] = newOpera
			repeated = true
			break
		}
	}

	if !repeated {
		s.list = append(s.list, newOpera)
	}

	if s.nextExecTime == "" {
		s.nextExecTime = s.AddTimer("", tools.NowTime()+int64(500*time.Millisecond), func(dt int64) {
			s.nextExecTime = ""
			s.execute()
		})
	}
}

func (s *processor) execute() {
	for _, v := range s.list {
		fmt.Println("exec ", v.t.TableName(), v.t.Key())
	}
	s.list = s.list[:0]
}
