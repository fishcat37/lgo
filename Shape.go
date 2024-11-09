package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Circle struct {
	r float64
}
type Rectangle struct {
	length, width float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}
func (r Rectangle) Area() float64 {
	return r.length * r.width
}
func Play(a Shape) float64 {
	return a.Area()
}
func main() {
	c := Circle{
		4,
	}
	area := Play(c)
	fmt.Println(area)
	b := Rectangle{
		3,
		4,
	}
	area = Play(b)
	fmt.Println(area)
}
