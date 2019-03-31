package main

import (
	"fmt"

	"github.com/yasukotelin/go-designpattern/5_singleton/singleton"
)

func main() {
	obj1 := singleton.NewSingleton()
	obj2 := singleton.NewSingleton()

	if obj1 == obj2 {
		fmt.Println("シングルトンです")
	} else {
		fmt.Println("シングルトンではありません")
	}
}
