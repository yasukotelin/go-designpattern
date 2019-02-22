package main

import (
	"github.com/yasukotelin/go-designpattern/strategy/duck"
	"github.com/yasukotelin/go-designpattern/strategy/fly"
)

func main() {
	// マガモ
	mDuck := duck.NewMallardDuck()
	show(mDuck)

	// 模型のアヒル
	modelDuck := duck.NewModelDuck()
	show(modelDuck)

	// アルゴリズムの動的設定
	// 模型のアヒルを飛べるようにする
	modelDuck.FlyBehavior = &fly.FlyWithRocketBehavior{}
	show(modelDuck)
}

func show(d duck.Duck) {
	d.Display()
	d.Quack()
	d.Fly()
}
