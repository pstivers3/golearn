package main

import "fmt"

func main() {
	s := []int{6, 7, 8}
	for index, value := range s { // returns the index and value if both receiving variables are specified
		fmt.Println(index, value)
	}
	for index := range s { // returns the index if only one receiving variable is specified
		fmt.Println(index)
	}
	for _, value := range s { // if only the value is needed, use the emty identifier for the index
		fmt.Println(value)
	}
}
