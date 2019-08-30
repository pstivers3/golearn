// go run error.go 1 2 3 4
// go run error.go a b c # to throw panic

package main

import (
	//	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sum int
	for _, a := range os.Args[1:] {
		i, err := strconv.Atoi(a)
		if err != nil {
			panic(fmt.Sprintf("Invalid Value: %v", err))
		}
		sum += i
	}
	fmt.Printf("sum = %v\n", sum)
}
