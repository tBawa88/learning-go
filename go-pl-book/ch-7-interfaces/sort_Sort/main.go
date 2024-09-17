package main

import (
	"fmt"
	"sort"
)

// To sort a custom type using the sort.Sort() method, that type must implement 3 methods Len() Less() Swap()

type StringSlice []string // Just a wrapper over a []string slice

func (s StringSlice) Len() int           { return len(s) }
func (s StringSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s StringSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	names := []string{"Tarun", "Teja", "Suvneet", "Sukhmani", "Jason", "Alice", "Bob"}

	sort.Sort(StringSlice(names))

	for _, name := range names {
		fmt.Printf("%s\n", name)
	}
}

// We can sort list of complex types using this sort.Sort() method as long as we provide our custom sorting logic in Less() and Swap()  and Len() method
