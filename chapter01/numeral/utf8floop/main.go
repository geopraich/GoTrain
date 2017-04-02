package main

import "fmt"

func main() {
	for c := 0;c <= 200; c+=1 {
		fmt.Printf("%d - %b - %X - %q\n", c, c, c, c)  // adds utf-8 to sequence (%q)

	}
}
