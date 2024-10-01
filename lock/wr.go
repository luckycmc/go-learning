package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x1     int64
	wg1    sync.WaitGroup
	rwlock sync.RWMutex
)

func write() {
	rwlock.Lock()
	x1 = x1 + 1
	time.Sleep(10 * time.Millisecond)
	rwlock.Unlock()
	wg1.Done()
}

func read() {
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
	wg1.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg1.Add(1)
		go read()
	}
	wg1.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
