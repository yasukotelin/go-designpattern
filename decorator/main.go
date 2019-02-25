package main

import (
	"fmt"

	"github.com/yasukotelin/go-designpattern/decorator/beverage"
)

func main() {

	// No decoration
	espresso := beverage.Espresso{}
	printBeverage(espresso)

	// Decoration
	mochaEspresso := beverage.Mocha{
		Beverage: espresso,
	}
	printBeverage(mochaEspresso)

	// More decoration
	soyMochaEspresso := beverage.Soy{
		Beverage: mochaEspresso,
	}
	printBeverage(soyMochaEspresso)
}

func printBeverage(b beverage.Beverage) {
	fmt.Printf("商品名: %v\n", b.Description())
	fmt.Printf("値段: %v\n", b.Cost())
}
