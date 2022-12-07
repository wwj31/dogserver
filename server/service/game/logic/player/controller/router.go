package controller

import (
	"fmt"
	"reflect"

	"github.com/wwj31/dogactor/log"
	"server/service/game/iface"
)

type Handler func(player iface.Player, v interface{})

var router = map[string]Handler{}

func registry(msg interface{}, fun Handler) bool {
	msgName := MsgName(msg)
	if _, ok := router[msgName]; ok {
		panic(fmt.Errorf("%v repeated ", msg))
	}
	router[msgName] = fun
	return true
}
func GetHandler(msgName string) (Handler, bool) {
	fn, ok := router[msgName]
	return fn, ok
}

func MsgName(v interface{}) string {
	str := reflect.TypeOf(v).String()
	if str[0] == '*' {
		str = str[1:]
	}
	return str
}

func argInfo(cb Handler) (reflect.Type, int) {
	cbType := reflect.TypeOf(cb)
	if cbType.Kind() != reflect.Func {
		log.SysLog.Errorw("nats: Handler needs to be a func")
		return nil, 0
	}
	numArgs := cbType.NumIn()
	if numArgs == 0 {
		return nil, numArgs
	}
	return cbType.In(numArgs - 1), numArgs
}
