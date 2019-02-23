package main

func main() {
	wd := WeatherData{}
	wd.Register(&ConditionDisplay{})

	wd.SetMeasurements(27, 65, 30.0)
	wd.SetMeasurements(28, 70, 29.2)
}
