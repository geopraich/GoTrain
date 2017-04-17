package main

import "fmt"

// assertion can only be used with interfaces
func main() {
	var name interface{} = "Sydney"
	str, ok := name.(string)  // asserts to string var.(assert type) while conversion type(var)
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("value not a string\n")
	}
}