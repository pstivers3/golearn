package main

import "fmt"

func main() {
	i := incI()
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())

	fmt.Println(incI()) // why do these return a memory location?
	fmt.Println(incI())
	fmt.Println(incI())
}

func incI() func() int {
	i := 0
	return func() int {
	    i++
		return i
	}
}

