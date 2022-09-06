package main

import (
	"fmt"
)

func main () {
    var p = f()
	fmt.Println(p) // prints the value stored at the address of x, aka in x
}

func f() *int {
	v := 1
	return &v
}
