package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	jim := person{
		firstName: "jimmy",
		age:       12,
	}
	jim.print()

	jimPointer := &jim //storing the address of 'jim' in a pointer variable
	doublePointer := &jimPointer
	jimPointer.updateName("Jaime")

	jim.print()
	(*doublePointer).updateName("Double jimmy")
	jim.print()

	triplePointer := &doublePointer
	(**triplePointer).updateName("Tripple jimmy")

	jim.print()

	// Just like C, it's possible to create double and triple and so on ... , pointers
	// Pointers store memory address of the value instead of storing the value itself
	// To extract the value stored in that memory address, we have to "Derefrence" the pointer using (*ptr).updateName() syntax

	// '*' in front of a type means we're dealing with a pointer of that type
	// '*' in front of a variable means that it's some pointer and we're derefrencing it

	printPointer(&jim)
	// printPointer(jim)   ---> THIS IS WRONG
	// Something to keep in mind, the shortcut of calling a receiver function that accepts a pointer of a type using a simple variable of that type works only for receiver functions
	// BUT if a function is expecting an argument of a pointer type, then that argument cannot be a simple variable, it must be an address of that type to satisfy the pointer type

}

func (p person) print() {
	fmt.Println(p.firstName)
}

func printPointer(p *person) {
	fmt.Println((*p).firstName)
}

func (p *person) updateName(name string) {
	(*p).firstName = name
}
