package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(remove(s, 0))
}

// Removes an element from the middle of the slice while maintainting the order of the slice
func remove(slice []int, p int) []int {
	copy(slice[p:], slice[p+1:])
	return slice[:len(slice)-1] // the last element is now a duplicate, so exclude that from the output slice
}
