package main

func main() {
	// Channels -> used to communicate between goroutines, can be buffered or unbuffered
	// Channels are typed, meaning you can only send and receive values of a specific type
	// Create a channel to communicate between goroutines
	messageChannel := make(chan string)

	// Start a goroutine that sends a message to the channel
	go func() { // -> anonymous function
		messageChannel <- "Hello from the goroutine!"
	}()

	// Receive the message from the channel and print it
	message := <-messageChannel
	println(message)
}