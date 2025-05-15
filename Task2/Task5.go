package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var count int = 0
var countWithNoLock int64 = 0

func main() {
	var syncLock sync.Mutex
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				syncLock.Lock()
				count++
				syncLock.Unlock()
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("count = %d\n", count)
	fmt.Println("-------------------------------------------")

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&countWithNoLock, 1)
			}
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Printf("countWithNoLock = %d", countWithNoLock)
}
