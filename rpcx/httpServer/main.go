package main

import (
	"flag"
	"fmt"
	"github.com/smallnest/rpcx/server"
	"log"
	service "mysql/rpcx"
	"net/http"
)

var (
	addr = flag.String("addr", "192.168.72.130:8003", "server address")
)

func main() {
	flag.Parse()
	s := server.NewServer()
	err := s.Register(new(service.Arith), "")
	if err != nil {
		log.Println("error")
		return
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		s.ServeHTTP(writer, request)
	})
	// 启动 HTTP 服务
	addr := "0.0.0.0:8003"
	fmt.Printf("Starting HTTP server on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
