package main

import "fmt"

func main() {
	//Defining a map where key is of type string, and value is of type string
	//just like in struct, each entry must be sepereated with a comma, even the last one
	/**colors := map[string]string{
		"red":   "#f00",
		"green": "#0f0",
		"blue":  "#00f",
		"black": "#000",
		"white": "#fff",
	}**/

	//2nd way of declaring a map. This approach is used when we want to dynamically add or remove key value pairs
	// var colors map[string]string //initializes it with a zero value (an empty map)

	//3rd way : using the "make()" function
	colors := make(map[string]string) //as we know using short variable declaration, we must use this value now or it wont compile
	colors["white"] = "#fff"          //unlike sturcts , we cannot use "." to access the keys or values. We must always use the [] brackets and provide the correct type
	colors["black"] = "#000"

	//deleting keys and values from an exisiting map
	delete(colors, "white") // delete is a specific function built in for deleting elements from a map, using their keys

	fmt.Println(colors) // map[black:#000 blue:#00f green:#0f0 red:#f00 white:#fff]

	bikes := map[string]string{
		"re":     "interceptor",
		"bajaj":  "dominar",
		"triump": "speed400",
		"honda":  "cb350",
	}

	printMap(bikes)

}

//A function that iterates over map passed as an argument
func printMap(m map[string]string) {
	m["re"] = "best bikes"
	//it's done using the same range for loop that we used to iterate over a slice
	for key, value := range m {
		fmt.Println(key, value)
	}

}
