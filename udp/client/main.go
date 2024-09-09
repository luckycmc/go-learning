package main

import (
	"fmt"
	"net"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 23301,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer socket.Close()
	sendData := []byte("Hello, Server!")
	_, err = socket.Write(sendData)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
}
