package router

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/golang/protobuf/proto"

	"server/common/log"

	"github.com/wwj31/dogactor/actor"
)

var router = sync.Map{}
var result = sync.Map{}

func Reg[ACTOR actor.Actor, MSG proto.Message](fn func(actor ACTOR, msg MSG) any, repeat ...bool) bool {
	var (
		msg     MSG
		act     ACTOR
		msgName = reflect.TypeOf(msg).String()
		actName = reflect.TypeOf(act).String()
	)

	v, _ := router.LoadOrStore(msgName, &sync.Map{})
	handlers, _ := v.(*sync.Map)

	if _, ok := handlers.Load(actName); ok {
		log.Warnw("repeat handler ", "actor", actName, "msg", msgName)
		if len(repeat) == 0 || !repeat[0] {
			return false
		}
	}

	handlers.Store(actName, func(actor actor.Actor, message proto.Message) {
		defer func() {
			if r := recover(); r != nil {
				stack := fmt.Sprintf("panic actor:[%v] message:[%T]  info:%v recover:%v", actor.ID(), message, message.String(), r)
				log.Errorf(stack)
			}
		}()
		log.Infow("router INPUT", "actor", actor.ID(), "msg", reflect.TypeOf(message), "message", message.String())

		ret := fn(actor.(ACTOR), message.(MSG))

		log.Infow("router OUTPUT", "actor", actor.ID(), "msg", reflect.TypeOf(ret), "ret", ret)
		val, ex := result.Load(actor.ID())
		if ex {
			response := val.(func(any))
			response(ret)
		}
	})
	return true
}

func Result(a actor.Actor, fn func(result any)) {
	result.Store(a.ID(), fn)
}

func Dispatch(a actor.Actor, msg proto.Message) error {
	msgName := reflect.TypeOf(msg).String()
	actorName := reflect.TypeOf(a).String()

	v, ok := router.Load(msgName)
	if !ok {
		return fmt.Errorf("not find msg:[%v] handler", msgName)
	}

	handlers, _ := v.(*sync.Map)
	v, ok = handlers.Load(actorName)
	if !ok {
		return fmt.Errorf("not find msg:[%v] handler without actor:[%v]", msgName, actorName)
	}

	handle, _ := v.(func(actor actor.Actor, message proto.Message))
	handle(a, msg)
	return nil
}
