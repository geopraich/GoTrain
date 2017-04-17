package main

import (
	"fmt"
	"time"
)
// channels
// allows communication between go-routines

func main() {
	c := make(chan int)  // channel declaration/initialisation slices, map and channels all use make()
	// unbuffered channel - a buffer channel would be c := make(chan int, 10) would take ten values before it is
	// full

	go func() {  // anon function self executing
		for count := 0; count < 10; count++ {
			c <- count  // putting count onto channel (<-) stops here until* passes data back and forth
		}

	}()

	go func() {
		for {
			fmt.Println(<-c)  // taking number off of the channel (<-) give next value off of the channel
			// *takes the value then passing back and so on
		}
	}()

	time.Sleep(time.Second)
}