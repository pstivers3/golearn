// from introducing go

package main

import (
	"fmt"
	"time"
	"math/rand"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		// //  sleep random milliseconds up to 250
		// amt := time.Duration(rand.Intn(250))
		// time.Sleep(time.Millisecond * amt)
		time.Sleep(100 * time.Millisecond) // sleep for 100 milliseconds
		fmt.Println(s)
	}
}

func main() {
	go say("world") // if the second say() is a goroutine or if both are grs, the grs don't print ?? q
	say("hello")
}
