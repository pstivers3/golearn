// starts with solution to exercise-loops-and-functions.go

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) float64 {
	z := 1.0
	acc_delta := 0.01
	if x < 0 {
		fmt.Println(ErrNegativeSqrt(x))
		return 0
	}
	for math.Abs(z*z - x) > acc_delta {
		z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(-2))
	fmt.Println(Sqrt(2))
}
