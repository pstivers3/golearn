// this program prints its command-line arguments
// uses explicit variable initialization and a range in the for loop

package main

import (
	"fmt"
	"os"
)

func main() {
	index := 1
	for _, arg := range os.Args[1:] {
	    fmt.Println(index,arg)
		index++
	}
}
