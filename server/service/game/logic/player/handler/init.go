package handler

import (
	"fmt"
	"reflect"
	"server/service/game/iface"
)

type Handle func(player iface.Player, v interface{})

var MsgRouter = map[string]func(player iface.Player, v interface{}){}

func regist(msgName string, fun Handle) bool {
	if _, ok := MsgRouter[msgName]; ok {
		panic(fmt.Errorf("%v repeated ", msgName))
	}
	MsgRouter[msgName] = fun
	return true
}

func MsgName(v interface{}) string {
	return reflect.TypeOf(v).String()
}
