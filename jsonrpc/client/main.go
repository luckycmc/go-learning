package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Params struct {
	Width, Height int
}

func main() {
	conn, err := jsonrpc.Dial("tcp", ":1234")
	if err != nil {
		log.Panicln(err)
	}
	ret := 0
	err2 := conn.Call("Rect.Area", Params{10, 2}, &ret)
	if err2 != nil {
		log.Panicln(err2)
	}
	fmt.Printf("面积为%d", ret)
	err3 := conn.Call("Rect.Perimeter", Params{10, 2}, &ret)
	if err3 != nil {
		log.Panicln(err3)
	}
	fmt.Printf("周长为%d", ret)
}
