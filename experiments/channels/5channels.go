// from golangbot.com

package main

import (
    "fmt"
)

func hello(done chan bool) {
    fmt.Println("Hello world goroutine")
    done <- true
}
func main() {
    done := make(chan bool)
    go hello(done)
	// receives data from done channel
	// does not use or store that data
	// blocks until data is sent from the go routine to the done channel 
    <-done
    fmt.Println("main function")
}
