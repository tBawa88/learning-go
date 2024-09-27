package main

/*
	Data races occur when : Two or more goroutines, concurrently, access a 'shared' resource, and atleast one of those accesses is a write operations
	So 3 main reasons why data races happen
		1. Write operation is happening to shared variable inside goroutines
		2. A variable is being shared inside 2 or more goroutines
		3. Concurrent access (2 or more goroutines having the ability to mutate it at the same time)
*/

/*
	3 main solutions
		1. Don't write the variable in a goroutine at all. Do all the writing outside the scope of a goroutine, preferrably at the package level outside of main
		2. Don't let more than 1 goroutine have access to the shared variable. Confine all shared variables to a single goroutine.
			Let read/writes happen through a channel (do not communicate by sharing memory, share memory by communicating)
		3. Don't let multiple goroutines mutate the shared variables all at once. **Implement mutual exclusion**
*/

// Following solution 2 in the bank example, using channles to controll concurrency
var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits: // constantly listening for any value coming out of this channel
			balance = balance + amount
		case balances <- balance: // pushing latest balance to this channel untill a receive is called on this channel
		}
	}
}

func main() {
	go teller()
}
