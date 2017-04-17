package main

import "fmt"

func main() {
	 fmt.Println(evenSteven(5))
	r := evenFunc(5)
	fmt.Println(r(5))
	h := func (num float64) (float64, bool) {
	// return num/2, int(num)%2 == 0
		if int(num)%2 == 0 {
			return num/2, true
		}
		return num/2, false
	}
	fmt.Println(h(5))
	lst := []int{51,2,16,20,43}
	fmt.Println(greatest(lst...))
	fmt.Println(foo(1,2),foo(1,2,3),foo(lst...),foo()) // *here
}

// 1. Write a function which takes an integer returns arg divided by 2 and returns bool shows
// whether the arg is even or not
func evenSteven(num float64) (float64, bool) {
	// return num/2, int(num)%2 == 0
	if int(num)%2 == 0 {
		return num/2, true
	}
	return num/2, false
}

// 2. use func expression to do the above...
func evenFunc(num float64) func(float64) string {
	 return func(float64) string {
		if int(num)%2 == 0 {
			return fmt.Sprintf("%v - %v",num/2,true)
		}
		 return fmt.Sprintf("%v - %v",num/2,false)
	}
}

// 3. find greatest number in list of numbers
func greatest(num ...int) int {
	var greatest int
	for _, element := range num {
		if element > greatest {
			greatest = element
		}
	}
	return greatest
}

// 4. write foo that can be called in all of these ways*
func foo(num ...int) int {
	var t int
	for _, element := range num {
		t += element
	}
	return t
}
