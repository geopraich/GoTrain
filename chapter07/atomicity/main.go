package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int64  // int64 for atomic

func main() {
	wg.Add(2)
	go incrementor("foo:")
	go incrementor("boo:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for count := 0; count < 50; count++ {
		atomic.AddInt64(&counter, 1)  // adds 1 to counter prevents race condition
		fmt.Println(s, count, "Counter:", counter)
	}
	wg.Done()
}