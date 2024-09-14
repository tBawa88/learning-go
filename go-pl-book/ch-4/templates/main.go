package main

import (
	"fmt"
	"os"
	"text/template"
)

type CustomWriter struct{}

var count int = 0

func (c CustomWriter) Write(p []byte) (int, error) {
	count++
	fmt.Println(string(p), p, "Count = ", count)
	return len(p), nil
}

func main() {
	//define the template string
	tmpl := `Hello{{.Name}}!, Welcome to {{.Place}}`

	// create a template object using and parse the string to it
	t := template.Must(template.New("welcome").Parse(tmpl))

	// Create a data structure to pass to the template
	// Make sure that all fields of the data structure are exported or else they won't be accessible to the tempalate
	data := struct {
		Name  string
		Place string
	}{"Tejas", "New York"}

	c := CustomWriter{}
	// Execute the template, it takes in a writer and the data to be passed to the template. It writes the resultant string to the writer
	t.Execute(c, data)
	t.Execute(os.Stdout, data)

	// The template engine writes the string output in chunks,Using os.Stdout might seems like it printed it out in a single string
	// but it's not that. Stdout is desgined to handle chunked input

}
