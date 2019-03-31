package main

import "fmt"

func main() {
	printer := NewPrinterProxy("Alice")

	fmt.Printf("名前は現在%sです\n", printer.GetPrinterName())
	printer.SetPrinterName("Bob")
	fmt.Printf("名前は現在%sです\n", printer.GetPrinterName())

	printer.Print("Hello, world")
}
