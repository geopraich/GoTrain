package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup  // enables the wait group

func main() {
	wg.Add(2)  // add to the wait group
	go foo()  // go routines
	go bar()
	wg.Wait()  // waits here until wait group is zero
}

func foo() {
	for count := 0; count < 45; count++{
		fmt.Println("Foo:",count)
	}
	wg.Done()  // this takes one away from the wait group
}

func bar() {
	for count := 0 ; count < 45; count++{
		fmt.Println("Bar:",count)
	}
	wg.Done()  // this takes another away from teh wait group
}

// concurrency is not always parallelism