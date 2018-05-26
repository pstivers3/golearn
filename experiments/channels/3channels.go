package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 1
	} ()
	// why do these lines need to be outside the goroutine in order to print ?? q
	// no exception is thrown if lines are inside the goroutine, but nothing is printed
	x := <-c
	fmt.Printf("Type: %T\n", x)
	fmt.Println("x: ", x)
}

