package main

import "fmt"

func main() {
	// 创建有缓冲区的通道
	ch := make(chan int, 1)
	ch <- 10
	fmt.Println("发送成功")
}
