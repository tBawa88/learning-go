package main

import "fmt"

// Map is a collection of key value pairs, similar to structs, but key-value pairs in Maps are statically typed
// All the keys and All the values must be of same type

func main() {

	// colors2 := make(map[string]string) //another way of creating a map, using the make() function

	colors := map[string]string{
		"red":   "ff0000",
		"green": "00ff00",
		"blue":  "0000ff",
		"white": "fff",
		"black": "000",
	}

	colors["grey"] = "dd0022"

	fmt.Println(colors)

	//Deleting key value paris in maps is also simple
	delete(colors, "red") // first arg is the map variable, second is the key name

	printMap(colors)

}

// Iterating over maps using for loops (first variable is the key, second is it's corresponding value)
func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}

// NOTES :
// - values in maps are not stored sequentially like in a slice or an array
// - maps has constant time access
