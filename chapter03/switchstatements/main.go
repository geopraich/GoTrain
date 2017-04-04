package main

import "fmt"

func main() {
	// switch statements runs through cases until evaluates to TRUE or hits default
	switch "Wibble" {  // check condition (no default fallthrough break not needed)
	case "Dibble":
		fmt.Println("Dibble You!")
	case "Wibble":  // will run this as evaluates to TRUE breaks at this point
		fmt.Println("Wibble You!")
	case "Mibble":
		fmt.Println("Mibble You!")
	default:  // would run if no case runs
		fmt.Println("No 'ibble For You!")
	}
	exampleFallTro()
	exampleUr()
	exampleUrTwo()
	exampleOnType("s")
}

func exampleFallTro() {
	// a switch with fallthrough will keep evaluating cases
	switch "aodann" {
	case "ceann":
		fmt.Println("head")
	case "aodann":  // this satisfies check runs
		fmt.Println("face")
		fallthrough  // this makes the next case run no check
	case "sùilean"	:
		fmt.Println("eyes")  // switch breaks at this point if more fallthrough would run next etc...
	default:
		fmt.Println("carson an t-aodann agad?")
	}
}

func exampleUr() {
	// two checks
	switch "Jobby" {
	case "Gaiket", "Jobby":  // runs as one of the checks satisfies the condition (|| in effect)
		fmt.Println("Gaiket jobby heid!")
	case "Scunt", "Cacan":
		fmt.Println("Scuntit cacan!")
	default:
		fmt.Println("dinnae kin fit yer oan aboot!")
	}
}

func exampleUrTwo() {

	myWord := "sgoilt"

	switch {  // switch with no check
	case len(myWord) == 4:
		fmt.Println("teuchter")
	case myWord == "roch", myWord == "chav":
		fmt.Println("rochal")
	default:
		fmt.Println("nothing matched")
	}
}

type Contact struct {
	greeting string
	name string
}

func exampleOnType(x interface{}) {  // argument can be of any type
	// interface{} allows any type to be stored (so x can be any type)
	// can switch on type -normally on value
	switch x.(type) { // this is asserting - x is of this "type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}
