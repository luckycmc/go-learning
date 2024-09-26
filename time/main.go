package main

import (
	"fmt"
	"time"
)

func main() {
	// 基本使用
	timer1 := time.NewTimer(2 * time.Second)
	t1 := time.Now()
	fmt.Println(t1)
	t2 := <-timer1.C
	fmt.Println(t2)

	// 验证timer只能响应一次
	// timer2 := time.NewTimer(time.Second)
	// for {
	//	<-timer2.C
	//	fmt.Println("时间到")
	// }

	// timer实现延时功能
	time.Sleep(time.Second)
	timer3 := time.NewTimer(2 * time.Second)
	<-timer3.C
	fmt.Println("2秒到")
	<-time.After(2 * time.Second)
	fmt.Println("2秒到")

	// 停止定时器
	timer4 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer4.C
		fmt.Println("定时器执行")
	}()
	b := timer4.Stop()
	if b {
		fmt.Println("timer4关闭")
	}

	// 重置定时器
	timer5 := time.NewTimer(3 * time.Second)
	timer5.Reset(1 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer5.C)
	for {

	}
}
