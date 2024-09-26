package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取ticker对象
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	// 子协程
	go func() {
		for {
			<-ticker.C
			i++
			fmt.Println(i)
			if i == 5 {
				// 停止
				ticker.Stop()
			}
		}
	}()
	for {

	}
}
