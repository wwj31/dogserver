package db

import (
	"fmt"
	"reflect"
	"server/common/log"
	"server/db/table"
	"sync/atomic"

	"github.com/spf13/cast"

	"github.com/wwj31/dogactor/actor"
	"gorm.io/gorm"
)

type op int

const (
	_INSERT op = iota + 1
	_UPDATE
	_LOAD
)

const (
	STOP = iota + 1
	RUNING
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
		session *gorm.DB
		set     []operator
		state   atomic.Value
	}
)

func (s *processor) OnInit() {
	s.state.Store(STOP)
	s.set = make([]operator, 0, 100)
}

func (s *processor) OnHandleMessage(sourceId, targetId string, msg interface{}) {
	exec, ok := msg.(string)
	if ok && exec == "exec" {
		s.cas()
		return
	}

	newOpera, ok := msg.(operator)
	if !ok {
		log.Errorw("processor receive invalid data", "type", reflect.TypeOf(msg).String())
		return
	}
	add := true
	for i := len(s.set) - 1; i >= 0; i-- {
		opera := s.set[i]
		if tableName(opera.tab.ModelName(), opera.tab.Count()) != tableName(newOpera.tab.ModelName(), newOpera.tab.Count()) {
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
						s.set[i] = newOpera
						add = false
						break
					}
				case _LOAD:
					switch opera.status {
					case _INSERT, _UPDATE:
						newOpera.tab = opera.tab
						if newOpera.finish != nil {
							if cap(newOpera.finish) == 0 {
								add = false
								log.Warnw("operator _LOAD finish cap == 0 can't push", "state", newOpera.status, "tab name", newOpera.tab.ModelName())
								break
							}
							newOpera.finish <- struct{}{}
						}
						add = false
						break
					}
				}
			} else {
				if newOpera.status == _INSERT && opera.status == _INSERT {
					// 同表不同key，合并insert操作
					opera.inserts = append(opera.inserts, newOpera.tab)
					add = false
					break
				}
			}
		}
		if !add {
			break
		}
	}
	if add {
		s.set = append(s.set, newOpera)
	}
	s.cas()
}

func (s *processor) cas() {
	if s.state.CompareAndSwap(STOP, RUNING) {
		arr := make([]operator, len(s.set))
		copy(arr, s.set)
		s.set = s.set[:0]
		go func() {
			//fmt.Println("start processor ", s.ID(), "arr len ", len(arr))
			for _, v := range arr {
				s.execute(v)
			}
			s.state.Store(STOP)
			s.Send(s.ID(), "exec")
		}()
	}
}

func (s *processor) execute(op operator) {
	tn := op.tab.ModelName()
	if op.tab.Count() > 0 {
		tn = tn + cast.ToString(op.tab.Count())
	}
	db := s.session.Table(op.tab.ModelName())
	if op.status == _INSERT {
		inserts := append([]table.Tabler{op.tab}, op.inserts...)
		for _, v := range inserts {
			fmt.Println("insert", v.Key())
			db.Create(v) // todo create 不支持 接口切片 怎么处理批量插入？？？？
		}
	} else if op.status == _UPDATE {
		fmt.Println("update", op.tab.Key())
		db.Updates(op.tab)
	} else if op.status == _LOAD {
		db.Take(op.tab)
	}
}
