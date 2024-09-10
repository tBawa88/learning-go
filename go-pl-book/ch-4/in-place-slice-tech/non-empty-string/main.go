package main

import "fmt"

func main() {
	colors := []string{"red", "", "orange", "", "blue"}
	nColors := nonempty(colors)

	fmt.Println(nColors)
	fmt.Println(colors)

}

// nonempty removes all the empty strings from a slice of given strings. This is an example of in-place function
// in-place function modify the exisiting data structure (slice, struct, array) without using any extra space
// the returned string slice share the same memory address as the input slice.
// Data of the input slice gets overwritten in these functions
func nonempty(strs []string) []string {
	i := 0
	for _, s := range strs {
		if s != "" {
			strs[i] = s
			i++
		}
	}
	return strs[:i] // 'i' is the new length of the string slice after this loop ends, therefore only return 0 to i elements slice
}
