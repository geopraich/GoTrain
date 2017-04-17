package main

import "fmt"

// struct declaration - fields can also be tagged re json storage within SQL database
type person struct {
	// fields
	first string
	last string
	age int
}

type agent struct {
	person
	first string
	status string
}

// can create your own type (alias of type int(for methods etc...) bad practice)
type foo int

func main() {
	// initialisation of person type assign to p1 and p2 variables
	p1 := person{"James", "Mitchell", 20}
	p2 := person{"Miss", "Moneypenny", 34}
	fmt.Println(p1.first, p1.age, p1.last)
	fmt.Println(p2.last, p2.age, p2.first)

	// creating type agent with fields inherited from type person
	p3 := agent{person{"Roger","Whittlestone",25},"Twoface","active"}
	p4 := agent{person: person{first: "Thomas", last:"Wentworth", age:37}, first:"Breet",status:"KIA"}  // explicit
	fmt.Println(p4)
	fmt.Println(p3.person.first, p3.first)  // prints out first value from person.first (explicit) then value at
	// then first field from agent type outer type (promotion) over inner type

	// initialisation of type foo which has underlying type of int
	var age foo
	age = 44
	fmt.Printf("%T - %v\n", age, age)

	fmt.Println(p1.fullName())  // method call
	fmt.Println(p3.hello())  // agent method hello()
	fmt.Println(p3.person.hello())  // person method hello() direct notation (explicit)

	// struct pointers - can pull values straight from pointers
	p5 := &person{"Anthony","Zach", 30}
	fmt.Println(p5)
	fmt.Printf("%T\n", p5)
	fmt.Println(p5.first)
}

// methods for user defined types (struct(objects))
// creating a method for person type
// func receiver functionName(parameters) returns {functionBody}
func (p person) fullName() string {  // fullName() name of method - (p person) receiver
	return p.first +" "+ p.last
}

// overriding method example
func (p person) hello() string {
	return "Hello "+ p.first
}
// blow method overrides above method when agent type calls method hello() outer type overrides inner type
// (fields/methods)
func (a agent) hello() string {
	return "Hello "+a.first
}