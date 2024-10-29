package main

import (
	"github.com/jung-kurt/gofpdf"
)

func main() {
	err := GeneratePdf("hello.pdf")
	if err != nil {
		panic(err)
	}
}

func GeneratePdf(filename string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)

	pdf.CellFormat(190, 7, "hello", "0", 0, "CM", false, 0, "")
	pdf.ImageOptions(
		"laradock-full-logo.jpg",
		75.89, 18.41,
		0, 0,
		false,
		gofpdf.ImageOptions{
			ImageType: "JPG",
			ReadDpi:   true,
		},
		0,
		"",
	)
	return pdf.OutputFileAndClose(filename)
}
