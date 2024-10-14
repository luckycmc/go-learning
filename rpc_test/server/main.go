package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
)

type Params struct {
	A, B int
}

// Response 返回给客户端的结果
type Response struct {
	Pro, Quo, Rem int
}

// Rect 用户注册的结构体
type Rect struct{}

// Product RPC 服务端方法，求和
func (r *Rect) Product(p Params, ret *Response) error {
	ret.Pro = p.A * p.B
	return nil
}

func (r *Rect) Reminder(p Params, ret *Response) error {
	if p.B == 0 {
		return errors.New("reminder failed")
	}
	// 除
	ret.Quo = p.A / p.B
	// 取模
	ret.Rem = p.A % p.B
	return nil
}

func main() {
	// 注册服务
	rect := new(Rect)
	// 注册一个rect服务
	rpc.Register(rect)
	// 绑定服务到http协议
	rpc.HandleHTTP()
	// 监听服务
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal(err)
	}
}
