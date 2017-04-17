package main

import "fmt"

func main() {
	greet("Mark")  // greet() called and pass argument of string type -"Mark"-
	greetTwo("Mark","Jenny")  // greetTwo() called passed to args
}

func greet(nam string) {  // func greet([parameter_name-nam] [of_type_string])
	fmt.Println(nam)  // prints out the argument passed that is set to param nam
}
// this func takes two arguments of type string therefore has two parameters set that take type string
func greetTwo(nameOne string, nameTwo string) {  // could also be written func greetTwo(nameOne, nameTwo string)
	fmt.Println(nameOne, nameTwo)
}
