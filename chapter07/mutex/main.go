package main

import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex  // mutex one way to prevent race conditions

func main() {
	wg.Add(2)  // adds two to wait group
	go incrementor("Foo:")
	go incrementor("Boo:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for count := 0; count < 50; count++ {
		mutex.Lock()  // any threads waits as it lock to a single thread all others have to wait
		x := counter
		x++
		counter = x
		fmt.Println(s, count, "counter:", counter)
		mutex.Unlock()  // unlocks threads
	}
	wg.Done()
}

// go run -race main.go - determines whether you have a race condition
// vs
// go run main.go