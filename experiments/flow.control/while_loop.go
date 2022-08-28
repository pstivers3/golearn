package main

import "fmt"

func main() {
	var i int
	for {
		i++
		if i%10 == 0 {
			fmt.Println(i)
		}
		if i == 50 {
			break
		}
	}
}
