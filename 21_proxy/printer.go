package main

import (
	"fmt"
	"time"
)

type Printable interface {
	SetPrinterName(string)
	GetPrinterName() string
	Print(string)
}

type Printer struct {
	name string
}

func NewPrinter(name string) *Printer {
	// インスタンス生成に時間がかかる
	heavyJob("Printerインスタンス生成中")
	return &Printer{
		name: name,
	}
}

func heavyJob(msg string) {
	fmt.Print(msg)
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
		fmt.Print(".")
	}
	fmt.Println()
	fmt.Println("完了")
}

func (p *Printer) GetPrinterName() string {
	return p.name
}

func (p *Printer) SetPrinterName(name string) {
	p.name = name
}

func (p *Printer) Print(s string) {
	fmt.Printf("=== %s ===\n", p.name)
	fmt.Println(s)
}
