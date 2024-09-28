package main

import "sync"

/*
	The idea is same as binary semaphore. Before accessing the shared resource, goroutine must acquire the lock
	If the lock is currently acquired by another goroutine, then current goroutine gets blocked untill the lock becomes available again (other goroutine calls Unlock())
*/

var (
	mu       sync.Mutex
	balance2 int
)

func Deposit2(amount int) {
	mu.Lock()
	balance2 += amount
	mu.Unlock()
}
