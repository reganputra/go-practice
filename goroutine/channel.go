package main

import (
	"fmt"
	"time"
)

func giveResponse(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Response from goroutine"

}

func main() {
	channel := make(chan string)

	go func() {
		defer close(channel)
		time.Sleep(2 * time.Second)
		channel <- "Hello, World!"
		fmt.Println("Finished sending message to channel")
	}()

	fmt.Println("Waiting for message from channel...")

	data := <-channel
	time.Sleep(3 * time.Second)
	fmt.Println("Received from channel:", data)

	responseChannel := make(chan string)
	go giveResponse(responseChannel)
	defer close(responseChannel)
	fmt.Println("Waiting for response from goroutine...")
	response := <-responseChannel
	fmt.Println("Received response:", response)

}
