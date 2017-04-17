package main

import (
	"fmt"
	"math"
)

type Square struct {  // user defined type of type Square underlying struct type
	side float64  // field side type float64
}

// method for Square finds area of square
func (sqr Square) area() float64 {
	return sqr.side * sqr.side  // takes side float64 multiples by another side of sqr
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Square and Circle implement the Shape interface
type Shape interface{  // creates type Shape underlying type of interface
	area() float64  // signature method can be called by a Shape interface type - signature needs to match exactly
			// interfaces with Square type and its methods and Circle type and its method
}

func info(sqr Shape) {  // function that takes Shape type and prints out the area via Shape interface
	fmt.Println(sqr)
	fmt.Println(sqr.area())
}

func main() {
	s := Square{side:12}
	c := Circle{12}
	info(s)  // shape can be used on both circle and square
	info(c) // concrete type of c is Circle
}
