package main

import (
	"fmt"
)

func main() {
	// coinflip
	fmt.Println(flipcoin("heads"))
	fmt.Println(flipcoin("tails"))
	fmt.Println(flipcoin("oops"))

	//signnum
	fmt.Println(Signum(22))
	fmt.Println(Signum(-11))
	fmt.Println(Signum(0))
}

func flipcoin(outcome string) int {
	//clunky but illustrates some things 
	heads := 0
	tails := 0
	result :=0
	switch outcome {
	case "heads":
		heads++
		result = heads
	case "tails":
		tails--
		result = tails
	default:
	}
	return result
}

func Signum(x int) int {
	switch {
	case x > 0:
		return +1
	case x < 0:
		return -1
	default:
		return 0
	}
}

