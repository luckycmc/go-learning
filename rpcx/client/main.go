package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/share"
	"log"
	service "mysql/rpcx"
	"time"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()
	share.Trace = true
	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()
	args := &service.Args{
		A: 10,
		B: 20,
	}
	for {
		reply := &service.Reply{}
		err := xclient.Call(context.Background(), "Multiply", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}
		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		time.Sleep(1e9)
	}
}