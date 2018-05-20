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
		return 1, ErrNegativeSqrt(x)
	}
	for math.Abs(z*z - x) > acc_delta {
		z -= (z*z - x) / (2*z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
