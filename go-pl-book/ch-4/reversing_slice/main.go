package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(nums)

	// reverse(nums)
	// fmt.Println(nums)

	//Since nums[n] is a pointer to a specific index in the slice
	// We can pass in a range of a slice and reverse that specific range without reversing the entire slice
	// reverse(nums[0:3])
	// fmt.Println(nums)

	// Rotating a slice 3 times to the left [1,2,3,4,5,6] => [4,5,6,1,2,3]
	// 1. Reverse the first 3 elements [3,2,1,4,5,6]
	reverse(nums[:3])
	// 2. Reverse the remaining elements [3,2,1,6,5,4]
	reverse(nums[3:])
	// 3. Reverse the entire slice [4,5,6,1,2,3]
	reverse(nums[:])

	fmt.Println("Slice rotated 3 times to the left")
	fmt.Println(nums)

}

// slices are pass by reference, so any change made to the slice by this func will mutate the original slice
func reverse(s []int) {
	//start 2 pointers one from index 0, another from the last index
	//keep swapping the elements untill the indices meet in the middle
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
