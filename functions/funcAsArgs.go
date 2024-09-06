package main

import "fmt"

var a int = 10
var b int = 20

func runMathOperation(fn func(a, b int) int, opName string) {
	fmt.Printf("Result of %s is %d \n", opName, fn(a, b))
}
