package redis

import (
	"fmt"
	"sync"
	"testing"
)

func TestLock(t *testing.T) {
	NewBuilder().OnConnect(func() {
		fmt.Println("redis connect success")
	}).Connect()

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
