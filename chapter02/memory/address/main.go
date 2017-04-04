package main

import "fmt"

func main() {

	s := 42

	fmt.Println("s -", s)
	fmt.Println("s's memory address is -", &s)  // shows address of var s (single &)
	fmt.Printf("s address in decimal - %d\n", &s)
	fmt.Printf("s address in bits - %b", &s)
}