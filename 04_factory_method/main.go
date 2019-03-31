package main

import (
	"github.com/yasukotelin/go-designpattern/4_factory_method/framework"
	"github.com/yasukotelin/go-designpattern/4_factory_method/idcard"
)

func main() {
	factory := new(idcard.IDCardFactory)
	cards := []framework.Product{
		factory.Create("taro tanaka"),
		factory.Create("hanako yamada"),
		factory.Create("yoshinori morita"),
	}

	for _, card := range cards {
		card.Use()
	}
}
