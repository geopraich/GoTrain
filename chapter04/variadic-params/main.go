package main

import (
	"fmt"
)

func main() {
	n := average(43, 56, 87, 12, 45, 47)
	fmt.Println(n)
	data := []float64{43, 56,87, 21, 45, 57}  // assigning a list of type float64 to data
	d := average(data...)  // passing a list of values as a variadic argument ie multi args one at a time via ...
	fmt.Println(d)
	fmt.Println(avg(data))
}

// for [index], element := range iterable (for index/count, element in enumerate(iterable):
func average(nameParam ...float64) float64 {  // parameter is variadic ...syntax means it can take from 0+ arguments
	total := 0.0
	for _, v := range nameParam {  // range iterates through an iterable ie a list etc... (for .. in.. python)
		total += v
	}
	return total/float64(len(nameParam))
}

// without variadic params and args
func avg(nameParam []float64) float64 {
	total := 0.0
	for _, v := range nameParam {
		total += v
	}
	return total/float64(len(nameParam))
}
