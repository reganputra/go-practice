package main

import (
	"fmt"
	"sync"
	"time"
)

type bankAccount struct {
	rwMutes sync.RWMutex
	balance int
}

func (ba *bankAccount) Deposit(amount int) {
	ba.rwMutes.Lock() // Lock for writing
	ba.balance += amount
	ba.rwMutes.Unlock() // Unlock after writing
}

func (ba *bankAccount) GetBalance() int {
	ba.rwMutes.RLock() // Lock for reading
	ba.rwMutes.RUnlock()
	return ba.balance

}

func main() {
	// Example usage of bankAccount
	acc := bankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				acc.Deposit(1) // Deposit 1 unit
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Final balance:", acc.GetBalance())

}
