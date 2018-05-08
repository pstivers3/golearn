package main

import (
	"errors"
	"fmt"
)

// include error handling in this simple function that adds 3 to the function's argument
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

//declare a custom error struct type
type argError struct {
	arg int
	prob string
}

// what is Error() ?
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func main() {
	x := 42
	if v, e := f1(x); e != nil {
		fmt.Println("f1 failed:", e)
	} else {
		fmt.Println("f1 worked:", v)
	}
}
