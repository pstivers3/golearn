// this program prints its command-line arguments
// uses explicit variable initialization and a range in the for loop

package main

import (
    "fmt"
    "os"
)

func main() {
    // var s, sep string
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}
