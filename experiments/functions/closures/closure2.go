package main

import "fmt"

// outer1 calls closure foo()
func outer1(name string) {
	// variable in outer function
	text := "outer1 says " + name

	// foo is an inner function, and is a closure.
	// closures have access to outer variables even after the block is exited
	foo := func() {
		fmt.Println(text)
	}

	// call the closure
	foo()
}

// outer2 returns the closure foo
func outer2(name string) func() {
	text := "outer2 says " + name

	foo := func() {
		fmt.Println(text)
	}

	return foo
}

func main() {
	outer1("hello")
	boo := outer2("hello")
	boo()
}
