package main

import "fmt"

func main() {
	// map/dict key:value storage declaration (type for both key:value)
	mDict := make(map[string]int)  // you can pass capacity optional

	// assigning key:value pairs
	mDict["key1"] = 42

	fmt.Println("key1: ", mDict["key1"])  // prints out value of pair of key1

	_, ok := mDict["key1"]  // checks whether map has key:value at "key1" (tuple unpack - value,key)
	fmt.Println("ok: ", ok)

	delete(mDict, "key1")  // deletes key:value pair at "key1"
	fmt.Println(mDict)

	// other ways to declare maps
	n := map[string]int{"k1": 1, "k2": 2}  // this makes the underliying data structure
	fmt.Println(n)
	var w map[string]int  // nil reference not rec as no append no elements may be added
	fmt.Println(w)

	// check to see whether value exists in map
	myGreeting := map[string]int{
		"k1":1,
		"k2":2,
		"k3":3,  // if bracket not on same line trailing comma needed
	}

	if value, ok := myGreeting["k1"]; ok{  // exists returns bool
		fmt.Println(value, ok)
	}else {
		fmt.Println("value does not exist")
	}
	fmt.Println(myGreeting)

	// range loop with maps
	for key, val := range myGreeting {
		fmt.Println(key, "-", val)
	}
}
