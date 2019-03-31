package main

import "fmt"

type Product interface {
	Use(string)
	CreateClone() Product
}

type Manager struct {
	showcase map[string]Product
}

func NewManager() *Manager {
	return &Manager{
		showcase: make(map[string]Product),
	}
}

func (m *Manager) Register(name string, product Product) {
	m.showcase[name] = product
}

func (m *Manager) Create(protoname string) Product {
	return m.showcase[protoname].CreateClone()
}

type MessageBox struct {
	decoRune rune
}

func NewMessageBox(r rune) *MessageBox {
	return &MessageBox{
		decoRune: r,
	}
}

func (mb *MessageBox) Use(s string) {
	decostr := string(mb.decoRune)
	for i := 0; i < len(s)+4; i++ {
		fmt.Print(decostr)
	}
	fmt.Println()
	fmt.Printf("%s %s %s\n", decostr, s, decostr)
	for i := 0; i < len(s)+4; i++ {
		fmt.Print(decostr)
	}
	fmt.Println()
}

func (mb *MessageBox) CreateClone() Product {
	clone := *mb
	return &clone
}

type UnderlinePen struct {
	urrune rune
}

func NewUnderlinePen(r rune) *UnderlinePen {
	return &UnderlinePen{
		urrune: r,
	}
}

func (ulp *UnderlinePen) Use(s string) {
	urstr := string(ulp.urrune)

	fmt.Printf("\"%s\"\n", s)
	fmt.Print(" ")
	for i := 0; i < len(s); i++ {
		fmt.Print(urstr)
	}
	fmt.Println()
}

func (ulp *UnderlinePen) CreateClone() Product {
	clone := *ulp
	return &clone
}
