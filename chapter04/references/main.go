package main

import "fmt"

func main() {
	m := make([]string, 1, 25)  // sets a list/slice of length 1 capacity of 25 to m
	fmt.Println(m)
	changeMe(m)
	fmt.Println(m)
	k := make(map[string]int)  // key:value pair keyType:string elementType:int type:dict-map[string]int
	changeMe2(k)
	fmt.Println(k["Todd"])
	fmt.Println(k)
}

func changeMe(lst []string) {
	lst[0] = "Wibble"  // inserts "Wibble" at index 0 in lst []string
	fmt.Println(lst)
}

func changeMe2(m map[string]int) {
	m["Todd"] = 44  // sets keyType:"Todd and elementType:44
}
