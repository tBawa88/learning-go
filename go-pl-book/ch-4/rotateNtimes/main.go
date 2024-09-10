package main

import "fmt"

// Rotating the given slice "r" number of times using rotation method
// Let's say we're rotating to the left
// Step1: Reverse the first r elemens s[:r]
// Step2: Reverse the rest of the elements s[r:]
// Step3: Reverse the entire slice s
func main() {
	s := []int{1, 2, 3, 4, 5}
	r := 2

	reverse(s[:r])
	reverse(s[r:])
	reverse(s)
	fmt.Println(s)

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
