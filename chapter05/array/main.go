package main

import "fmt"

func main() {
	// array declaration
	var aExample [58]string  // numbers in bracket (length definition part of type)
	// defines array rather than slice/list - array of type string length 58

	fmt.Println(aExample)
	fmt.Println(len(aExample))
	fmt.Println(aExample[12])  // print element at index 12 (13th element)

	for count := 65; count <= 122; count++ {
		aExample[count-65] = string(count)  // string() casts int to ASCII char inserts at [count-65] index
	}

	fmt.Println(aExample)
	fmt.Println(len(aExample))
	fmt.Println(aExample[50])

}
