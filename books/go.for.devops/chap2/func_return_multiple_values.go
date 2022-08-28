package main

import "fmt"

func main() {
	result, remainder := divide(11, 3)
	fmt.Printf("Results: %d, Remainder %d", result, remainder)
}

// num == numberator
// div == divider
// res == result
// rem == remainder
func divide(num, div int) (res, rem int) {
	// variables are already initialized in return declaration in line immediately above.
	res = num / div
	rem = num % div
	return res, rem
}
