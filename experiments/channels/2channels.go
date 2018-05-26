// from introducing go

package main

import (
	"fmt"
	"time"
)

// c is type channel of strings
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string // channel declaration. Default value is nil
	fmt.Println("channel c is: ", c) // nil
	fmt.Printf("Type of c is %T\n", c)
	c  = make(chan string)
	fmt.Println("channel c is: ", c) // memory location ?? q
	fmt.Printf("Type of c is %T\n", c)
	// why does ping or pong sometimes get printed twice in the beginning ?? q
	// may not happen with the extra code above in main 
	// var c chan string = make(chan string) // original initialization in example
	// c := make(chan string) // alternative initialization
	go pinger(c)
	go ponger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
}
