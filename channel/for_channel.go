package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0-100的数发送到cha1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中取值，并将该值的平方发送到ch2中
	go func() {
		for i := 0; i < 100; i++ {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中获取值并打印
	for i := range ch2 {
		fmt.Println(i)
	}
}
