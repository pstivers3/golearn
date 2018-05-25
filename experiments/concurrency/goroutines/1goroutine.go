// from introducing go

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond) // sleep for 100 milliseconds
		fmt.Println(s)
	}
}

func main() {
	go say("world") // goroutine
	say("hello")
}
