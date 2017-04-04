package main

import (
	"fmt"
)
// rune is unicode character || just a char
func main() {
	fmt.Println([]byte("Hello"))  // casts string into bytes appends each to a list
	for c := 50; c <= 140; c++ {
		fmt.Println(c,"-",string(c),"-",[]byte(string(c)))  // takes int from c casts to string ie unicode char
		// then casts to byte equivalent form appends to list
		fmt.Println('c')  // single quotes represents a rune not a string in go therefore
		// this will print out the unicode/ASCII decimal for lower c which is 99
		// also note '9' only works with one rune not '99' as rune is one character [57 57] is 99
	}
	fmt.Println(string(69))  // casts to ASCII character at 69 so E rather than "69" string
	b := 'a'  // sets rune to b - int32 - 97
	fmt.Println(string(b))  // this will print character a as it casts from 97 ASCII code to rune ie character
	fmt.Println(rune(b))  // this gives the code for rune a which is unicode/ASCII 97
	// `this is a raw string literal` - uninterpreted (UTF-8-encoded) characters between quotes
	// can have new lines no escape character (maintains spacing etc...)
	// "this is a interpreted string literal" character sequence
	// 'r' this is a rune literal
}