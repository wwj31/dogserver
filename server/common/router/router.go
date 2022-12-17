package router

import (
	"reflect"
	"server/common/log"
	"sync"

	gogo "github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
)

var router = sync.Map{}

func Reg[A actor.Actor, T gogo.Message](fn func(actor A, msg T), repeat ...bool) bool {
	var (
		t     T
		a     A
		tname = reflect.TypeOf(t).String()
		aname = reflect.TypeOf(a).String()
	)

	v, _ := router.LoadOrStore(tname, &sync.Map{})
	handlers, _ := v.(*sync.Map)

	if _, ok := handlers.Load(aname); ok {
		log.Warnw("repeat handler ", "actor", aname, "msg", tname)
		if len(repeat) == 0 || !repeat[0] {
			return false
		}
	}

	handlers.Store(aname, func(actor actor.Actor, message gogo.Message) {
		fn(actor.(A), message.(T))
	})
	return true
}

func On(a actor.Actor, msg gogo.Message) {
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
