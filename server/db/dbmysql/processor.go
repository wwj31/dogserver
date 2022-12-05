package dbmysql

import (
	"fmt"
	"reflect"
	"server/common/log"
	"server/db/dbmysql/table"
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
	RUNNING
)

const notify = "check"

type (
	operator struct {
		state  op
		tabler table.Tabler

		// 当state为INSERT时，extrOpera用于存储update，
		// insert和update无法合并，因为update使用gorm的update()结构体零值不会保存
		extrOpera *operator

		finish chan<- struct{}
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

func (s *processor) OnHandle(m actor.Message) {
	rawMsg := m.RawMsg()
	check, ok := rawMsg.(string)
	if ok && check == notify {
		if len(s.set) != 0 {
			s.processing()
		}
		return
	}

	newOpera, ex := rawMsg.(operator)
	if !ex {
		log.Errorw("processor receive invalid data", "type", reflect.TypeOf(rawMsg).String())
		return
	}

	s.mergeAndCover(newOpera)
	s.processing()
}

func (s *processor) mergeAndCover(newOpera operator) {
	newTableName := tableName(newOpera.tabler.ModelName(), newOpera.tabler.SplitNum())
	tableNameKey := newTableName + "_" + cast.ToString(newOpera.tabler.Key())
	if oldOpera, ok := s.set[tableNameKey]; !ok {
		s.set[tableNameKey] = newOpera
	} else {
		// 1.相同key的数据store操作执行替换
		// 2.load操作优先从队列里取,取不到再读库
		switch newOpera.state {
		case _INSERT:
			log.Errorw("exception error operation before insert", "op", oldOpera.state, "tabler", oldOpera.tabler.Key())
		case _UPDATE:
			switch oldOpera.state {
			case _UPDATE:
				oldOpera.tabler = newOpera.tabler // cover
			case _INSERT:
				oldOpera.extrOpera = &newOpera // insert with update
			}
		case _LOAD:
			switch oldOpera.state {
			case _INSERT, _UPDATE:
				if oldOpera.extrOpera == nil {
					newOpera.tabler = oldOpera.tabler
				} else {
					newOpera.tabler = oldOpera.extrOpera.tabler
				}
				if newOpera.finish != nil {
					if cap(newOpera.finish) == 0 {
						log.Warnw("operator _LOAD finish cap == 0 can't push", "state", newOpera.state, "tabler name", newOpera.tabler.ModelName())
						return
					}

					select {
					case newOpera.finish <- struct{}{}:
					default:
					}
				}
			case _LOAD:
				s.set[tableNameKey] = newOpera
			}
		}
	}
	return
}

func (s *processor) processing() {
	if s.state.CompareAndSwap(STOP, RUNNING) {
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
			_ = s.Send(s.ID(), notify)
		}()
	}
}

func (s *processor) execute(op operator) {
	tn := op.tabler.ModelName()
	if op.tabler.SplitNum() > 0 {
		tn = tn + cast.ToString(op.tabler.SplitNum())
	}
	db := s.session.Table(tn)
	if op.state == _INSERT {
		fmt.Println("insert")
		db.Create(op.tabler) // todo create 不支持 接口切片 怎么处理批量插入？？？？
		if op.extrOpera != nil {
			db.Updates(op.extrOpera.tabler)
		}
	} else if op.state == _UPDATE {
		db.Updates(op.tabler)
	} else if op.state == _LOAD {
		err := db.Take(op.tabler).Error
		if err != nil {
			log.Errorw("load faild ", "tabler", op.tabler.ModelName(), "key", op.tabler.Key())
		}

		select {
		case op.finish <- struct{}{}:
		default:
		}
	}
}
