package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine %d\n", i)
			time.Sleep(time.Second)
			if i == 2 {
				break
			}
		}
	}()
}
