package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/test", myHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "connect success")
	fmt.Println("method:", r.Method)
	fmt.Println("url:", r.URL)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	// response
	w.Write([]byte("this is a test"))
}
