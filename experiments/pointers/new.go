package main

import (
	"fmt"
)

func main () {
    fmt.Println(newInt1())
    fmt.Println(newInt2())
    fmt.Println(delta(2,5))
}

func newInt1() *int {
	// new(type) returns new memory location for the type
	return new(int)
}

func newInt2() *int {
	var dummy int
	return  &dummy
}

func delta(old, new int) int {
	return new - old
}
