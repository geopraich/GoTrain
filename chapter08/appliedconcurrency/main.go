package main

import "fmt"

func main() {
	c1 := incrementor("Foo:")
	c2 := incrementor("Boo:")
	c3 := puller(c1)
	c4 := puller(c2)
	fmt.Println("Final Counter:", <-c3+<-c4)  // adds c3 to c4 all on int on channel stops main exiting
}

func incrementor(s string) chan int {
	out := make(chan int)
	go func() {
		for count := 0; count < 20; count++ {
			out <-1
			fmt.Println(s, count)
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for num := range c {
			sum += num
		}
		out <- sum
		close(out)
	}()
	return out
}
