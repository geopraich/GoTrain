package main

import "fmt"

func main() {

	a := 42

	fmt.Println(a)
	fmt.Println(&a)


	var b *int = &a  // b is a pointer to the memory address where an int is stored - b is type "int pointer"
// * is part of the type - b is type *int
	fmt.Println(b)
	fmt.Println(*b)  // dereference b so it gives the value - *pointer

	*b = 9  // makes the value at the memory address 9 - in effect reassigns a variable to 9
	fmt.Println(a)  // a is now 9
	fmt.Println(&a,"\n")

	x := 5
	fmt.Println(&x)
	zero(x)
	fmt.Println(x)
	zeroWithPointer(&x)
	fmt.Println(x)  // x is now zero
}

func zero(x int) {  // without pointer
	fmt.Println(&x)  // different address different x form mains x
	  x = 0
}

func zeroWithPointer(x *int) {  // takes type pointer (address of a var) int - *int is own type
	*x = 0  // dereferences to effect changes to x at same address as mains x
}
