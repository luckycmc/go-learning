package main

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"os"
)

func main() {
	qrCode, _ := qr.Encode("http://www.baidu.com", qr.M, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 256, 256)
	file, _ := os.Create("qrcode.png")
	defer file.Close()

	png.Encode(file, qrCode)
}
