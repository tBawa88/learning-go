package main

import "fmt"

// rotate the slice in a single pass
func main() {
	s := []int{1, 2, 3, 4, 5}
	s = rotateLeft(s)
	s = rotateLeft(s)
	s = rotateLeft(s)
	fmt.Println(s)
}

// [1,2,3,4,5] = 3 = [3,4,5,1,2]
func rotateLeft(s []int) []int {
	temp := s[0]
	for i := 1; i < len(s); i++ {
		s[i-1] = s[i]
	}
	s[len(s)-1] = temp
	return s
}
