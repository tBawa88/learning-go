package main

import "fmt"

func main() {
	strings := []string{"a", "b", "b", "c", "d", "d"}
	strs := removeDups(strings)
	fmt.Println(strs)
}

// ["a", "b", "b", "c", "d", "d"] = ["a", "b", "c", "d"]
func removeDups(str []string) []string {
	i := 0
	for j := 1; j < len(str); j++ {
		if str[i] != str[j] {
			i++
			str[i] = str[j]
		}
	}
	return str[:i+1]
}
