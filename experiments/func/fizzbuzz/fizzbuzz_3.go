// print 1 to 100, except;
// if factor of 3, print fizz
// if factor of 5, print buzz
// if factor of 3 and 5, print fizzbuzz

package main

import (
	"fmt"
)

func factorOf (num int, factor int) bool {
	if num%factor == 0 {
		return true
	} else {
		return false
	}
}

func bothTrue (test1 bool, test2 bool) bool {
	if test1 && test2 {
		return true
	} else {
		return false
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		isFactor3 := factorOf(i,3)
		isFactor5 := factorOf(i,5)
		isBoth := bothTrue(isFactor3, isFactor5)
		switch {
			case isBoth:
				fmt.Println("fizzbuzz")
			case isFactor3:
				fmt.Println("fizz")
			case isFactor5:
				fmt.Println("buzz")
			default:
				fmt.Println(i)
		}
	}
}

