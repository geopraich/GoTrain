package anonfunc

import "fmt"

// anonymous function contained within Anon() non named func set to var increment - func expression
// is the way to write nested func
func Anon() {
	x := 0
	increment := func() int {  // int return type
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
