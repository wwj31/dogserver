package common

import (
	"testing"
)

func TestUids(t *testing.T) {
	uid := New()
	m := make(map[uint64]bool)
	for {
		u := uid.Uuid()
		if m[u] {
			panic(u)
		}
		m[u] = true
	}
}
