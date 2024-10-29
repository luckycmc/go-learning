package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
)

func main() {
	err := qrcode.WriteFile("http://www.topgoer.com", qrcode.Medium, 256, "qr.png")
	if err != nil {
		fmt.Println("write err")
	}
}
