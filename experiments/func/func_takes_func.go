package main

import (
	"fmt"
	"math"
)

// a function can take another function as an argument
// compute will return func(3,4) of any func that takes two float64 arguments and returns a float 64
func compute(fn func(float64, float64) float64) float64 { // declaration ?? q
	return fn(3, 4)
}


func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
	fmt.Println(math.Pow(3,4))
	fmt.Println(compute(math.Pow) == math.Pow(3,4))
}

