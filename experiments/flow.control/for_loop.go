package main

import "fmt"

func main() {
	// conditional is checked at the top of the loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		// i increments at the bottom of the loop
	}
}
