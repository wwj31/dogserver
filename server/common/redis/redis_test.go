package redis

import (
	"context"
	"fmt"
	"testing"
)

func TestRedis(t *testing.T) {
	err := Builder().OnConnect(func() {
		fmt.Println("redis connect success")
	}).Connect()

	if err != nil {
		fmt.Println(err)
		return
	}
	ctx := context.Background()
	sr := Ins.Set(ctx, "foo", "bar", 0)
	fmt.Println(sr.String())
	gr := Ins.Get(ctx, "foo")
	fmt.Println(gr.String())
}
