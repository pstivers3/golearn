package main

import "fmt"

// returns an anonumous func. The anonymous func returns an int
func incI() func() int {
	i := 0
	return func() int {
	    i++
		return i
	}
}

func main() {
	// a variable needs to be initialized with the outer func and the variable used to reach the inner func
	i := incI()
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())

	fmt.Println(incI()) // why do these return a memory location ?? q
	fmt.Println(incI())
	fmt.Println(incI())
}

