package main

import "fmt"

func main() {
	// declare an array of 5 strings, initialized to its zero value
	var fruits [5]string
	fruits[0] = "apple"
	fruits[1] = "orange"
	fruits[2] = "banana"
	fruits[3] = "grape"
	fruits[4] = "plum"

	// iterage over the array fo stings
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}

	// declare an array of 4 integers that is initialized with some values
	numbers := [4]int{10, 20, 30, 40}

	// iterate over the array of numbers
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}
}
