package db

import (
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
	newOper, ok := msg.(operatorTable)
	if !ok {
		log.Errorw("processor receive invalid data", "type", reflect.TypeOf(msg).String())
		return
	}

	for index, oper := range s.list {
		// 相同数据执行替换操作
		if oper.t.TableName() == newOper.t.TableName() && oper.t.Key() == newOper.t.Key() {
			s.list[index] = newOper
		}
	}

	s.list = append(s.list, newOper)

	if s.nextExecTime == "" {
		s.nextExecTime = s.AddTimer("", tools.NowTime()+int64(500*time.Millisecond), func(dt int64) {
			s.nextExecTime = ""
			s.execute()
		})
	}
}

func (s *processor) execute() {

}
