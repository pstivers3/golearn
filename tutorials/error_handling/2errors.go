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

// create custom type that implements the Error() method
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// include custom *argError in this simple function that adds 3 to the function's argment
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
    // use the error handling built into f1
	for _, i := range []int{7,42} {
		if r, e := f1(i); e!= nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7,42} {
		if r, e := f2(i); e!= nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

    // Question. Not sure how this one works ?? q
	_,e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
