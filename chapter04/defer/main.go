package main

import "fmt"

func main() {
	// without defer would run world() then hello()
	defer world()  // defers the run order of function until right before main() exits
	hello()
}

func hello() {
	fmt.Println("hello")
}

func world() {
	fmt.Println("world")
}
