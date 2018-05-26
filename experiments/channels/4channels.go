// doesn't run ?? q

package main

import "fmt"

var x int

func put(x int) {
	c := make(chan int)
	c <- x
}
func get() int {
	x <- c
	return x
}

func main() {
	go put(3)
	y := get()
	fmt.Println("y: ", y)

//	var input string
//	fmt.Scanln(&input)
}

