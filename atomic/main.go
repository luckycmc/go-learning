package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	x    int64
	lock sync.Mutex
	wg   sync.WaitGroup
)

// 普通加函数
func add() {
	x++
	wg.Done()
}

// 互斥锁加函数
func addMutex() {
	lock.Lock()
	x++
	lock.Unlock()
	wg.Done()
}

// atomic加函数
func addAtomic() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		//go add()
		//go addMutex()
		go addAtomic()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x)
	fmt.Println(end.Sub(start))
}
