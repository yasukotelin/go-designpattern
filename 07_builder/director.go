package main

import (
	"github.com/yasukotelin/go-designpattern/07_builder/builder"
)

type Director struct {
	builder builder.Builder
}

func NewDirector(builder builder.Builder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) Construct() {
	d.builder.MakeTitle("Greeting")
	d.builder.MakeString("朝から昼にかけて")
	d.builder.MakeItems([]string{
		"おはようございます",
		"こんにちは",
	})
	d.builder.MakeString("夜に")
	d.builder.MakeItems([]string{
		"こんばんは",
		"おやすみなさい",
		"さようなら",
	})
	d.builder.Close()
}
