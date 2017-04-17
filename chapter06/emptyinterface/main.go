package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

// empty interface takes all types
func specs(a interface{}) {
	fmt.Println(a)
}


func main() {
	fido := dog{animal{"bark"}, false}
	fifi := cat{animal{"hiss"},false}
	specs(fifi)
	specs(fido)
}
