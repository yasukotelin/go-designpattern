package duck

import (
	"fmt"

	"github.com/yasukotelin/go-designpattern/strategy/fly"
	"github.com/yasukotelin/go-designpattern/strategy/quack"
)

type MallardDuck struct {
	Name          string
	QuackBehavior quack.QuackBehavior
	FlyBehavior   fly.FlyBehavior
}

func NewMallardDuck() *MallardDuck {
	return &MallardDuck{
		Name:          "マガモ",
		QuackBehavior: &quack.StdQuackBehavior{},
		FlyBehavior:   &fly.StdFlyBehavior{},
	}
}

func (m *MallardDuck) Display() {
	fmt.Printf("=== %s ===\n", m.Name)
}

func (m *MallardDuck) Quack() {
	m.QuackBehavior.Quack()
}

func (m *MallardDuck) Fly() {
	m.FlyBehavior.Fly()
}
