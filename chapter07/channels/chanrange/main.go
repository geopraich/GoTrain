package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for count := 0; count < 10; count++ {
			c <- count
		}
		close(c)  // closes channel can no longer put values onto channel but can receive values
	}()

	for num := range c {  // drains channel (no longer needs to wait) continues until nothing left on the channel
		fmt.Println(num)
	}
}
