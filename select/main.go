package main

import (
	"fmt"
	"time"
)

func test1(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test1"
}

func test2(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test2"
}

func main() {
	// 2个管道
	output1 := make(chan string)
	output2 := make(chan string)
	// 开启两个子协程，向管道内写数据
	go test1(output1)
	go test2(output2)
	// 使用select监控
	select {
	case s1 := <-output1:
		fmt.Println("s1=", s1)
	case s2 := <-output2:
		fmt.Println("s2=", s2)
	}
}
