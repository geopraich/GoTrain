package main

import "fmt"

// functional programming
func main() {
	visit([]int{1,2,3,4}, func(num int) {  // arg1: pass a slice/list of int, arg2: pass a anon func that takes
		// an int
		fmt.Println(num)  // *which then prints out via num from anon func
	})  // needs to all go inside function parenthesis

	xs := filter([]int{1,2,3,4}, func(num int) bool {
		return num > 1  // ^greater than 1
	})
	fmt.Println(xs)
}

// a callback is when you pass a function as an argument
func visit(numbers []int, callback func(int)) {  // param: name:numbers ofType:[]int, param2: name:callback ofType:func
	//() which also takes an arg of type int
	for _, element := range numbers {  // iterates through the []int{1,2,3,4}
		callback(element)  // calls callback function that passes element of type int*
	}
}

//another callback example
func filter(numbers []int, callbackName func(n int) bool) []int {  // callbackName function takes int returns a boolean
	xs := []int{}  // creates empty list/slice
	for _, element := range numbers {
		if callbackName(element) {  // calls function returns true if element > 1^
			xs = append(xs, element)  // appends to empty list/slice if^
		}
	}
	return xs  // returns a []int as stated by filter func declaration
}
