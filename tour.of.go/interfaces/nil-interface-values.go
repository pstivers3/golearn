// A nil interface means that at least one method is
// declared in the interface, but the method fuction
// does not exist.

package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	// calling the method M throws a runtime error
	// because there is no method M defined
	i.M()
}

func describe(i I) {
	fmt.Printf("(%f, %T)\n", i, i)
}