package main

import (
	"fmt"
	"strings"
)

type contactInfo struct {
	houseNo int
	street  string
	phone   int
}

type person struct {
	firstName string
	lastName  string
	age       int
}

// Embedding a struct inside another struct to create more complex data structures
type personWithContact struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//First way (best way of declaring a struct, specifying which value goes into which field)
	p := person{
		firstName: "Tarun",
		lastName:  "Bawa",
		age:       26,
	}

	//Second way of defining a struct (relying on the order in which the fields were defined)
	alex := person{"alex", "jason", 23}
	alex.printFullName()

	alex.firstName = "updated first name"
	alex.printFullName()

	// Declaring struct with "Zero values", then populating each field individually
	var testPerson person
	testPerson.firstName = "test"
	// testPerson.lastName = "person"
	testPerson.age = 22
	fmt.Printf("%+v \n", testPerson) //printing the struct along with it's fieldnames to visualize the zero values

	// p.greet()
	// p.upperCase()
	p.printFullName()

	var tbawa personWithContact
	tbawa.firstName = "tejas"
	tbawa.lastName = "bawa"
	tbawa.contact.houseNo = 420
	tbawa.contact.street = "Arlong park street"
	tbawa.contact.phone = 6969696969

	tbawa.printAddress()

	tabahi := personWithContact{
		firstName: "Divyam",
		contact: contactInfo{
			houseNo: 340,
			street:  "Raftael",
			phone:   12121212,
		},
	}

	fmt.Printf("%+v", tabahi)

	tabahi.printAddress()
}

func (p personWithContact) printAddress() {
	fmt.Printf("Street %s, House %d , Phone %d \n", p.contact.street, p.contact.houseNo, p.contact.phone)
}

func (p *person) upperCase() {
	p.firstName = strings.ToUpper(p.firstName)
}

func (p person) printFullName() {
	fmt.Println(p.firstName + " " + p.lastName)
}

func (p person) greet() {
	fmt.Println("Welcome home ", p.firstName)
}
