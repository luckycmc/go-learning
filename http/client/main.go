package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8000/test")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	// 200 OK
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)

	buf := make([]byte, 1024)
	for {
		// receive data from server
		n, err := resp.Body.Read(buf)

		if err != nil && err != io.EOF {
			// EOF means we're done
			fmt.Println("error:", err)
			return
		} else {
			fmt.Println("read done")
			res := string(buf[:n])
			fmt.Println("Bytes Read:", res)
			break
		}
	}
}
