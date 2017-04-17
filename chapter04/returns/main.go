package main

import "fmt"

func main() {
	fmt.Println(greet("Joe","Smith"))
	fmt.Println(greetAgain("Joe","Smith"))
	fmt.Println(greetMore("Joe","Smith"))
	fmt.Println(greetLast("Joe","Smith"))
}

// this function returns a string - func [receiver] nameOfFunc(params type) [returns(type)]
func greet(nameOne, nameTwo string) string {
	return fmt.Sprint(nameOne," ", nameTwo)  // Sprint concatenates and returns a string
}

// this function names the return set it to s - named return
func greetAgain(nameOne, nameTwo string) (s string) {
	s = fmt.Sprint(nameOne," ",nameTwo)
	return s  // can be written without the s and just return
}

// this function returns multiple values
func greetMore(nameOne, nameTwo string) (string, string) {
	return fmt.Sprint(nameOne," ", nameTwo), fmt.Sprint(nameTwo," ",nameOne)  // this returns two strings
}

// same as above but with named returns
func greetLast(nameOne, nameTwo string) (s string, b string) {
	s = fmt.Sprint(nameOne," ",nameTwo)
	b = fmt.Sprint(nameTwo," ",nameOne)
	return  // can also be written as - return s, b - or reverse order of return - return b, s
}
