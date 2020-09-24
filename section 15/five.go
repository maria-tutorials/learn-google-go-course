package main

import (
	"fmt"
	"math"
)

type square struct {
	lenght float32
	width  float32
}

type circle struct {
	radius float32
}

type shape interface {
	area() float32
}

func (s square) area() float32 {
	return s.lenght * s.width
}

func (c circle) area() float32 {
	return math.Pi * (c.radius * c.radius)
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{
		lenght: 3.0,
		width:  5.0,
	}

	c := circle{
		radius: 4.0,
	}

	info(s)
	info(c)
}
