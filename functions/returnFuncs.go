package main

// It's also possible to create such functions that return functions as their return values
// This features also allow us to create  something called closures in which a function captures the values that are present in it's context where it was created

func createMultiplier(factor int) func(int) int {
	return func(i int) int { return i * factor }
}
