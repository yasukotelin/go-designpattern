package main

func main() {
	displays := []Display{
		NewRuneDisplay('A'),
		NewStringDisplay("hello, golang"),
	}

	for _, d := range displays {
		d.Display()
	}
}
