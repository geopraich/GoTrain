package exerciseTwo

import "fmt"

// 2. fix to drain channel
func main() {
	c := make(chan int)

	go func() {
		for count := 0; count < 10; count++ {
			c <- count
		}
		// close(c) goes here to use range
	}()

	for n := 0; n < 10; n++ {
		fmt.Println(<-c)
	}
}

// to use range you need to close(c) the channel before ranging