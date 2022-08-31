package main

import "fmt"

func main() {
	x := [2]int{}
	changeValueAtZeroIndex(x)
	fmt.Println(x) // Will print 0
}

func changeValueAtZeroIndex(array [2]int) {
	array[0] = 3
	fmt.Println("inside: ", array[0]) // Will print 3
}
