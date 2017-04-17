package main

import (
	"fmt"

	"strconv"
)

func main() {
	var x = 12
	var y = 12.12221
	fmt.Println(y + float64(x))  // conversion to float64 (widening conversion)
	fmt.Println(int(y) + x)  // conversion from float64 to int (narrowing conversion)

	var r rune = 'a'  // rune is an alias for int32
	var f int32 = 'a'
	fmt.Println(r)
	fmt.Println(f)
	fmt.Println(string(f))  // converts form in32 to string
	fmt.Println(string(r))
	fmt.Println([]byte("hello"))  // string to slice of bytes (parenthesis for conversion)
	g := []int32("hello")
	fmt.Println(string(g))

	// ASCII to INT
	var z = "12"
	var u = 62
	k, _ := strconv.Atoi(z)
	fmt.Println(k + u)
	var j = 34
	i := strconv.Itoa(j)  // converts int to ascii
	fmt.Println(i + " Wibbles")
}