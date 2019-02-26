package beverage

// Decorator is interface for decorate Beverage
type Decorator interface {
	Description() string
	Cost() float64
}

// Mocha decorator struct
type Mocha struct {
	Beverage Beverage
}

func (m *Mocha) Description() string {
	return "モカ" + m.Beverage.Description()
}

func (m *Mocha) Cost() float64 {
	return m.Beverage.Cost() + 0.20
}

// Soy decorator struct
type Soy struct {
	Beverage Beverage
}

func (s *Soy) Description() string {
	return "豆乳" + s.Beverage.Description()
}

func (s *Soy) Cost() float64 {
	return s.Beverage.Cost() + 0.15
}
