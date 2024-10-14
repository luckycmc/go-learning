package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type params struct {
	A, B int
}

type result struct {
	Pro, Quo, Rem int
}

func main() {
	conn, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	req := params{10, 2}
	var res result
	err2 := conn.Call("Rect.Product", req, &res)
	if err2 != nil {
		log.Fatal("call:", err)
	}
	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Quo)
	err3 := conn.Call("Rect.Reminder", req, &res)
	if err3 != nil {
		log.Fatal("call:", err)
	}
	fmt.Printf("%d / %d 商 %d, 余数 %d", req.A, req.B, res.Quo, res.Rem)
}
