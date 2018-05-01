package main

import (
	"fmt"
)

func main () {
    fmt.Println(newInt1())
    fmt.Println(newInt2())
}

func newInt1() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return  &dummy
}


