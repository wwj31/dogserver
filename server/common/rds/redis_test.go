package rds

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestRedis(t *testing.T) {
	if err := Connect("redis://:123456@localhost:6379/1", false); err != nil {
		log.Fatalf(err.Error())
	}
	ctx := context.Background()
	sr := Ins.Set(ctx, "foo", "bar", 0)
	fmt.Println(sr.String())
	gr := Ins.Get(ctx, "foo")
	fmt.Println(gr.String())
}
