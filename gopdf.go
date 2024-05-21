package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jung-kurt/gofpdf"
)


func main() {
// fmt.Println("hello")
file := "txtForPdf.txt"
content, err := os.ReadFile(file)
if err != nil {
	log.Fatalf("%s file not found", file)
} else {
	fmt.Println("read file")
}

pdf := gofpdf.New("P", "mm", "A4", "")

pdf.AddPage()
pdf.SetFont("Arial", "B", 14)
pdf.MultiCell(190.0, 277.0, string(content), "", "", false)

err = pdf.OutputFileAndClose("test.pdf")

if err != nil {
	log.Fatal("failed to make pdf")
} else {
	fmt.Println("Created text.pdf")
}

}