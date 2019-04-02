package main

import (
	"fmt"

	"github.com/yasukotelin/go-designpattern/07_builder/builder"
)

func main() {
	// plane text
	textBuilder := new(builder.TextBuilder)
	director := NewDirector(textBuilder)
	director.Construct()
	res := textBuilder.GetResult()
	fmt.Println(res)

	// HTML fiel
	htmlBuilder, err := builder.NewHTMLBuilder("C:\\Users\\yasu\\Desktop\\out.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	director = NewDirector(htmlBuilder)
	director.Construct()
	res = htmlBuilder.GetResult()
	fmt.Printf("%sにHTMLを出力しました\n", res)
}
