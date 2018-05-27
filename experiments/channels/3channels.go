// me experimenting
package main

import "fmt"

var x int

func put(c chan int, x int) {
	c <- x
}
func get(c chan int) int {
	x := <- c
	return x
}

func main() {
	c := make(chan int)
	go put(c, 3)
	fmt.Println("get(c): ", get(c))
	go put(c, 5)
	x := <- c
	fmt.Println("x: ", x)
}

