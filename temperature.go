package main

import "fmt"

type Temperature struct {
	Fahrenheit float64
	Celsius    float64
}

func main() {
	p := ToFahrenheit(37.3)
	fmt.Println(*p)
	q := ToCelsius(77)
	fmt.Println(*q)
}

func ToFahrenheit(celsius float64) *Temperature {
	return &Temperature{
		Fahrenheit: celsius*1.8 + 32,
		Celsius:    celsius,
	}
}
func ToCelsius(fahrenheit float64) *Temperature {
	return &Temperature{
		Fahrenheit: fahrenheit,
		Celsius:    (fahrenheit - 32) / 1.8,
	}
}
