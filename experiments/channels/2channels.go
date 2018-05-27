package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 1
	} ()
	x := <-c
	fmt.Printf("Type: %T\n", x)
	fmt.Println("x: ", x)
}

