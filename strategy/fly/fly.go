package fly

import "fmt"

type FlyBehavior interface {
	Fly()
}

type StdFlyBehavior struct {
}

func (f *StdFlyBehavior) Fly() {
	fmt.Println("飛んでます")
}

type FlyNoWayBehavior struct {
}

func (f *FlyNoWayBehavior) Fly() {
	fmt.Println("飛べません")
}

type FlyWithRocketBehavior struct {
}

func (f *FlyWithRocketBehavior) Fly() {
	fmt.Println("ロケットで飛んでます")
}
