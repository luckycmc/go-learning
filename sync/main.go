package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello() {
	defer wg.Done()
	fmt.Println("hello sync")
}

func main() {
	wg.Add(1)
	go hello()
	fmt.Println("hello main")
	wg.Wait()
}
