package main

import "fmt"

// Any other type, in this program, that implements a function called getGreeting() string {} on it, automatically becomes a member of bot interface
type bot interface {
	getGreeting() string
}

type englishBot struct {
	greeting string
}
type spanishBot struct {
	greeting string
}

func (e englishBot) getGreeting() string { return e.greeting }
func (s spanishBot) getGreeting() string { return s.greeting }

//These 2 functions are pretty generic as in their funcitonality is identical
// func printGreeting(e englishBot)  { fmt.Println(e.getGreeting()) }
// func printGreetingS(s spanishBot) { fmt.Println(s.getGreeting()) }

func main() {
	e := englishBot{greeting: "Hello Friend"}
	s := spanishBot{greeting: "Hola amigo"}

	printGreeting(e)
	printGreeting(s)

}

// any type, that satisfies the requirement of bot interface would be able to call this function
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//NOTES :
//An interface in go defines a method signature
//Any concrete type then can choose to implement all those methods to implictly become a memeber of that interface.
//This allows us to achieve polymorphism, like we did with our printGreeting() example
