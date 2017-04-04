package main

import (
	"fmt"
	"github.com/geopraich/GoTrain/chapter02/scope/closure/anonfunc"
)

func main() {
	x := 42
	fmt.Println(x)  // outer scope
	{
		fmt.Println(x)
		y := "answer to the universe and everything"
		fmt.Println(y)  // inner scope
	}  // inner scope ends here
	anonfunc.Anon()
	increment := wrapper()  // func expression
	fmt.Println(increment())
	fmt.Println(increment())
}

// func wrapper() returns function (that returns int)
func wrapper() func() int {
	x := 0
	return func() int {  // this is the function that is returned - which in turn returns x (int)
		x++
		return x
	}
}
/*scope blocks:
universe
package
file
function
 */