package main

import (
	"fmt"
)

type ConditionDisplay struct {
	Temperature float64
	Humidity    float64
}

func (c *ConditionDisplay) Update(templ float64, humidity float64, pressure float64) {
	c.Temperature = templ
	c.Humidity = humidity

	c.Display()
}

func (c *ConditionDisplay) Display() {
	fmt.Printf("現在の気象状況: 温度%v度 湿度%v％\n", c.Temperature, c.Humidity)
}
