package main

import (
	"github.com/yasukotelin/go-designpattern/observer/observer"
)

type WeatherData struct {
	observers []observer.Observer

	Temperature float64
	Humidity    float64
	Pressure    float64
}

func (w *WeatherData) Register(o observer.Observer) {
	w.observers = append(w.observers, o)
}

func (w *WeatherData) Remove(o observer.Observer) {
	var result []observer.Observer
	for _, e := range w.observers {
		if e != o {
			result = append(result, e)
		}
	}
	w.observers = result
}

func (w *WeatherData) Notify() {
	for _, o := range w.observers {
		o.Update(w.Temperature, w.Humidity, w.Pressure)
	}
}

func (w *WeatherData) SetMeasurements(temperature float64, humidity float64, pressure float64) {
	w.Temperature = temperature
	w.Humidity = humidity
	w.Pressure = pressure

	w.Notify()
}
