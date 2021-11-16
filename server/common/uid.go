package common

import (
	"time"
)

const max = 100000

type UID struct {
	ids    []uint64
	lastAt int64
}

func New() *UID {
	uid := &UID{
		ids:    make([]uint64, 0, max),
		lastAt: time.Now().UnixNano(),
	}
	time.Sleep(500 * time.Millisecond)
	uid.gen()
	return uid
}

func (s *UID) Uuid() (uuid uint64) {
	if len(s.ids) == 0 {
		time.Sleep(1 * time.Second)
		s.gen()
	}

	last := len(s.ids) - 1
	uuid = s.ids[last]
	s.ids = s.ids[:last]

	if len(s.ids) < max/2 && s.lastAt < time.Now().UnixNano() {
		s.gen()
	}

	return uuid
}

func (s *UID) gen() {
	span := time.Now().UnixNano() - s.lastAt
	for i := int64(0); i < span; i++ {
		s.lastAt++
		s.ids = append(s.ids, uint64(s.lastAt))
		if len(s.ids) >= max {
			return
		}
	}
}
