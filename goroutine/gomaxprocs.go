package main

import (
	"fmt"
	"runtime"
)

func main() {

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU:", totalCpu)

	totalThreads := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Threads:", totalThreads)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine:", totalGoroutine)

}
