package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello world")
}

func main() {
	go hello()
	fmt.Println("main goroutine done")
	time.Sleep(time.Second)
}
