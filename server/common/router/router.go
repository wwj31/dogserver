package router

import (
	"fmt"
	"reflect"
	"sync"

	"server/common/log"

	gogo "github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
)

var router = sync.Map{}
var result = sync.Map{}

func Reg[ACTOR actor.Actor, MSG gogo.Message](fn func(actor ACTOR, msg MSG) any, repeat ...bool) bool {
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

	handlers.Store(actName, func(actor actor.Actor, message gogo.Message) {
		log.Infow("router INPUT", "actor", actor.ID(), "msg", reflect.TypeOf(message), "message", message.String())

		ret := fn(actor.(ACTOR), message.(MSG))

		var retLog string
		if reflect.ValueOf(ret).IsNil() {
			retLog = "nil"
		} else {
			if returnMessage, ok := ret.(gogo.Message); ok {
				retLog = returnMessage.String()
			} else {
				retLog = "unknown"
			}
		}
		log.Infow("router OUTPUT", "actor", actor.ID(), "msg", reflect.TypeOf(ret), "ret", retLog)
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

func Dispatch(a actor.Actor, msg gogo.Message) error {
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

	handle, _ := v.(func(actor actor.Actor, message gogo.Message))
	handle(a, msg)
	return nil
}
