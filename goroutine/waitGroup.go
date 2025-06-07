package main

import (
	"fmt"
	"sync"
	"time"
)

func runGroup(id int, group *sync.WaitGroup) {
	defer group.Done()

	fmt.Printf("Goroutine %d: Mulai...\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Goroutine %d: Selesai.\n", id)
}

func main() {
	group := &sync.WaitGroup{}

	numGoroutine := 100
	fmt.Printf("Meluncurkan %d goroutine...\n", numGoroutine)

	for i := 0; i < numGoroutine; i++ {
		go runGroup(i, group)
	}

	fmt.Println("Menunggu semua goroutine selesai...")
	group.Wait()
	fmt.Println("Done! Semua goroutine telah selesai.")

}
