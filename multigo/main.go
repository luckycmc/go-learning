package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	// goroutine结束就登记-1
	defer wg.Done()
	fmt.Println("hello goroutine", i)
}

func main() {
	for i := 0; i < 10; i++ {
		// 启动一个goroutine就登记+1
		wg.Add(1)
		go hello(i)
	}
	// 等待所有登记的goroutine都结束
	wg.Wait()
}
