package main

import (
	"fmt"
	"time"
)

func main() {

	// Create a new ticker that ticks every 1 second
	ticker := time.NewTicker(1 * time.Second)

	// Create a channel to signal when we want to stop
	done := make(chan bool)

	// Run a goroutine that receives from the ticker
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Let the ticker run for 5 seconds
	time.Sleep(5 * time.Second)

	// Stop the ticker and signal the goroutine to exit
	ticker.Stop()
	done <- true

	fmt.Println("Ticker stopped")
}
