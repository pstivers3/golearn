package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	i := 1
	fmt.Println(math.Abs(z*z - x))
	for math.Abs(z*z - x) > 0.01 || i < 10 {
		z -= (z*z - x) / (2*z)
		fmt.Println("z:",z, "delta:", (z*z - x))
		i++
	}
	return z
}

func main() {
	fmt.Println(Sqrt(24))
}
