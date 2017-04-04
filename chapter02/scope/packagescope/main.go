package main

import "fmt"

var x int = 42  //  x at package level scope - accessible to whole package

func main() {
	fmt.Println(x)
	foo()
	y := 3.14  // y at block level scope
	fmt.Println(y)
	fmt.Println(z)
}

func foo() {
	fmt.Println(x)
}

var z int = 42  // package level accessible throughout package - order not important
