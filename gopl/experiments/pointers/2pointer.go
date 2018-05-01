package main

import (
	"fmt"
)

func main () {
	var p = f()
	fmt.Println(f()) // mem address 1
	fmt.Println(f()) // mem address 2
	fmt.Println(p)   // mem address 3
	fmt.Println(p)   // mem address 3. Why not a 4th address ??
}

func f() *int {
	v := 1
	return &v
}

