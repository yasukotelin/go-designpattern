package observer

type Observer interface {
	Update(templ float64, humidity float64, pressure float64)
}
