package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func isOpen(host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Millisecond * 1)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err == nil {
		_ = conn.Close()
		return true
	}
	return false
}

func main() {
	hostname := flag.String("hostname", "", "hostname")
	startPort := flag.Int("p", 80, "start port")
	endPort := flag.Int("e", 1000, "end port")
	timeout := flag.Duration("t", time.Millisecond*200, "timeout")
	flag.Parse()

	ports := []int{}
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	for port := *startPort; port <= *endPort; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen(*hostname, p, *timeout)
			if opened {
				mutex.Lock()
				ports = append(ports, p)
				mutex.Unlock()
			}
			wg.Done()
		}(port)
		wg.Wait()
		fmt.Printf("opened ports:%v\n", ports)
	}
}
