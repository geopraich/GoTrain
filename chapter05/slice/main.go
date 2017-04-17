package main

import "fmt"

func main() {
	// slice/list declaration (can have ending comma)
	mySlice := []int{1,2,3,4,5,11}  // slice/list assign to mySlice of slice type int with 1,2,3... initialised
	var aArray [10]int  // array zero declaration
	a := [10]int{}  // array shorthand
	b := [10]int{1,2,3,3}  // array
	fmt.Println(mySlice)
	fmt.Println(mySlice[0:2])  // slice of a slice/list (not inclusive)
	fmt.Println("A string"[0:4])  // slice on a string (when accessing single index returns ASCII for char at index)
	aSlice := make([]int, 10, 20)  // make() allows definition of slice ie set length specify capacity
	fmt.Println(aSlice)  // default len() of slice initialise to zero (still dynamic) creates new underlining array
	anotherSlice := new([50]int)[0:50]  // another way of initialising slice/list
	fmt.Printf("slice: %T -slice: %T -slice: %T -array: %T -array: %T -array: %T\n",anotherSlice,aSlice,
		mySlice,aArray,a,b)
	fmt.Println(anotherSlice)
	aSlice[0] = 7  // insert shorthand - insert 7 (int) at index 0 when adding new item over
	// and above the len() of slice append() is used
	aSlice = append(aSlice, 12)
	fmt.Println(aSlice)
	mySlice = append(mySlice, aSlice...)  // appending slice to slice syntax - ... -
	fmt.Println(mySlice)
	mySlice = append(mySlice[:0],mySlice[1:]...)  // deletes from slice ... included first up to where you del
	// second where to continue from then ...
	fmt.Println(mySlice)
	mySlice[0]++  // this is the way to increment a value at index 0

	matrixSlice()
	matrixNum()
}
/* defining capacity gives slice room to grow without golang having to create a
 new underlying array every-time our slice grows (specifying only one number in make() sets to both len and cap)
*/

// function with nested slices
func matrixSlice() {
	records := make([][]string, 0)  // making a nested list/slice
	// entry 1
	st1 := make([]string, 4)
	st1[0] = "Fosters"
	st1[1] = "Beer"
	st1[2] = "100.00"
	st1[3] = "74.00"
	// store entry
	records = append(records, st1)
	// entry 2
	st2 := make([]string, 4)
	st2[0] = "Gnome"
	st2[1] = "Jolly"
	st2[2] = "99.00"
	st2[3] = "98.90"
	// store entry
	records = append(records, st2)
	// print
	fmt.Println(records)
}

// nested slice of int
func matrixNum() {
	transact := make([][]int, 5)

	for count := 0; count < 5; count++ {
		numSlice := make([]int, 0)
		for c := 10;c <= 13 ; c++  {
			numSlice = append(numSlice, c)
		}
		transact[count] = numSlice
	}
	fmt.Println(transact)
}

// make() makes the underlying data structure for slices ie array that underlies reference types
// therefore you can immediately use indexing with var and shorthand declaration you can only use append
// until all values are associated with said slice and cannot use indexing to insert values
