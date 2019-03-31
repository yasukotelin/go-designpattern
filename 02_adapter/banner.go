package main

import "fmt"

// Banner はすでに存在している構造体とする
type Banner struct {
	s string
}

func NewBanner(s string) *Banner {
	return &Banner{
		s: s,
	}
}

func (b *Banner) ShowWithParen() {
	fmt.Printf("(%s)\n", b.s)
}

func (b *Banner) ShowWithAster() {
	fmt.Printf("*%s*\n", b.s)
}
