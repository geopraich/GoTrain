package main

import (
	"fmt"
)

func main() {
	//for initialise; condition; post - c++ also accepted
	for c := 0; c <= 100; c+=1 {
		fmt.Println(c)
	}
	nestedForLoop()
	forLoopConditional()
	forLoopNoCondition()
	anotherWay()
	loopSomeTime()
}

func nestedForLoop() {
	//outer loop runs inner runs to completion (3 times) outer runs inner runs to com... until outer runs 6 times
	for c := 0; c <= 5; c++ {
		fmt.Println(c,"- outer loop - tha sinn a-muigh an seo")
		for c := 0; c < 3 ; c++  {
			fmt.Println(c,"- inner loop - tha sinn a-staigh an seo")
		}
	}
}

func forLoopConditional() {
	c := 0  // initialise the count (c)
	for c < 10 {  // condition
		fmt.Println("still less than ten - a bheil a-nis?")
		c++  // increments count
	}
}

func forLoopNoCondition() {
	var input string
	// for doesn't need a condition will run forever (unless specified)
	for {
		fmt.Println("Write Message Here: ")
		fmt.Scan(&input)
		fmt.Println(input)
		break
	}
}

func anotherWay() {
	c := 0
	// another way to break from for loop with no condition set
	for {
		fmt.Println(c)
		if c >= 10 {
			break
		}
		c++
	}
}

func loopSomeTime() {
	c := 0
	for {
		c++
		if c%2 == 0 {
			fmt.Println("going back to top of for loop")
			continue
		}
		fmt.Println("printing odds -", c)
		if c >= 30 {
			fmt.Println("breaking for loop now")
			break

		}
	}
}