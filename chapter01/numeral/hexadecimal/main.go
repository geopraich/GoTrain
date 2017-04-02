package main

import "fmt"

func main() {
	fmt.Printf("%d - %b - %X\n", 42, 42, 42)  // print format args to decimal, binary, hex (with A-F)
	fmt.Printf("%#X", 42)  // 0X signifier added
}
