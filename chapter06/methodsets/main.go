package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

// method with a value receiver (c circle) value receiver
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	side float64
}

// method with pointer receiver
func (a *square) area() float64 {
	return a.side * a.side
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := square{5}
	info(&c)  // for pointer receiver needs address
	b := circle{5}
	info(b)  // value receiver can use both address or not
	info(&b)
}
