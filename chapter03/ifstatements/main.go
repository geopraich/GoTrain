package main

import "fmt"

func main() {
	// if condition - code body
	if true {
		fmt.Println("seo a' ruith")
	}

	if false {
		fmt.Println("chan eil a' ruith")
	}

	if !true {  // not true which is false
		fmt.Println("this didn't run")
	}

	if !false {  // not false which is true
		fmt.Println("this runs")
	}
	b := 5
	if f := 5; b == f {  // f initialised before the condition check (after; which is b)
		fmt.Println(f)
	}

	// if else statement
	if false {
		fmt.Println("will not run")
	} else {
		fmt.Println("will run")
	}

	// elif statements
	if false {
		fmt.Println("will not run")
	} else if true {
		fmt.Println("will run")
	} else if true {
		fmt.Println("will not run")
	} else {
		fmt.Println("will not run")
	}

}