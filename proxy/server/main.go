package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	u, _ := url.Parse("http://localhost:5051/")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)

}

func main() {
	http.HandleFunc("/test_proxy", sayHello)
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
