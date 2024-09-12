package main

import "fmt"

var m = make(map[string]int)

// Sprintf(format, a ...any) takes in variardic input and formats them into a single string and returns the resultant string
// %q is for double quoting the string
func k(list []string) string { return fmt.Sprintf("%q", list) }

func Add(list []string) { m[k(list)]++ }

func Count(list []string) {
	fmt.Printf("For List %q, Add() has been called %d times\n", list, m[k(list)])
}

func main() {
	m := []string{"bob"}
	m = append(m, "alice")
	Add(m)
	Add(m)
	Add(m)
	Add(m)
	Add(m)

	Add(append(m, "jason"))
	Add(append(m, "jason"))
	Count(append(m, "jason"))
}

// We cannot use a slice as a key of a map since we cannot compare slices
// One workaround is to create a helper function that converts a slice into a single string. Now that string can be used as a map key
// Then we create a map[string]

// In this program, the Count(list []string) function counts that for a given slice of string, how many times has the Add() function been called
