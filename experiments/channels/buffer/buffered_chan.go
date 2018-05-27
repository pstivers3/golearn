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

	// no bufer
	ch := make (chan string)
	// why doesn't this next line work, if comment it in, and comment out the separate goroutine below. ?? q
	// ch <- "hello"
	// apparently if unbuffered, sender has to be inside a separate goroutine
	go func() {
		 ch <- "hello"
	}()

	fmt.Println(<-ch)
}
