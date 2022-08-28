package main

import "fmt"

// returns 2 anonumous funcs. One anon func returns and int, the other doesn't return anything
func counter(start int) (func() int, func()) {
	ctr := func() int {
		return start
	}

	incr := func() {
		start++
	}

	return ctr, incr
}

func main() {
	ctr, incr := counter(100)
	ctr1, incr1 := counter(100)

	incr()
	fmt.Println("counter: ", ctr())
	fmt.Println("counter1: ", ctr1())

	incr1()
	incr1()

	fmt.Println("counter: ", ctr())
	fmt.Println("counter1: ", ctr1())
}
