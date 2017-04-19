package exerciseOne

import "fmt"

// 1. fix deadlock
func main() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
	}()
	fmt.Println(<-c)
	fmt.Println(<-c)
	close(c)
}
