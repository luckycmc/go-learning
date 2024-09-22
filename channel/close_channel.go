package main

import "fmt"

func main() {
	// 创建通道
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		// 关闭通道
		close(c)
	}()
	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("end")
}
