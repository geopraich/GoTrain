package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

// recursive function calls itself
func factorial(num int) int {
	if num == 0 {  // breaks recursion loop by returning 1
		return 1
	}
	return num * factorial(num-1)  // recursive call - 4*3*2*1*1=24 first call then second etc...
}
