package main

import (
	"fmt"
	"errors"
)

func Divide(num int, div int) (int, error) {
	if div == 0 {
		// We return zer value of int (0) and error.
		return 0, errors.New("cannot divide by 0")
	}
	return num / div, nil
}

func main() {
	divideBy := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, div := range divideBy {
		result, err := Divide(100, div)
		if err != nil {
			fmt.Printf("100 / %d error: %s\n", div, err)
		} else {
			fmt.Printf("100 / %d = %d\n", div, result) 
		}
	}
}