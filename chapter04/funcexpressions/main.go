package main

import "fmt"

func main() {
	// only way to get a function inside another function
	greeting := func() {  // function expression (anon func)
		fmt.Println("Hello World!")
	}

	greeting()  // calls function expression (anon func)
	greet := makeMoreFunc()
	fmt.Println(greet())  // prints out the return of func expression

	func() {  // this is an anonymous self executing function (can have params)
		fmt.Println("I'm Driving!")
	}()  // this parenthesis set executes the function
}

func makeMoreFunc() func() string {
	return func() string {
		return "Hello World"
	}
}