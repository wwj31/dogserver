package redis

import (
	"context"
	"fmt"
	"testing"
)

func TestRedis(t *testing.T) {
	Builder().OnConnect(func() {
		fmt.Println("redis connect success")
	}).OK()

	ctx := context.Background()
	sr := Ins.Set(ctx, "foo", "bar", 0)
	fmt.Println(sr.String())
	gr := Ins.Get(ctx, "foo")
	fmt.Println(gr.String())
}
