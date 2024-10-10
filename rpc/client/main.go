package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 传参
type Params struct {
	Width, Height int
}

func main() {
	// 连接远程rpc服务
	conn, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// 调用方法
	// 面积
	ret := 0
	err2 := conn.Call("Rect.Area", Params{Width: 400, Height: 400}, &ret)
	if err2 != nil {
		log.Fatal("rect error:", err2)
	}
	fmt.Println("面积： ", ret)
	// 周长
	err3 := conn.Call("Rect.Perimeter", Params{Width: 400, Height: 400}, &ret)
	if err3 != nil {
		log.Fatal("rect error:", err3)
	}
	fmt.Println("周长: ", ret)
}
