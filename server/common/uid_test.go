package common

import (
	"fmt"
	"testing"
	"time"
)

func TestUids(t *testing.T) {
	uid := NewUID(1)
	m := make(map[uint64]bool)
	for i := 0; i < 10000; i++ {
		u := uid.GenUuid()
		if m[u] {
			panic(u)
		}
		m[u] = true
		fmt.Println(u>>24, time.Now().UnixMilli())
		time.Sleep(1 * time.Millisecond)
	}
}
