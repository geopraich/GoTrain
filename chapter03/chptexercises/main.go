package main

import (
	"fmt"
	"strings"
)

func main() {
	helloWorld()
	namePlease()
	remainderDiv()
	evens()
	fizzBuzz()
	mulOfThreeFive()
}

// 1. Create a func which prints "Hello World" to the terminal
// 2. Modify func so it prints hello and your name
func helloWorld() {
	fmt.Println("Hello World")
	name := "your name"
	fmt.Printf("Hello, my name is %v\n", name)
}

// 3. Create func that prints to terminal asking for users name then print hello <INPUT>
func namePlease() {
	var name string
	fmt.Print("Please Enter Name Here: ")
	fmt.Scan(&name)
	fmt.Printf("Hello %v\n", strings.Title(name))
}

// 4. Create func that prints to terminal asking for small and large number divide print remainder
func remainderDiv() {
	var (
		numLarge int64
		numSmall int64
	)
	fmt.Print("Please Enter Large Number Here: ")
	fmt.Scan(&numLarge)
	fmt.Print("Please Enter Small Number Here: ")
	fmt.Scan(&numSmall)
	for numSmall >= numLarge {
		fmt.Print("Number Needs to be Smaller: ")
		fmt.Scan(&numSmall)
		if numSmall < numLarge{
			fmt.Println("Thanks!")
			break
		}
	}
	fmt.Printf("remainder of %v divided by %v is %v", numLarge, numSmall, numLarge%numSmall)
}

// 5. Print all even numbers to terminal 0 - 100
func evens() {
	for c := 0; c <= 100; c++ {
		if c%2 == 0 {
			fmt.Println(c)
		}
	}
}

// 6. FizzBuzz - Fizz for multiples of 3 - Buzz for multiples of 5 - fizzbuzz for both 1-100
func fizzBuzz() {
	for c := 1; c <= 100; c++ {
		if c%15 == 0 {
			fmt.Println("FizzBuzz -",c)
		}else if c%3 == 0 {
			fmt.Println("Fizz -", c)
		}else if c%5 == 0 {
			fmt.Println("Buzz -", c)
		}else {
			fmt.Println(c)
		}
	}
}

// 7. find the sum multiples of 3 and 5 below 1000
func mulOfThreeFive() {
	t := 0
	for c := 0; c < 1000; c++ {
		if c%3 == 0 {
			t+=c
		}else if c%5 == 0 {
			t+=c
		}
	}
	fmt.Println(t)
}