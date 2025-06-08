package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return "New Value"
		}, // New function to create a new value if the pool is empty
	}
	pool.Put("Hello")
	pool.Put("World")
	pool.Put("!")

	for i := 0; i < 10; i++ {
		go func() {
			value := pool.Get()
			fmt.Println(value)
			pool.Put(value)
		}()
	}

	time.Sleep(2 * time.Second)
}
