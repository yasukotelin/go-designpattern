package singleton

import "fmt"

var singleInstance *singleton

func init() {
	singleInstance = newSingleton()
}

type singleton struct{}

func newSingleton() *singleton {
	fmt.Println("インスタンスを生成しました")
	return new(singleton)
}

func NewSingleton() *singleton {
	return singleInstance
}
