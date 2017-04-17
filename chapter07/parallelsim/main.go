package main

import (
	"fmt"
	"sync"
	"runtime"
)

var wg sync.WaitGroup  // enables the wait group

func init() {  // this enables parallelism  (init) special func name to initialise set for program)
	runtime.GOMAXPROCS(runtime.NumCPU())  // uses all cpu's cores (this is set by (after 1.7) default)
}

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
	wg.Done()  // this takes another away from the wait group
}