package main

import "fmt"

func main() {
	//total := 1
	var f int64
	fmt.Print("Enter Number:")
	fmt.Scan(&f)

	//c := make(chan int)

	//go func() {
	//	for count := f; count > 0; count--{
	//		c <- count
	//	}
	//	close(c)
	//}()

	//for n := range c {
	//	total *= n
	//}
	fmt.Println("Total:", factorial(f))

}

func factorial(number int64) int64 {
	var total int64 = 1
	c := make(chan int64)

	go func() {
		for count := number; count > 0; count--{
			c <- count
		}
		close(c)
	}()

	for n := range c {
		total *= n
	}
	return total
}
