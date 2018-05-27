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
	// fmt.Println("channel c is: ", c) // nil
	// fmt.Printf("Type of c is %T\n", c) // chan string
	c  = make(chan string)
	// fmt.Println("channel c is: ", c) // memory location ?? q
	// fmt.Printf("Type of c is %T\n", c) // chan string
	// why does ping or pong sometimes get printed twice in the beginning ?? q
	// why does it not happen with with extra print statements above in main ?? g
	// why does it not happen with the wait command below ?? q
	go pinger(c)
	// time.Sleep(time.Millisecond * 5)
	go ponger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
}
