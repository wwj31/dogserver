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

type op int

const (
	_INSERT op = iota + 1
	_UPDATE
	_LOAD
)

type (
	operator struct {
		status  op
		tab     table.Tabler
		inserts []table.Tabler
		finish  chan<- struct{}
	}
	processor struct {
		actor.Base
		session      *gorm.DB
		list         []operator
		nextExecTime string
	}
)

func (s *processor) OnInit() {
	s.list = make([]operator, 0, 100)
}

func (s *processor) OnHandleMessage(sourceId, targetId string, msg interface{}) {
	newOpera, ok := msg.(operator)
	if !ok {
		log.Errorw("processor receive invalid data", "type", reflect.TypeOf(msg).String())
		return
	}

	var repeated bool
	for i := len(s.list) - 1; i >= 0; i-- {
		opera := s.list[i]
		if opera.tab.TableName() != newOpera.tab.TableName() {
			continue
		}
		batches := append([]table.Tabler{opera.tab}, opera.inserts...)
		for _, tab := range batches {
			// 1.相同数据store操作执行替换
			// 2.load操作优先从队列里取,取不到再读库
			if tab.Key() == newOpera.tab.Key() {
				switch newOpera.status {
				case _INSERT, _UPDATE:
					if opera.status == newOpera.status {
						s.list[i] = newOpera
						repeated = true
						break
					}
				case _LOAD:
					switch opera.status {
					case _INSERT, _UPDATE:
						newOpera.tab = opera.tab
						if newOpera.finish != nil {
							newOpera.finish <- struct{}{}
						}
						return
					}
				}
			} else {
				if newOpera.status == _INSERT && opera.status == _INSERT {
					// 同表不同key，合并insert操作
					opera.inserts = append(opera.inserts, newOpera.tab)
					break
				}
			}
		}
	}
	if !repeated {
		s.list = append(s.list, newOpera)
	}

	if s.nextExecTime == "" && len(s.list) > 0 {
		if s.list[0].status == _LOAD {
			s.execute()
			return
		}
		s.delayExec()
	}
}

func (s *processor) delayExec() {
	s.nextExecTime = s.AddTimer("", tools.NowTime()+int64(500*time.Millisecond), func(dt int64) {
		s.nextExecTime = ""
		s.execute()
	})
}

const maxCount = 100

func (s *processor) execute() {
	for i, v := range s.list {
		fmt.Println("exec ", v.tab.TableName(), v.tab.Key())
		db := s.session.Table(v.tab.TableName())
		if v.status == _INSERT {
			inserts := append([]table.Tabler{v.tab}, v.inserts...)
			db.Create(inserts)
		} else if v.status == _UPDATE {
			db.Updates(v.tab)
		} else if v.status == _LOAD {
			db.Take(v.tab)
		}

		count := i + 1
		if count == maxCount && count < len(s.list) {
			s.list = s.list[i+1:]
			s.delayExec()
			return
		}
	}
	s.list = s.list[:0]
}
