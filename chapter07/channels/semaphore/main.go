package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for count := 0; count < 10; count++ {
			c <- count
		}
		done <- true  // acting as a flag
	}()

	go func() {
		for count := 0; count < 10; count++ {
			c <- count
		}
		done <- true
	}()

	go func() {
		<-done  // waits until value comes through throws value away off the channel
		<-done
		close(c)
	}()

	for num := range c {  // drains channel when received takes off the channel
		fmt.Println(num)
	}
}
