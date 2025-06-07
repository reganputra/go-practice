package main

import (
	"fmt"
	"sync"
	"time"
)

// Mutex is to deal with race conditions in goroutines
func main() {
	x := 0
	// Using a mutex to protect the shared variable x
	var mu sync.Mutex
	for i := 1; i <= 100; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				mu.Lock() // Lock the mutex before accessing the shared variable
				x += 1
				mu.Unlock() // Unlock the mutex after accessing the shared variable
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Count:", x)

}
