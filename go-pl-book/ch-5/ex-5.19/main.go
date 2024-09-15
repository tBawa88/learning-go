package main

import "fmt"

// Using panic and recover , Write a function that contains no return statement but still returns a non-zero value

func main() {
	fmt.Println(noReturn())
}

func noReturn() (val int) {
	defer func() {
		if p := recover(); p != nil {
			val = p.(int) // since the panic value is of 'any' type (p {}interface), we have to assert it's type using this syntax
		}
	}()
	panic(99)
}
