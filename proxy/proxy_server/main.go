package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("我是5051")
}

func main() {
	http.HandleFunc("/test_proxy", sayHello)
	err := http.ListenAndServe(":5051", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
