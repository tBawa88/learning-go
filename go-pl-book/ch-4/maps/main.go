package main

import "fmt"

// Comparing 2 maps. Maps cannot be compared with each other. The only correct comparison of a map is with nil

func main() {
	x := make(map[string]int)
	y := make(map[string]int)

	x["alice"] = 20
	x["bob"] = 22

	y["alice"] = 20
	y["bob"] = 22

	fmt.Println(isEqual(x, y))
}

func isEqual(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for key, valX := range x {
		// check if each key exists inside y, and if it does, compare their values
		if valY, ok := y[key]; !ok || valY != valX {
			return false
		}
	}
	return true
}
