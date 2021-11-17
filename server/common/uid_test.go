package common

import (
	"testing"
)

func TestUids(t *testing.T) {
	uid := NewUID()
	m := make(map[uint64]bool)
	for {
		u := uid.Uuid()
		if m[u] {
			panic(u)
		}
		m[u] = true
	}
}
