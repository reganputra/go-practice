package main

import "fmt"

func main() {

	channel := make(chan string, 2) // Buffered channel with a capacity of 2

	channel <- "Hello, World!"            // Send a message to the channel
	channel <- "Buffered channel message" // Send another message to the channel

	fmt.Println(<-channel)
	fmt.Println(<-channel) // Receive messages from the channel

}
