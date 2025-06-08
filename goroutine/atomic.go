package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var ops atomic.Int64

	var wg sync.WaitGroup

	for range 1000 {
		wg.Add(1)
		go func() {
			for range 100 {
				ops.Add(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops:", ops.Load())

}
