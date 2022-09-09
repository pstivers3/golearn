// Empty interface means no method is declared in the interface
// An empty interface may hold values of any type

package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

// %v default format
// %T the type of the value
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}