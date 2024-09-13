package main

// The string literal syntax after Year and Color is called a field tag
// The field tag is used to overwrite the name of the fiels when this struct is Marshalled into JSON
//`json:"newName",option` => The options is optional. We can decide whether we want to include a particular field in the JSON string or not,
// or omit the field if it has a zero value for it's type. (false, "", 0, empty struct map interface)
import (
	"encoding/json"
	"fmt"
	"log"
	"tbawa/go-course/go-pl-book/ch-4/JSON/placeholder"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

// Suppose we needed a list of Movies. Data structures like this are perfect use case for JSON
var movies = []Movie{
	{Title: "3 Idiots", Year: 2014, Color: false, Actors: []string{"Amir khan", "R Madhavan", "Boman Irani"}},
	{Title: "3 Idiots", Year: 2014, Color: false, Actors: []string{"Amir khan", "R Madhavan", "Boman Irani"}},
	{Title: "3 Idiots", Year: 2014, Color: false, Actors: []string{"Amir khan", "R Madhavan", "Boman Irani"}},
}

func main() {
	//Converting a Go data strucutre like []Movies to JSON is called "Marshalling"
	data, err := json.Marshal(movies)
	data2, err := json.MarshalIndent(movies, "", "  ")

	if err != nil {
		log.Fatal("Error encoding movies to JSON", err)
	}

	fmt.Printf("%s\n", string(data))

	fmt.Println("Printing JSON data in more readable form")
	fmt.Printf("%s\n", string(data2))

	// The inverse of Marshal is UnMarshal.
	// By defining suitable GO data structures, we can filter out which part of the JSON data we want to unmarshal
	var titles []struct {
		Title    string
		released int
	}
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatal("Error Unmarshaling the JSON data")
	}
	fmt.Println(titles)

	fmt.Println("Printing user data")
	placeholder.GetUserData()
}

// Don't forget to pass in a pointer to the type as the 2nd argument to json.Unmashal(). Passing by reference will make sure that the function modifies our original variable
// var titles [] struct{Title string} => this is also how you can decalre custom types var title struct {Title string}
