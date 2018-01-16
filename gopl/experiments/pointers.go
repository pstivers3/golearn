package main

import (
	"fmt"
)

func main () {
	x := 1
	p := &x // the address of x
	fmt.Println(*p) // prints the value stored at the address of x, aka in x
	*p = 2 // equivalent to x = 2. Stores 2 at the address of x.
	fmt.Println(x) // 2
	fmt.Println(*p) // 2
	fmt.Println(*&x) // 2
	if *&x == x {
		fmt.Println("yes")
	}
}
