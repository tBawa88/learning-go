package main

import "fmt"

/*
	Another classic example is the bank balance deposit and retreival example
*/

var balance int

func Deposit(amount int) { balance = balance + amount }
func Balance() int       { return balance }

func main() {
	// update for alice
	go func() {
		Deposit(200)
		fmt.Println(Balance())
	}()

	// update for bob
	go func() {
		Deposit(100)
	}()
}

/*
	Both of these goroutines will be fired at the same time
	Imagine while updating the balance for alice, the read happends (balance = 0), and just before writing, bob's query has finished updating the balance to 100
	After this alice's query will update the balance to 200. If that happends, bob's update of 100 get's lost
*/
