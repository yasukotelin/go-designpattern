package quack

import "fmt"

type QuackBehavior interface {
	Quack()
}

type StdQuackBehavior struct {
}

func (q *StdQuackBehavior) Quack() {
	fmt.Println("がーがー")
}

type MuteQuackBehavior struct {
}

func (q *MuteQuackBehavior) Quack() {
	fmt.Println("...")
}
