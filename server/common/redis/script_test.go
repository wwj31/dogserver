package redis

import (
	"fmt"
	"testing"
)

func TestScript(t *testing.T) {
	Builder().OnConnect(func() {
		fmt.Println("redis connect success")
	}).Connect()

	Ins.Set(Ins.Context(), "abc", "123", 0)

	result := unlockScript.Run(Ins.Context(), Ins, []string{"abc"}, []string{"123"})
	fmt.Println(result.Result())
}
