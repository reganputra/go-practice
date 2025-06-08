package main

import "sync"

func doSomething(data *sync.Map, key int, group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)
	data.Store(key, key)
}

func main() {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	numGoroutines := 100

	for i := 0; i < numGoroutines; i++ {
		go doSomething(data, i, group)
	}

	group.Wait()

	// Print the results
	data.Range(func(key, value interface{}) bool {
		println("Key:", key.(int), "Value:", value.(int))
		return true
	})

}
