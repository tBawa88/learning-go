package main

// They way to prevent a function panic to crash the program is using the built-in recover() function

// A defer call is placed inside a function that could potentially panic in the future
// recover() function is called inside that defer call
// If in future, this function panics, recover() will end the state of panic, and return the panic value
// the panicked function doesn't resume the execution, instead it's forced to return normally
// recover() call only works if the function panics, otherwise it has no effect and returns nil

// The reason why recover() is placed inside a defer call : When a function panics, all pending defer calls are immidiately executed, which in turns call recover() function that ends panic
import "fmt"

func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
			return
		}
	}()
	f(x - 1)
}

// In this case we won't get a stack trace, but the program will execute normally, since we used recover() inside the defer call
