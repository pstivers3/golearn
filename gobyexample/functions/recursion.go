package main

import "fmt"

func main() {
	fmt.Println(fact(0))
	fmt.Println(fact(1))
	fmt.Println(fact(2))
	fmt.Println(fact(3))
	fmt.Println(fact(4))
	fmt.Println(fact(5))
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

