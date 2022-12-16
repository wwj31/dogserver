package ge_test

import (
	"fmt"
	gogo "github.com/gogo/protobuf/proto"
	"reflect"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"testing"
)

var _ = reg(func(val int, msg *outer.LoginReq) {
	fmt.Println(val, msg.PlatformName)
})

var _ = reg(func(val int, msg *inner.ItemInfo) {
	fmt.Println(val, msg.RID)
})

func TestGeneric(t *testing.T) {
	msg := &outer.LoginReq{
		PlatformUUID:  "吃了一粒",
		PlatformName:  "布洛芬",
		OS:            "我tm羊了",
		ClientVersion: "38度写代码",
		Token:         "爽！",
	}
	var v gogo.Message
	v = msg
	On(v)
	On(&inner.ItemInfo{
		RID:   "foo",
		Items: nil,
	})
}

var router = map[string]func(msg gogo.Message){}

func reg[T gogo.Message](fn func(val int, msg T)) error {
	var t T
	name := reflect.TypeOf(t).String()
	if _, exist := router[name]; exist {
		fmt.Println("error")
		return nil
	}

	router[name] = func(message gogo.Message) {
		fn(123, message.(T))
	}
	return nil
}

func On(msg gogo.Message) {
	name := reflect.TypeOf(msg).String()
	v, _ := router[name]
	v(msg)
}
