package main

import "fmt"

func main() {
	var x int
	x = 5
	y := &x          // memory address of x (pointer to x)
	fmt.Println(x)   // 5
	fmt.Println(y)   // memory address memory address of x
	*y = 6           // store 6 in dereferenced address of x
	fmt.Println(x)   // 6
	fmt.Println(*y)  // 6
	fmt.Println(*&x) // 6
	fmt.Println(&x)  // memory address of x
}
