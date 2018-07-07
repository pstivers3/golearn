// print the numbers from 1 to 100, exept:
// if multiple of 3 print Fizz
// if multiple of 5 print Buzz
// if multiple of 3 and 5 print FizzBuzz

package main

import (
	"fmt"
)

func isFactor3 (num int) bool {
	if num%3 == 0 {
		return true
	} else {
		return false
	}
}

func isFactor5 (num int) bool {
	if num%5 == 0 {
		return true
	} else {
		return false
	}
}

func isFactorBoth (num int) bool {
	if num%3 == 0 && num%5 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		if isFactorBoth (i) {
			fmt.Println("FizzBuzz")
		} else if isFactor3(i) {
			fmt.Println("Fizz")
		} else if isFactor5(i) {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

