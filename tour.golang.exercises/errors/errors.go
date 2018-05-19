// starts with solution to exercise-loops-and-functions.go

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	z := 1.0
	acc_delta := 0.01
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	fmt.Printf("\n%14s %20s %20s\n", "estimate", "est squared", "delta")
	for math.Abs(z*z - x) > acc_delta {
		z -= (z*z - x) / (2*z)
		fmt.Printf("%14.5f %20.5f %20.5f\n", z , z*z, (z*z - x))
	}
	return z, nil
}

func main() {
	x := float64(-2); fmt.Println(Sqrt(x))
/*	
	x := float64(-2); fmt.Printf("\n  Square root of %.2f = %.4f\n", x, Sqrt(x))
	x = float64(200); fmt.Printf("\n  Square root of %.2f = %.4f\n", x, Sqrt(x))
	x = float64(10000); fmt.Printf("\n  Square root of %.2f = %.4f\n\n", x, Sqrt(x))
*/
}
