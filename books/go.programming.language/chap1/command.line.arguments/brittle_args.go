// echos command line args
// the Args index must not exceed the end of the Args array
// needs two command line arguments
// brittle

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
}
