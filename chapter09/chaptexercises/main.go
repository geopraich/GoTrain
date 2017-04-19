package main

import "fmt"

func main() {
	in := gen()
	f := factorail(in)

	for n:= range f {
		fmt.Println(n)
	}
}

func gen() chan int {
	out := make(chan int)
	go func() {
		for c := 0; c < 10; c++ {
			for c := 1; c <= 10; c++ {
				out <- c
			}
		}
		close(out)
	}()
	return out
}

func factorail(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for c := n; c > 0; c-- {
		total *= c
	}
	return total
}