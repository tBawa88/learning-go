package main

/*
	One way of implementing mutual exclusion is using a binary semaphore.
	Basically a buffered channel with the capacity of 1
	Before accessing the shared resource the goroutine shall acquire the token and release it after it's done
*/

var (
	sema    = make(chan struct{}, 1) //guards balance
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{}
	balance += amount
	<-sema
}

func Balance() int { return balance }
