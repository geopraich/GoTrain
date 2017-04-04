package main

import "fmt"


func main() {
	// shorthand declarations - can only be used inside func
	// infers type - syntax -
	a := 10
	b := "golang"
	c := 3.14
	d := true

	fmt.Printf("%v \n", a) // %v use default format for value
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%T - %T - %T - %T \n", a, b, c, d)  // %T shows type of variables
}