package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// my additional tests
	fmt.Println(float64(2.1)) // converts to float?
	fmt.Println(math.Sqrt2)   // apparently a builtin method for Sqrt(2)
	fmt.Println(math.Sqrt(16))

}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
