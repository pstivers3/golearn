package main

import "fmt"

func main() {
	s:= make([]int, 5, 10) // length 5, capacity 10
	printSlice(s[:])
	printSlice(s[:])
	s = append(s, 1, 1, 1, 1, 1)
	printSlice(s[:])
	s = append(s, 2, 2, 2, 2, 2)
	printSlice(s[:])
	s = append(s, 3, 3, 3, 3, 3, 3)
	printSlice(s[:])
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d %v\n",len(x), cap(x), x)
}

