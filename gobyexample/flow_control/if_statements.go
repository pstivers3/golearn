package main

import "fmt"

func testIf(arg int) (string, int) {
	if arg == 42 {
		return "42 is bad: ", 1
	}
	return "arg is good: ", 0
}

func main() {
	error_message, error_code := testIf(42)
	fmt.Println(error_message, error_code)

	error_message, error_code = testIf(5)
	fmt.Println(error_message, error_code)
}
