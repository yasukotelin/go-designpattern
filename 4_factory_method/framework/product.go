package framework

type Product interface {
	Use()
}

type Factory interface {
	Create(owner string) Product
	CreateProduct(owner string) Product
	RegisterProduct(p Product)
}

// Create は Factory interfaceの共通Create関数
func Create(owner string, createProduct func(string) Product, registerProduct func(p Product)) Product {
	p := createProduct(owner)
	registerProduct(p)
	return p
}
