package main

import "fmt"

func main() {
	fmt.Println("Hello ")
	digits := []int{1, 2, 3, 4, 5}

	fmt.Printf("Sum of index 0 and 1 %d\n", add(digits[0], digits[1]))

	fmt.Printf("Sum of entire slice %d\n", reduceSum(digits...)) // Passes all elements as individual arguments, think of spread syntax in Javascript
	// Instead of using the elipsis in front of the slice, in GO we use it after the slice

	// printGreeting("Linking park is back, im so excited ")

	// Storing functions in variables as values
	// Passing those values as arguments to another function which evetually executes them
	adder := func(a, b int) int { return a + b }
	subtractor := func(a, b int) int { return a - b }
	multi := func(a, b int) int { return a * b }
	runMathOperation(adder, "addition")
	runMathOperation(subtractor, "subtraction")
	runMathOperation(multi, "multiplication")

	times10 := createMultiplier(10)
	fmt.Printf("50 times 10 is %d\n", times10(50))

	printMessage("Linking park is back babyyy")

}

// Function with named return type
func add(a, b int) (sum int) {
	return a + b
}

// Variardic functions - All the parameters must be of type int. And all of them will be collected into a single int slice
func reduceSum(nums ...int) (sum int) {
	sum = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
