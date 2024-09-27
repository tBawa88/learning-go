package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Add a function Withdraw(amount int) bool to the gopl.io/ch9/bank1
	program. The result should indicate whether the transaction succeeded or failed due to insuf-
	ficient funds. The message sent to the monitor goroutine must contain both the amount to
	withdraw and a new channel over which the monitor goroutine can send the boolean result
	back to Withdraw.
*/

var deposits = make(chan int)
var balances = make(chan int)

type withdraw struct {
	withdrawAmount  int
	withdrawChannel *chan bool
}

var withdrawals = make(chan withdraw)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func WithdrawAmount(amount int) bool {
	ch := make(chan bool)
	w := withdraw{amount, &ch}
	withdrawals <- w

	success := <-ch
	fmt.Println("withdrawal status ", success)
	return success
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case withdraw := <-withdrawals:
			if balance < withdraw.withdrawAmount {
				*withdraw.withdrawChannel <- false
			} else {
				balance -= withdraw.withdrawAmount
				*withdraw.withdrawChannel <- true
			}
		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	go teller()

	// deposit go routine
	wg.Add(1)
	go func() {
		defer wg.Done()
		Deposit(1000)
	}()

	// check balance goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 1)
		fmt.Println("Current balance ", Balance())
	}()

	// withdrawl goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		WithdrawAmount(500)
	}()

	// check balance goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 1)
		fmt.Println("Current balance ", Balance())
	}()
	wg.Wait()
}
