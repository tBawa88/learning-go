package main

import "fmt"

type englishBot struct{}
type hindiBot struct{}

type Bot interface {
	getGreeting() string
}

func (e englishBot) getGreeting() string {
	return "Hello there"
}

func (h hindiBot) getGreeting() string {
	return "Namaste"
}

type student struct {
	name string
}

type studentList []student

//Instead of creating 2 seperate functions for printing, we created an interface that specifies a contract of a getGreeting() string function
//Any type that has this receiver function, will implicitly inherit this interface

func printBotGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(e englishBot) {
// 	fmt.Println(e.getGreeting())
// }

// func printGreetingH(h hindiBot) {
// 	fmt.Println(h.getGreeting())
// }

func main() {
	e := englishBot{}
	h := hindiBot{}

	// printGreeting(e)
	// printGreetingH(h)

	// Since both englishBot and hindiBot have a receiver method called getGreeting() string, they satisfy the Bot interface and hence can be passed into printBotGreeting(b Bot)
	printBotGreeting(e)
	printBotGreeting(h)

	//Testing out type aliases of around structs
	classList := studentList{{name: "Tarun"}, {name: "Tejaswi"}, {name: "Suvneet"}, {name: "Sukhmani"}}
	classList = append(classList, student{"Testing"})
	fmt.Printf("%+v \n", classList)

}
