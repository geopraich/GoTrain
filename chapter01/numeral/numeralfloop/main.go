package main

import "fmt"

func main() {
	for c := 0; c < 101; c++  {
		fmt.Printf("%d - %b - %#X\n", c, c, c)
	}
}
