package main

import "fmt"

// formal declarations - outside/inside func

var a string = "golang"  // package scope
var b int = 10
var c float64 = 3.14
var d bool = false
var f int
var g = 43  // not rec

func main() {
	var e string = "world"
	f = 20
	var h = "hello"

	fmt.Printf("%v - %v - %v - %v - %v - %v - %v - %v\n", a, b, c, d, e, f, g, h)
	fmt.Printf("%T - %T - %T - %T - %T - %T - %T - %T\n", a, b, c, d, e, f, g, h)
}
