package main

import "fmt"

type Display interface {
	Open()
	Print()
	Close()
	Display()
}

// display はDisplay interface のテンプレートメソッド
func display(open, print, close func()) {
	open()
	for i := 0; i < 5; i++ {
		print()
	}
	close()
}

type RuneDisplay struct {
	r rune
}

func NewRuneDisplay(r rune) *RuneDisplay {
	return &RuneDisplay{
		r: r,
	}
}

func (rd *RuneDisplay) Open() {
	fmt.Print("<<")
}

func (rd *RuneDisplay) Print() {
	fmt.Print(string(rd.r))
}

func (rd *RuneDisplay) Close() {
	fmt.Println(">>")
}

func (rd *RuneDisplay) Display() {
	display(rd.Open, rd.Print, rd.Close)
}

type StringDisplay struct {
	s     string
	width int
}

func NewStringDisplay(s string) *StringDisplay {
	return &StringDisplay{
		s:     s,
		width: len(s),
	}
}

func (sd *StringDisplay) Open() {
	sd.PrintLine()
}

func (sd *StringDisplay) Print() {
	fmt.Printf("|%s|\n", sd.s)
}

func (sd *StringDisplay) Close() {
	sd.PrintLine()
}

func (sd *StringDisplay) Display() {
	display(sd.Open, sd.Print, sd.Close)
}

func (sd *StringDisplay) PrintLine() {
	fmt.Print("+")
	for i := 0; i < sd.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
