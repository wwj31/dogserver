package rds

import (
	"fmt"
	"log"
	"sync"
	"testing"
)

func TestLock(t *testing.T) {
	if err := Connect("redis://:123456@localhost:6379/1", false); err != nil {
		log.Fatalf(err.Error())
	}

	var n int

	waiter := sync.WaitGroup{}
	waiter.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			LockDo("n", func() {
				n++
			})
			waiter.Done()
		}()
	}
	waiter.Wait()
	fmt.Println(n)
}
