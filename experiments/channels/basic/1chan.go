package main

import "fmt"

func main() {
	// make a channel of type string
	messages := make(chan string)

	// send a text value into the messages channel using a go routine
	go func() {
		messages <- "ping"
	}()

	// receive a value from the message channel
	msg := <-messages
	fmt.Println(msg)
}
