package router

import (
	"reflect"
	"server/common/log"
	"sync"

	gogo "github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
)

var router = sync.Map{}

func Reg[ACTOR actor.Actor, MSG gogo.Message](fn func(actor ACTOR, msg MSG), repeat ...bool) bool {
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
		fn(actor.(ACTOR), message.(MSG))
	})
	return true
}

func Dispatch(a actor.Actor, msg gogo.Message) {
	msgName := reflect.TypeOf(msg).String()
	actorName := reflect.TypeOf(a).String()

	v, ok := router.Load(msgName)
	if !ok {
		log.Errorw("not find msg handler", "msg", msgName)
		return
	}

	handlers, _ := v.(*sync.Map)
	v, ok = handlers.Load(actorName)
	if !ok {
		log.Errorw("not find msg handler without actor", "msg", msgName, "actor", actorName)
		return
	}

	handle, _ := v.(func(actor actor.Actor, message gogo.Message))
	handle(a, msg)
}
