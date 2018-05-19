package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	acc_err := 0.01
	fmt.Printf("\n%14s %20s %20s\n", "estimate", "est squared", "error")
	for math.Abs(z*z - x) > acc_err {
		z -= (z*z - x) / (2*z)
		fmt.Printf("%14.5f %20.5f %20.5f\n", z , z*z, (z*z - x))
	}
	return z
}

func main() {
	x := float64(2); fmt.Printf("\n  Square root of %.2f = %.4f\n", x, Sqrt(x))
	x = float64(200); fmt.Printf("\n  Square root of %.2f = %.4f\n", x, Sqrt(x))
	x = float64(10000); fmt.Printf("\n  Square root of %.2f = %.4f\n\n", x, Sqrt(x))
}
