package main

import "fmt"

func main() {
	// make a channel of type string with buffer size 2
	messages := make(chan string, 2)

	// send a text value into the messages channel using a go routine
	messages <- "ping"
	messages <- "pong"

	// receive a value from the message channel
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	fmt.Println("start ping pong loop")
	for i:=0; i <3; i++ {
		messages <- "ping"
		fmt.Println(<-messages)
		messages <- "pong"
		fmt.Println(<-messages)
	}
}
