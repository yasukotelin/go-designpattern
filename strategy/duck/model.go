package duck

import (
	"fmt"

	"github.com/yasukotelin/go-designpattern/strategy/fly"
	"github.com/yasukotelin/go-designpattern/strategy/quack"
)

type ModelDuck struct {
	Name          string
	QuackBehavior quack.QuackBehavior
	FlyBehavior   fly.FlyBehavior
}

func NewModelDuck() *ModelDuck {
	return &ModelDuck{
		Name:          "模型のアヒル",
		QuackBehavior: &quack.MuteQuackBehavior{},
		FlyBehavior:   &fly.FlyNoWayBehavior{},
	}
}

func (m *ModelDuck) Display() {
	fmt.Printf("=== %s ===\n", m.Name)
}

func (m *ModelDuck) Quack() {
	m.QuackBehavior.Quack()
}

func (m *ModelDuck) Fly() {
	m.FlyBehavior.Fly()
}
