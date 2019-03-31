package idcard

import (
	"fmt"

	"github.com/yasukotelin/go-designpattern/4_factory_method/framework"
)

type IDCard struct {
	Owner string
}

func NewIDCard(owner string) *IDCard {
	fmt.Printf("%sのカードを作ります\n", owner)
	return &IDCard{
		Owner: owner,
	}
}

func (card *IDCard) Use() {
	fmt.Printf("%sのカードを使います\n", card.Owner)
}

type IDCardFactory struct {
	Owners []string
}

func (f IDCardFactory) Create(owner string) framework.Product {
	return framework.Create(owner, f.CreateProduct, f.RegisterProduct)
}

func (f IDCardFactory) CreateProduct(owner string) framework.Product {
	return NewIDCard(owner)
}

func (f IDCardFactory) RegisterProduct(p framework.Product) {
	if v, ok := p.(*IDCard); ok {
		f.Owners = append(f.Owners, v.Owner)
	}
}
