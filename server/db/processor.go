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

const notify = "check"

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
		set     map[string]operator
		state   atomic.Value
	}
)

func (s *processor) OnInit() {
	s.state.Store(STOP)
	s.set = make(map[string]operator, 10)
}

func (s *processor) OnHandleMessage(sourceId, targetId string, msg interface{}) {
	check, ok := msg.(string)
	if ok && check == notify {
		if len(s.set) != 0 {
			s.store()
		}
		return
	}

	newOpera, ex := msg.(operator)
	if !ex {
		log.Errorw("processor receive invalid data", "type", reflect.TypeOf(msg).String())
		return
	}
	s.merageAndCover(newOpera)
	s.store()
}

func (s *processor) merageAndCover(newOpera operator) {
	newTableName := tableName(newOpera.tab.ModelName(), newOpera.tab.Count())
	tableNameKey := newTableName + "_" + cast.ToString(newOpera.tab.Key())
	if oldOpera, ok := s.set[tableNameKey]; !ok {
		s.set[tableNameKey] = newOpera
	} else {
		// 1.相同key的数据store操作执行替换
		// 2.load操作优先从队列里取,取不到再读库
		switch newOpera.status {
		case _INSERT:
			log.Errorw("exception error operation before insert", "op", oldOpera.status, "tab", oldOpera.tab.Key())
		case _UPDATE:
			if oldOpera.status == newOpera.status {
				oldOpera.tab = newOpera.tab // cover
			} else {
				s.set[tableNameKey] = newOpera
			}
		case _LOAD:
			switch oldOpera.status {
			case _INSERT, _UPDATE:
				newOpera.tab = oldOpera.tab
				if newOpera.finish != nil {
					if cap(newOpera.finish) == 0 {
						log.Warnw("operator _LOAD finish cap == 0 can't push", "state", newOpera.status, "tab name", newOpera.tab.ModelName())
						return
					}
					newOpera.finish <- struct{}{}
				}
			case _LOAD:
				s.set[tableNameKey] = newOpera
			}
		}
	}
	return
}

func (s *processor) store() {
	if s.state.CompareAndSwap(STOP, RUNING) {
		arr := make([]operator, 0, len(s.set))
		for _, v := range s.set {
			arr = append(arr, v)
		}
		s.set = make(map[string]operator, 10)
		go func() {
			for _, v := range arr {
				s.execute(v)
			}
			s.state.Store(STOP)
			s.Send(s.ID(), notify)
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
		fmt.Println("insert", len(inserts))
		for _, v := range inserts {
			db.Create(v) // todo create 不支持 接口切片 怎么处理批量插入？？？？
		}
	} else if op.status == _UPDATE {
		db.Updates(op.tab)
	} else if op.status == _LOAD {
		db.Take(op.tab)
	}
}
