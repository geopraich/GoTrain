package main

import "fmt"

func main() {
	operandExample()
	operandExample2()
}

func operandExample() {
	if true || false {  // or operand: ||
		fmt.Println("This runs")
	}
}

func operandExample2() {
	if true && false {  // and operand: &&
		fmt.Println("Not running")
	}else {
		fmt.Println("This runs")
	}
}