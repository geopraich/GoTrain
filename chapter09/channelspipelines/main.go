package main

import "fmt"

func main() {
	// Set up the pipeline
	c := gen(2, 3)
	out := sq(c)

	// consume the output
	fmt.Println(<-out)
	fmt.Println(<-out)

	for ns := range sq(sq(gen(2,3))) {  // sq(gen(2,3)) returns 4,9 sq(gen(4,9)) returns 16 and 81
		fmt.Println(ns)
	}
}

func gen(num ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range num {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {  // could make receive only <-chan int
	out := make(chan int)
	go func() {
		for n := range in {  // on channel 2 and 3 (int)
			out <- n * n  // n * n will be first n 2*2 by itself and then second n 3*3
		}
		close(out)
	}()
	return out
}
