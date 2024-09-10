package main

import "fmt"

// Slices declared with var s []int , are not only empty slices(len == 0) but they're also equal to nil s == nil // true
// Slices created using s := []int{} are empty but not nil
// Same with slices created using make() function s:= make([]int, 0), it's length is 0, but it's not nil

// Slices have length and a capacity which determines how much a slice can grow before it needs to be allocated new memory

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	// Both slices are starting out as empty slices
	// we're appending x and storing it in y, you can see the capacity being doubled every time the new length exceeds the old one
}

func appendInt(s []int, val int) []int {
	var z []int
	zLen := len(s) + 1 // new length will oldLen + 1

	//check if the new length exceeds the current capacity or not
	if zLen <= cap(s) {
		// this is not copying, we're extending z to the current slice s (z will have same length "and capacity as slice s")
		// also this will not produce array out of bounds error since zlen < cap(s)
		z = s[:zLen]
	} else {

		// This is the minimum new capacity to accomodate the old elements and the new one
		// The reason for doubling mem alloc if the is so that we don't have to copy on every single interation (which is O(n)) and keep using the same underlying slice
		//as much as we can on future append calls
		zcap := zLen
		if zcap < 2*len(s) {
			zcap = 2 * len(s)
		}
		z = make([]int, zLen, zcap) // 2nd is the length of the new slice, and 3rd arg is the new capacity of this new slice
		copy(z, s)                  // copy(target, source)	, special function made to copy slices
	}

	z[len(s)] = val
	return z
}
