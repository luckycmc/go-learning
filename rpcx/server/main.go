package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	service "mysql/rpcx"
)

var addr = flag.String("addr", "0.0.0.0:8003", "server address")

func main() {
	s := server.NewServer()
	s.Register(new(service.Arith), "")
	s.Serve("tcp", *addr)
}
