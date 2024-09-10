package main

import "fmt"

// The built in appen() function let's us append more than one elements at a time
// append(slice , ...int) // we can pass entire slice by spreading it using the elipisis operator
// we can modify our appendInt() function to match this behaviour

func main() {
	a := []int{1, 2, 3}

	fmt.Println(a)

	a = appendInt(a, 4)
	fmt.Println(a)

	a = appendInt(a, 5, 6, 7, 8)
	fmt.Println(a)

}

// Using variadic function, we can accept a slice instead of accepting a single element
func appendInt(s []int, y ...int) []int {
	var z []int
	zlen := len(s) + len(y)

	if zlen <= cap(s) {
		z = s[:zlen] //in this case, simply make z point to s
	} else {
		zcap := zlen
		if zcap < 2*len(s) {
			zcap = 2 * len(s)
		}

		z = make([]int, zlen, zcap)
		copy(z[:len(s)], s) //First extend z to where s is currently
	}

	//finally copy all the elements of y into z
	copy(z[len(s):], y)

	return z
}

// We're checking if the combined len of both slices is less than the capacity of slice s, then just simply point z to s[:zlen]
// else , check if the current capacity (which is combined length of len(s) + len(y)) is less than twice the len(s)
// in that case make new capacity twice the combined length
// If incase len(y) is more than 2 * len(s) then in that case the capacity will anyways be greater and able to accomodate all elements of y
// after allocating new memory, copy over all elements of slice s to z (the first len(s) of z should be equal to s)
// then copy the slice y into z (starting from the index len(s))
