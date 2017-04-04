package main

import "fmt"


func main() {
	// zero value declarations
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("%v - %v - %v -%v\n", a, b, c, d)
	fmt.Printf("%T - %T - %T - %T\n", a ,b ,c ,d)
}
