package beverage

// Beverage is interface
type Beverage interface {
	Description() string
	Cost() float64
}

// Espresso is implements Beverage interface
type Espresso struct{}

func (e *Espresso) Description() string {
	return "エスプレッソ"
}

func (e *Espresso) Cost() float64 {
	return 1.99
}

// HouseBlend is implements Beverage interface
type HouseBlend struct{}

func (h *HouseBlend) Description() string {
	return "ハウスブレンドコーヒー"
}

func (h *HouseBlend) Cost() float64 {
	return 0.89
}
