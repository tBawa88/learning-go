package main

import "fmt"

type contactInfo struct {
	email string
	pin   int
}

// dont seperate by any commas or any colons like in TS
type person struct {
	firstname string
	lastname  string
	age       int
	// contact   contactInfo
	contactInfo //we can leave the field name out, but then we must use the same name as the struct while declaring an instance of this struct
}

type voterList []person //creating more complex types using struct types

func main() {
	//Here we're implicitly assigning values, and go will automatically assign them in ther order thery're defined inside the actual struct. Like "jason" will go to firstname and so on
	//but what if we swapped the position of the fields inside the struct definition? Then it would mess up our representation
	// p := person{"Jason", "kang", 33}

	// //Explicitly assingin values to the fields using property names, just like we do it in JS
	// tarun := person{firstname: "Tarun", age: 26, lastname: "Bawa"}

	// //declaring a struct like this, go compiler will initialize all the properties with zero values, just like it does with any variable type
	// var per person

	// fmt.Println(tarun) //{Tarun Bawa 26} the printing will happen in the same order tho
	// fmt.Println(p.firstname)
	// fmt.Printf("%+v", per) //%+v prints the structs along with it's property names {firstname: lastname: age:0}

	// per.firstname = "Alex"
	// per.lastname = "Prince"
	// per.age = 69
	// fmt.Printf("\n%+v", per)

	//To define a struct which is embedded inside a struct, we use the same syntax that we use to define the outer struct p := structType {a, b, c}
	//NOTE: another thing to keep in mind is that when defining properties of structs inside multiple lines, each property must end with a comma "," even if it's the last property
	jim := person{
		firstname: "JIM",
		lastname:  "Purdy",
		age:       30,
		contactInfo: contactInfo{
			email: "aaa@gmail.com",
			pin:   121212,
		},
	}
	// fmt.Printf("%+v", jim)
	fmt.Println(jim.greet())
	jim.updateLastname("Keller")
	fmt.Println(jim.greet())

	//pointers
	jp := &jim
	jp.UpdateLastname("Keller")

	jim.updateLastname("lastname") //has the same effect as calling it with the pointer go will automatically pass in the address of this struct
	//Go automatically handles the method calls in a way that makes sure that receiver type is respected

	fmt.Println(jp.greet())
	fmt.Println(jim.greet())

	var x person
	fmt.Printf("\n%+v\n", x)

	// studentList := []person{jim, *jp, jim }	//you can also do this, create an array of person type

}

func (p person) greet() string {
	return "Good morning " + p.firstname + " " + p.lastname
}

//this is the non pointer one
func (p person) updateLastname(newLastname string) {
	p.lastname = newLastname
}

//this is the pointer one
//NOTE: even if we dont use the defrencing sytax, Go will automatically derefrence it for us. This is syntantical sugar Go uses to improve the code when working with pointers
func (pointerToPerson *person) UpdateLastname(newLastname string) {
	(*pointerToPerson).lastname = newLastname
}
