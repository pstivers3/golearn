// this program prints its command-line arguments
// uses default Println formatting

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:])
}
