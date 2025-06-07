package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			channel <- i * 2 // send even numbers to the channel
		}
		close(channel) // close the channel to indicate no more values will be sent
	}()

	for value := range channel {
		fmt.Println(value) // print each value received from the channel
	}
	fmt.Println("Done!")

}
