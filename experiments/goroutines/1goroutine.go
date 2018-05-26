// from introducing go

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond) // sleep for 100 milliseconds
		fmt.Println(s)
	}
}

func main() {
	// the non-gr call needs to be second, to give the gr call time to run 
	go say("one")
	say("two")
}
