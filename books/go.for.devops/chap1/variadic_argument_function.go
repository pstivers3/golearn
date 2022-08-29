package main

import "fmt"

func main() {

	// using an array
	array := []int{1, 2, 3, 4, 5}
	fmt.Println("sum array: ", sumArray(array))

	// using variatic
	// fmt.Println(sum_variatic
	fmt.Println("sum variatic: ", sumVariatic(1, 2, 3))
}

// takes an array
func sumArray(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// same thing but variatic notation
func sumVariatic(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}
